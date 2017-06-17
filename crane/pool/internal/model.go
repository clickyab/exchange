package internal

import (
	"clickyab.com/exchange/crane/entity"
	"clickyab.com/exchange/services/mysql"
)

// TODO: needs to have migration
const poolQuery string = ``

// Manager is a manager
type Manager struct {
	mysql.Manager
}

// Initialize this package initializer
func (Manager) Initialize() {
}

// GetAllActiveAds is needed for ad pool
func (m *Manager) GetAllActiveAds() ([]entity.Advertise, error) {
	var resp []entity.Advertise
	_, err := m.GetRDbMap().Select(&resp, poolQuery)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// NewManager returns this manager
func NewManager() *Manager {
	return &Manager{}
}

func init() {
	mysql.Register(&Manager{})
}
