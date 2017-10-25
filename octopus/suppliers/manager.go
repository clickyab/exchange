package suppliers

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/suppliers/internal/base"
	"clickyab.com/exchange/octopus/suppliers/internal/ortb"
	"clickyab.com/exchange/octopus/suppliers/internal/srtb/restful"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/sirupsen/logrus"
)

var (
	sm *supplierManager
)

type supplierManager struct {
	suppliers map[string]exchange.Supplier
	lock      *sync.RWMutex
}

func (sm *supplierManager) loadSuppliers() {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	m := NewManager()
	sm.suppliers = m.GetSuppliers()
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
		return s, nil
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
	case "rest":
		return restful.GetBidRequest(sup, r)

	default:
		logrus.Panicf("Not a supported type: %s", sup.Type())
		return nil, fmt.Errorf("not supported type: %s", sup.Type())
	}
}

func init() {
	sm = &supplierManager{lock: &sync.RWMutex{}}
	mysql.Register(sm)
}

// Manager is the model manager
type Manager struct {
	mysql.Manager
}

// Initialize the model, its the interface, not really need this
func (m *Manager) Initialize() {

}

// NewManager return a new manager object
func NewManager() *Manager {
	return &Manager{}
}

func init() {
	mysql.Register(&Manager{})
}

// GetSuppliers return all suppliers
func (m *Manager) GetSuppliers() map[string]exchange.Supplier {
	q := "SELECT * FROM suppliers WHERE active <> 0"
	var res []base.Supplier
	_, err := m.GetRDbMap().Select(&res, q)
	assert.Nil(err)
	ret := make(map[string]exchange.Supplier, len(res))
	for i := range res {
		if res[i].Type() == exchange.SupplierORTB {
			ret[res[i].Key] = &ortb.Supplier{
				res[i],
			}
		}
		panic("[BUG] not a valid supplier type")
	}

	return ret
}
