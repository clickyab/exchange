package show

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
	return "show"
}

func (consumer) Queue() string {
	return "show_aggregate"
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
				obj := jsonbackend.ShowWorker{}
				err := del.Decode(&obj)
				assert.Nil(err)

				datamodels.ActiveAggregator().Channel() <- datamodels.TableModel{
					Supplier:     obj.Supplier,
					Source:       obj.Publisher,
					Demand:       obj.Demand,
					DeliverBid:   obj.Winner,
					DeliverCount: 1,
					Profit:       obj.Profit,
					// TODO : why this is different with other?? make it same.
					Time:         models.FactTableID(timestampToTime(obj.Time)),
					Acknowledger: del,
					WorkerID:     s.workerID,
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

func timestampToTime(s string) time.Time {
	res, err := time.Parse(time.RFC3339, s)
	assert.Nil(err)
	return res

}

func init() {
	initializer.Register(&consumer{workerID: <-random.ID}, 10000)
}
