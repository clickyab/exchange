package suppliers

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/suppliers/internal/models"
	"clickyab.com/exchange/octopus/suppliers/internal/renderer"
	"clickyab.com/exchange/octopus/suppliers/internal/restful"
	"github.com/clickyab/services/mysql"

	"github.com/sirupsen/logrus"

	"clickyab.com/exchange/octopus/suppliers/internal/openrtb"
)

const (
	rest = "rest"
	rtb  = "rtb"
)

var (
	sm *supplierManager
)

type supplierManager struct {
	suppliers map[string]models.Supplier
	lock      *sync.RWMutex
}

func restRendererFactory(sup exchange.Supplier, in string) exchange.Renderer {
	switch in {
	case "rest":
		// TODO : /api is hardcoded
		// TODO : restRenderFactory should have no arg
		return renderer.NewRenderer()
	default:
		logrus.Panicf("supplier with key %s not found", in)
	}
	return nil
}

func (sm *supplierManager) loadSuppliers() {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	m := models.NewManager()
	sm.suppliers = m.GetSuppliers(restRendererFactory)
}

func (sm *supplierManager) Initialize() {
	sm.loadSuppliers()
	reloadChan := make(chan os.Signal)
	signal.Notify(reloadChan, syscall.SIGHUP)
	go func() {
		for i := range reloadChan {
			logrus.Infof("Reloading supplier config, due to signal %s", i)
			sm.loadSuppliers()
		}
	}()
}

// GetSupplierByKey return a single supplier by its id
func GetSupplierByKey(key string) (exchange.Supplier, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()

	if s, ok := sm.suppliers[key]; ok {
		return &s, nil
	}

	return nil, fmt.Errorf("supplier with key %s not found", key)
}

// GetSupplierByName return a single supplier by its name
func GetSupplierByName(name string) (exchange.Supplier, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	for i := range sm.suppliers {
		if sm.suppliers[i].Name() == name {
			return sm.suppliers[i], nil
		}
	}

	return nil, fmt.Errorf("supplier with name %s not found", name)
}

// GetBidRequest try to get an bid request from a http request
func GetBidRequest(key string, r *http.Request) (exchange.BidRequest, error) {
	sup, err := GetSupplierByKey(key)
	if err != nil {
		return nil, err
	}

	// Make sure the profit margin is added to the request
	switch sup.Type() {
	case rest:
		return restful.GetBidRequest(sup, r)
	case rtb:
		return openrtb.GetBidRequest(sup, r)
	default:
		logrus.Panicf("Not a supported type: %s", sup.Type())
		return nil, fmt.Errorf("not supported type: %s", sup.Type())
	}
}

func init() {
	sm = &supplierManager{lock: &sync.RWMutex{}}
	mysql.Register(sm)
}
