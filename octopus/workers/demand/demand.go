package demand

import (
	"context"
	"time"

	"clickyab.com/exchange/octopus/models"
	"clickyab.com/exchange/octopus/workers/internal/datamodels"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/safe"
)

type model struct {
	Request struct {
		TrackID   string    `json:"track_id"`
		Time      time.Time `json:"time"`
		Inventory struct {
			Name     string `json:"name"`
			Supplier struct {
				FloorCPM     int64  `json:"floor_cpm"`
				SoftFloorCPM int64  `json:"soft_floor_cpm"`
				Name         string `json:"name"`
				Share        int    `json:"share"`
			} `json:"supplier"`
			Domain string `json:"domain"`
		} `json:"inventory"`
	} `json:"request"`
	Response struct {
		ID   string `json:"id"`
		Bids []struct {
			ID         string   `json:"id"`
			ImpID      string   `json:"imp_id"`
			Price      int64    `json:"price"`
			WinURL     string   `json:"win_url"`
			Categories []string `json:"categories"`
			AdID       string   `json:"ad_id"`
			AdHeight   int      `json:"ad_height"`
			AdWidth    int      `json:"ad_width"`
			AdDomain   []string `json:"ad_domain"`
			Demand     struct {
				Name string `json:"name"`
			} `json:"demand"`
		} `json:"bids"`
	} `json:"response"`
}

var extraCount = config.RegisterInt("octopus.workers.extra.count", 10, "the consumer count for a worker")

type consumer struct {
	ctx      context.Context
	workerID string
}

func (s *consumer) Initialize(ctx context.Context) {
	s.ctx = ctx
	broker.RegisterConsumer(s)

	for i := 1; i < extraCount.Int(); i++ {
		broker.RegisterConsumer(
			&consumer{
				ctx:      ctx,
				workerID: <-random.ID,
			},
		)
	}
}

func (consumer) Topic() string {
	return "demand"
}

func (consumer) Queue() string {
	return "demand_aggregate"
}

func (s *consumer) Consume() chan<- broker.Delivery {
	chn := make(chan broker.Delivery, 0)
	done := s.ctx.Done()
	safe.ContinuesGoRoutine(func(cnl context.CancelFunc) {
		var del broker.Delivery
		defer func() {
			if del != nil {
				del.Reject(false)
			}
		}()
		for {
			select {
			case del = <-chn:
				obj := model{}
				err := del.Decode(&obj)
				assert.Nil(err)
				var win int64
				for _, v := range obj.Response.Bids {
					if cpm := v.Price; cpm > 0 {
						win++
					}
				}

				datamodels.ActiveAggregator().Channel() <- datamodels.TableModel{
					Supplier:           obj.Request.Inventory.Supplier.Name,
					Source:             obj.Request.Inventory.Domain,
					Demand:             obj.Response.Bids[0].Demand.Name,
					Time:               models.FactTableID(obj.Request.Time),
					RequestOutCount:    1,
					ImpressionOutCount: int64(len(obj.Response.Bids)),
					AdInCount:          win,
					Acknowledger:       del,
					WorkerID:           s.workerID,
				}
			case <-done:
				cnl()
				del = nil
				return
			}
		}
	}, time.Second)
	return chn
}

func init() {
	initializer.Register(&consumer{workerID: <-random.ID}, 1000)
}
