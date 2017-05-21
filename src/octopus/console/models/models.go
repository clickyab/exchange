package models

import (
	"octopus/console/modules/user/aaa"
	"services/mysql"
)

//UserTableFull table users
const UserTableFull = "users"

//Manager models
type Manager struct {
	mysql.Manager
}

//NewOctManager octopus manager
func NewOctManager() *Manager {
	return &Manager{}
}

//Initialize manager
func (m *Manager) Initialize() {
	m.AddTableWithName(
		aaa.User{},
		UserTableFull,
	).SetKeys(
		true,
		"ID",
	)
}

func init() {
	mysql.Register(NewOctManager())
}
