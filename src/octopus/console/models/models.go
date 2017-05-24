package models

import (
	"fmt"

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

// GetUser returns true if the user and pass is ok
func (m *Manager) GetUser(username, password string) bool {
	query := fmt.Sprintf(`SELECT * FROM %s where username=? AND password=?`, UserTableFull)
	user := &aaa.User{}
	err := m.GetRDbMap().SelectOne(
		user,
		query,
		username,
		password,
	)
	if err != nil || user == nil {
		return false
	}
	return true
}

func init() {
	mysql.Register(NewOctManager())
}
