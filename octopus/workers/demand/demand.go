package demand

import (
	"context"
	"time"

	"clickyab.com/exchange/octopus/exchange/materialize/jsonbackend"
	"clickyab.com/exchange/octopus/models"
	"clickyab.com/exchange/octopus/workers/internal/datamodels"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/safe"
)

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
				obj := jsonbackend.Demand{}
				err := del.Decode(&obj)
				assert.Nil(err)

				datamodels.ActiveAggregator().Channel() <- datamodels.TableModel{
					Supplier:        obj.Supplier,
					Source:          obj.Source,
					Demand:          obj.Demand,
					Time:            models.FactTableID(time.Now()),
					RequestOutCount: 1,
					BidOutCount:     int64(obj.BidLen),
					AdInCount:       int64(obj.TotalPrice),
					Acknowledger:    del,
					WorkerID:        s.workerID,
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
