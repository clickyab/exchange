package demands

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"clickyab.com/exchange/octopus/core"
	"github.com/clickyab/services/mysql"

	"clickyab.com/exchange/octopus/demands/internal/base"
	"clickyab.com/exchange/octopus/demands/internal/ortb"
	"clickyab.com/exchange/octopus/exchange"
	"github.com/sirupsen/logrus"
)

type demandManager struct {
	activeDemands []exchange.DemandBase
	lock          *sync.RWMutex
}

func (dm *demandManager) loadDemands() {
	dm.lock.Lock()
	defer dm.lock.Unlock()
	dm.activeDemands = base.NewManager().ActiveDemands()
	core.ResetProviders()
	for _, demand := range dm.activeDemands {
		switch demand.Type() {
		case exchange.DemandTypeSrtb:
			// TODO: register srtb
			//core.Register(restful.NewRestfulClient(demand, getRawBidRequest), demand.GetTimeout())
		case exchange.DemandTypeOrtb:
			core.Register(&ortb.Demand{DemandBase: demand}, demand.GetTimeout())
		default:
			logrus.Panicf("Not supported demand type : %s", demand.Type)
		}
	}
}

func (dm *demandManager) Initialize() {
	dm.loadDemands()
	reloadChan := make(chan os.Signal)
	signal.Notify(reloadChan, syscall.SIGHUP)
	go func() {
		for i := range reloadChan {
			logrus.Infof("Reloading demands config, due to signal %s", i)
			dm.loadDemands()
		}
	}()
}

func init() {
	mysql.Register(&demandManager{lock: &sync.RWMutex{}})
}
