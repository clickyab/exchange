package models

import (
	"assert"
	"database/sql/driver"
	"errors"
	"time"
)

type (
	DemandType string
)

const (
	DemandTypeRest DemandType = "rest"
)

var (
	// TODO : Watch it, if you add a demand type add it here too
	allDemandTypes = []DemandType{
		DemandTypeRest,
	}
)

// Demand is a single structure to handle demand data from database. loaded on initialize and on
// signals
type Demand struct {
	ID     int64      `db:"id" json:"id"`
	Name   string     `db:"name" json:"name"`
	Type   DemandType `db:"type" json:"type"`
	GetURL string     `db:"get_url" json:"get_url"`
	WinURL string     `db:"win_url" json:"win_url"`

	MinuteLimit int64 `db:"minute_limit" json:"minute_limit"`
	HourLimit   int64 `db:"hour_limit" json:"hour_limit"`
	DayLimit    int64 `db:"day_limit" json:"day_limit"`
	WeekLimit   int64 `db:"week_limit" json:"week_limit"`
	MonthLimit  int64 `db:"month_limit" json:"month_limit"`

	IdleConnections int   `db:"idle_connections" json:"idle_connections"`
	Timeout         int64 `db:"timeout" json:"timeout"`

	Active int `db:"active" json:"active"`
}

// IsValid try to validate enum value on ths type
func (e DemandType) IsValid() bool {
	for _, i := range allDemandTypes {
		if i == e {
			return true
		}
	}
	return false
}

// Scan convert the json array ino string slice
func (e *DemandType) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return errors.New("unsupported type")
	}
	if !DemandType(b).IsValid() {
		return errors.New("invaid value")
	}
	*e = DemandType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e DemandType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, errors.New("invalid demand type")
	}
	return string(e), nil
}

func (d Demand) GetTimeout() time.Duration {
	if time.Duration(d.Timeout) < 100*time.Millisecond {
		return 100 * time.Millisecond
	}
	if time.Duration(d.Timeout) > time.Second {
		return time.Second
	}
	return time.Duration(d.Timeout)
}

func (m *Manager) ActiveDemands() []Demand {
	var res []Demand
	_, err := m.GetRDbMap().Select(&res, "SELECT * FROM demands WHERE active <> 0")
	assert.Nil(err)

	return res
}