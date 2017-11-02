package base

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/mysql"
)

// Supplier is a supplier in our system
type Supplier struct {
	ID            int64                 `json:"id" db:"id"`
	SName         string                `json:"name" db:"name"`
	SType         string                `json:"type" db:"type"`
	Key           string                `json:"-" db:"key"`
	SFloorCPM     int64                 `json:"floor_cpm" db:"floor_cpm"`
	SSoftFloorCPM int64                 `json:"soft_floor_cpm" db:"soft_floor_cpm"`
	UnderFloor    int                   `json:"under_floor" db:"under_floor"`
	Excluded      mysql.StringJSONArray `json:"excluded_demands" db:"excluded_demands"`
	SShare        int                   `json:"-" db:"share"`
	SActive       int                   `json:"-" db:"active"`
	UserID        int64                 `json:"user_id" db:"user_id"`
	Test          bool                  `json:"test_mode" db:"test_mode"`
	Click         string                `json:"click_mode" db:"click_mode"`
}

// Name of this supplier
func (s Supplier) Name() string {
	return s.SName
}

// FloorCPM of this supplier
func (s Supplier) FloorCPM() int64 {
	return s.SFloorCPM
}

// SoftFloorCPM of this supplier
func (s Supplier) SoftFloorCPM() int64 {
	return s.SSoftFloorCPM
}

// ExcludedDemands of this supplire @TODO implement this
func (s Supplier) ExcludedDemands() []string {
	return s.Excluded
}

// CountryWhiteList is the country allowed by this supplier @TODO implement this
func (Supplier) CountryWhiteList() []exchange.Country {
	return nil
}

// Type is the supplier type
func (s Supplier) Type() string {
	return s.SType
}

// Share of the supplier
func (s Supplier) Share() int {
	if s.SShare > 100 {
		return 100
	}
	if s.SShare < 0 {
		s.SShare = 0
	}
	return s.SShare
}

// TestMode return true if this is a test demand
func (s Supplier) TestMode() bool {
	return s.Test
}
