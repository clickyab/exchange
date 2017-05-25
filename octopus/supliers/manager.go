package supliers

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/supliers/internal/models"
	"clickyab.com/exchange/octopus/supliers/internal/restful"
	"clickyab.com/exchange/octopus/supliers/internal/restful/renderer"
	"clickyab.com/exchange/services/mysql"

	"github.com/Sirupsen/logrus"
)

var sm *supplierManager

type supplierManager struct {
	suppliers map[string]models.Supplier
	lock      *sync.RWMutex
}

func restRendererFactory(sup exchange.Supplier, in string) exchange.Renderer {
	switch in {
	case "rest":
		return renderer.NewRestfulRenderer(sup, "/pixel/%s/%s")
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
			logrus.Infof("Reloding supplier config, due to signal %s", i)
			sm.loadSuppliers()
		}
	}()
}

// getSupplier return a single supplier by its id
func getSupplier(key string) (*models.Supplier, error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()

	if s, ok := sm.suppliers[key]; ok {
		return &s, nil
	}

	return nil, fmt.Errorf("supplier with key %s not found", key)
}

// GetImpression try to get an impression from a http request
func GetImpression(key string, r *http.Request) (exchange.Impression, error) {
	sup, err := getSupplier(key)
	if err != nil {
		return nil, err
	}
	// Make sure the profit margin is added to the request
	switch sup.SType {
	case "rest":
		return restful.GetImpression(sup, r)
	default:
		logrus.Panicf("Not a supported type: %s", sup.SType)
		return nil, fmt.Errorf("not supported type: %s", sup.SType)
	}
}

func init() {
	sm = &supplierManager{lock: &sync.RWMutex{}}
	mysql.Register(sm)
}