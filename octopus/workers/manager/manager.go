package manager

import (
	"fmt"
	"time"

	"clickyab.com/exchange/octopus/workers/internal"
	"clickyab.com/exchange/services/assert"
	"clickyab.com/exchange/services/config"
	"clickyab.com/exchange/services/mysql"
	"clickyab.com/exchange/services/safe"
)

var (
	limit   = config.RegisterInt("octopus.worker.manager.limit", 1000, "the limit for points in manager")
	timeout = config.RegisterDuration("octopus.worker.manager.timeout", time.Minute, "the timeout to flush data")
)

type starter struct {
	channel chan internal.TableModel
}

func (s *starter) Initialize() {
	internal.RegisterAggregator(s)
	safe.GoRoutine(func() {
		worker(s.channel)
	})
}

func (s *starter) Channel() chan<- internal.TableModel {
	return s.channel
}

func worker(c chan internal.TableModel) {
	supDemSrcTable := make(map[string]*internal.TableModel)
	supSrcTable := make(map[string]*internal.TableModel)

	t := *timeout
	if t < 10*time.Second {
		t = 10 * time.Second
	}
	var counter = 0
	var ack internal.Acknowledger

	defer func() {
		if ack != nil {
			assert.Nil(ack.Nack(true, true))
		}
	}()

	flushAndClean := func() {
		err := flush(supDemSrcTable, supSrcTable)
		if ack != nil {
			if err == nil {
				assert.Nil(ack.Ack(true))
			} else {
				assert.Nil(ack.Nack(true, true))
			}
		}
		ack = nil
		counter = 0
		supDemSrcTable = make(map[string]*internal.TableModel)
		supSrcTable = make(map[string]*internal.TableModel)
	}
	ticker := time.NewTicker(t)

	for {
		select {
		case p := <-c:

			if p.Time == 0 {
				assert.NotNil(nil, "Time should not be equal 0")
			}
			if p.Source == "" || p.Supplier == "" {
				assert.NotNil(nil, "Source and supplier can not be empty")
			}
			ack = p.Acknowledger

			supSrcTableKey := fmt.Sprint(p.Time, p.Supplier, p.Source)
			supSrcTable[supSrcTableKey] = aggregate(supSrcTable[supSrcTableKey], p)

			if p.Demand != "" {
				supDemSrcKey := fmt.Sprint(p.Time, p.Supplier, p.Source, p.Demand)
				supDemSrcTable[supDemSrcKey] = aggregate(supDemSrcTable[supDemSrcKey], p)
			}

			counter++

			if counter > *limit {
				flushAndClean()
			}

		case <-ticker.C:
			flushAndClean()
		}
	}
}

func aggregate(a *internal.TableModel, b internal.TableModel) *internal.TableModel {
	if a == nil {
		return &b
	}

	assert.True(a.Time == b.Time, "[BUG] times are not same")

	a.RequestInCount += b.RequestInCount
	a.RequestOutCount += b.RequestOutCount
	a.ImpressionInCount += b.ImpressionInCount
	a.ImpressionOutCount += b.ImpressionOutCount
	a.WinCount += b.WinCount
	a.DeliverCount += b.DeliverCount
	a.WinBid += b.WinBid
	a.DeliverBid += b.DeliverBid
	a.Profit += b.Profit

	return a
}

func init() {
	//make sure worker start after mysql
	mysql.Register(&starter{channel: make(chan internal.TableModel)})
}
