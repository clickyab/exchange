package suppliers

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/suppliers/internal/base"
	"github.com/clickyab/services/mysql"
	"github.com/sirupsen/logrus"
)

var (
	sm = &supplierManager{}
)

type supplierManager struct {
	suppliers map[string]exchange.Supplier
	sync.RWMutex
}

func (sm *supplierManager) loadSuppliers() {
	sm.Lock()
	defer sm.Unlock()

	m := base.NewManager()
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
	sm.RLock()
	defer sm.RUnlock()

	if s, ok := sm.suppliers[key]; ok {
		return s, nil
	}

	return nil, fmt.Errorf("supplier with key %s not found", key)
}

func init() {
	mysql.Register(sm)
}
