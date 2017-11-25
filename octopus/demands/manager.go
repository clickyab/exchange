package demands

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/clickyab/services/hub"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/safe"

	"clickyab.com/exchange/octopus/demands/internal/base"
	"clickyab.com/exchange/octopus/dispatcher"
	"clickyab.com/exchange/octopus/exchange"
	"github.com/sirupsen/logrus"
)

type demandManager struct {
	activeDemands []exchange.Demand
	sync.RWMutex
}

func (dm *demandManager) loadDemands() {
	dm.Lock()
	defer dm.Unlock()
	dm.activeDemands = base.NewManager().ActiveDemands()
	dispatcher.ResetProviders()
	for _, demand := range dm.activeDemands {
		dispatcher.Register(demand, demand.GetTimeout())
	}
}

func (dm *demandManager) Initialize() {
	dm.loadDemands()
	reloadChan := make(chan os.Signal)
	signal.Notify(reloadChan, syscall.SIGHUP)
	s := hub.Subscribe("reload")
	safe.GoRoutine(func() {
		for {
			select {
			case <-s:
				logrus.Infof("Reloading demands config, due to publisher")
				dm.loadDemands()
			case i := <-reloadChan:
				logrus.Infof("Reloading demands config, due to signal %s", i)
				dm.loadDemands()
			}
		}
	})
}

func init() {
	mysql.Register(&demandManager{})
}
