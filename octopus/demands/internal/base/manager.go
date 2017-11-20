package base

import (
	"context"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/mysql"
)

// Manager is the model manager
type Manager struct {
	mysql.Manager
}

// Initialize the manager. nothing to do, just keep it in interface shape
func (m *Manager) Initialize() {

}

// NewManager return a new manager object
func NewManager() *Manager {
	return &Manager{}
}

func init() {
	mysql.Register(&Manager{})
}

// Provide provide ad for specified bid request
func Provide(ctx context.Context, dem exchange.Demand, bq exchange.BidRequest, ch chan exchange.BidResponse) {

}
