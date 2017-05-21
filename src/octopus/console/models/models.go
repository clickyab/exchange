package models

import (
	"octopus/console/modules/user/aaa"
	"services/mysql"
)

const UserTableFull = "users"

type Manager struct {
	mysql.Manager
}

func NewOctManager() *Manager {
	return &Manager{}
}

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
