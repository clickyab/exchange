package base

import (
	"context"
	"io"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/ortb"
	"clickyab.com/exchange/octopus/exchange/srtb"
	ortb2 "clickyab.com/exchange/octopus/suppliers/internal/ortb"
	srtb2 "clickyab.com/exchange/octopus/suppliers/internal/srtb"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/sirupsen/logrus"
)

// supplier is a supplier in our system
type supplier struct {
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
	ICurrency     string                `json:"currency" db:"currency"`
}

func (s *supplier) GetBidRequest(_ context.Context, r *http.Request) (exchange.BidRequest, error) {
	switch s.Type() {
	case exchange.SupplierSRTB:
		return srtb.NewSimpleRTBFromRequest(s, r)
	case exchange.SupplierORTB:
		return ortb.NewOpenRTBFromRequest(s, r)
	default:
		logrus.Panicf("invalid type %s", s.Type())
		return nil, nil
	}
}

func (s *supplier) RenderBidResponse(ctx context.Context, w io.Writer, b exchange.BidResponse) http.Header {
	switch s.Type() {
	case exchange.SupplierSRTB:
		return srtb2.RenderBidResponse(ctx, s, w, b)
	case exchange.SupplierORTB:
		return ortb2.RenderBidResponse(ctx, s, w, b)
	default:
		logrus.Panicf("invalid type %s", s.Type())
		return nil
	}
}

// Currency is the suppliers currency considered in database
func (s supplier) Currency() string {
	return s.ICurrency
}

// Name of this supplier
func (s supplier) Name() string {
	return s.SName
}

// FloorCPM of this supplier
func (s supplier) FloorCPM() int64 {
	return s.SFloorCPM
}

// SoftFloorCPM of this supplier
func (s supplier) SoftFloorCPM() int64 {
	return s.SSoftFloorCPM
}

// ExcludedDemands of this supplire @TODO implement this
func (s supplier) ExcludedDemands() []string {
	return s.Excluded
}

// CountryWhiteList is the country allowed by this supplier @TODO implement this
func (supplier) CountryWhiteList() []exchange.Country {
	return nil
}

// Type is the supplier type
func (s supplier) Type() string {
	return s.SType
}

// Share of the supplier
func (s supplier) Share() int {
	if s.SShare > 100 {
		return 100
	}
	if s.SShare < 0 {
		s.SShare = 0
	}
	return s.SShare
}

// TestMode return true if this is a test demand
func (s supplier) TestMode() bool {
	return s.Test
}

// GetSuppliers return all suppliers
func (m *Manager) GetSuppliers() map[string]exchange.Supplier {
	q := "SELECT * FROM suppliers WHERE active <> 0"
	var res []supplier
	_, err := m.GetRDbMap().Select(&res, q)
	assert.Nil(err)
	ret := make(map[string]exchange.Supplier, len(res))
	for i := range res {
		ret[res[i].Key] = &res[i]
	}

	return ret
}
