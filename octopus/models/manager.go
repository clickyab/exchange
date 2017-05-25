package models

import (
	"clickyab.com/exchange/services/mysql"
)

// Manager is the model manager
type Manager struct {
	mysql.Manager
}

// Initialize the manager. nothing to do, just keep it in interface shape
func (m *Manager) Initialize() {
	m.AddTableWithName(
		SupplierSourceDemand{},
		"sup_dem_src",
	).SetKeys(
		false,
		"Supplier",
		"Demand",
		"Source",
		"TimeID",
	)
	m.AddTableWithName(
		SupplierSource{},
		"sup_src",
	).SetKeys(
		false,
		"Supplier",
		"Source",
		"TimeID",
	)
}

// NewManager return a new manager object
func NewManager() *Manager {
	return &Manager{}
}

func init() {
	mysql.Register(NewManager())
}