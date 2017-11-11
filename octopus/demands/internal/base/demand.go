package base

import (
	"time"

	"context"
	"net/http"

	"fmt"

	"net/url"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/safe"
	"github.com/clickyab/services/xlog"
)

// Demand is a single structure to handle demand data from database.
// loaded on initialize and on signals
type Demand struct {
	ID                  int64                 `db:"id" json:"id"`
	FName               string                `db:"name" json:"name"`
	FType               exchange.DemandType   `db:"type" json:"type"`
	GetURL              string                `db:"get_url" json:"get_url"`
	MinuteLimit         int64                 `db:"minute_limit" json:"minute_limit"`
	HourLimit           int64                 `db:"hour_limit" json:"hour_limit"`
	DayLimit            int64                 `db:"day_limit" json:"day_limit"`
	WeekLimit           int64                 `db:"week_limit" json:"week_limit"`
	MonthLimit          int64                 `db:"month_limit" json:"month_limit"`
	IdleConnections     int                   `db:"idle_connection" json:"idle_connection"`
	Timeout             int64                 `db:"timeout" json:"timeout"`
	Active              int                   `db:"active" json:"active"`
	FHandicap           int64                 `json:"handicap" db:"handicap"`
	Share               int                   `json:"-" db:"share"`
	Rate                int                   `json:"-" db:"call_rate"`
	FWhiteListCountries mysql.StringJSONArray `json:"white_countrie" db:"white_countrie" `
	FExcludedSuppliers  mysql.StringJSONArray `json:"excluded_suppliers" db:"excluded_suppliers"`
	UserID              int64                 `json:"user_id" db:"user_id"`
	FTestMode           bool                  `json:"test_mode" db:"test_mode"`
	client              *http.Client
}

// Client get client
func (d *Demand) Client() *http.Client {
	return d.client
}

// HasLimits check demand limit
func (d *Demand) HasLimits() bool {
	if d.MinuteLimit == 0 &&
		d.HourLimit == 0 &&
		d.DayLimit == 0 &&
		d.WeekLimit == 0 &&
		d.MonthLimit == 0 {
		return true
	}
	mo, we, da, ho, mi := getCPM(d.Name())
	if mo > 0 && mo >= d.MonthLimit {
		return false
	}
	if we > 0 && we >= d.WeekLimit {
		return false
	}
	if da > 0 && da >= d.DayLimit {
		return false
	}
	if ho > 0 && ho >= d.HourLimit {
		return false
	}
	if mi > 0 && mi >= d.MinuteLimit {
		return false
	}
	return true
}

//EndPoint demand end point
func (d *Demand) EndPoint() string {
	return d.GetURL
}

//Name demand name
func (d *Demand) Name() string {
	return d.FName
}

//Win demand win action
func (d *Demand) Win(ctx context.Context, b exchange.Bid) {
	incCPM(d.Name(), b.Price())
	safe.GoRoutine(func() {
		u, err := url.Parse(b.WinURL())
		if err != nil {
			xlog.GetWithError(ctx, err).Debug("bid win url is not valid")
			return
		}
		tmp := u.Query()
		tmp.Add("win", b.ID())
		tmp.Add("cpm", fmt.Sprint(b.Price()))
		u.RawQuery = tmp.Encode()
		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			xlog.GetWithError(ctx, err).Debug("demand making win request failure")
			return
		}

		_, err = d.Client().Do(req)
		if err != nil {
			xlog.GetWithError(ctx, err).Debug("demand making win request failure")
			return
		}
	})

}

//Status demand status
func (d *Demand) Status(context.Context, http.ResponseWriter, *http.Request) {
	panic("implement me")
}

// Handicap demand handicap
func (d *Demand) Handicap() int64 {
	return d.FHandicap
}

// CallRate demand callrate
func (d *Demand) CallRate() int {
	if d.Rate > 1 {
		d.Rate = 1
	}
	if d.Rate > 100 {
		d.Rate = 100
	}
	return d.Rate
}

// WhiteListCountries  demand whiteListCountries
func (d *Demand) WhiteListCountries() []string {
	return d.FWhiteListCountries
}

// ExcludedSuppliers demand excludedSuppliers
func (d *Demand) ExcludedSuppliers() []string {
	return d.FExcludedSuppliers
}

// TestMode test mode
func (d *Demand) TestMode() bool {
	return d.FTestMode
}

// Type demand type (srtb/ortb)
func (d *Demand) Type() exchange.DemandType {
	return d.FType
}

// GetTimeout return the timeout for this demand
func (d *Demand) GetTimeout() time.Duration {
	if time.Duration(d.Timeout) < 100*time.Millisecond {
		return 100 * time.Millisecond
	}
	if time.Duration(d.Timeout) > time.Second {
		return time.Second
	}
	return time.Duration(d.Timeout)
}

// ActiveDemands list all active demands
func (m *Manager) ActiveDemands() []exchange.DemandBase {
	var res []exchange.DemandBase
	var demands []Demand
	_, err := m.GetRDbMap().Select(&demands, "SELECT * FROM demands WHERE active <> 0")
	assert.Nil(err)
	for i := range demands {
		demands[i].client = &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: demands[i].IdleConnections,
			},
			Timeout: func() time.Duration {
				if time.Duration(demands[i].Timeout) < 100*time.Millisecond {
					return 100 * time.Millisecond
				}
				if time.Duration(demands[i].Timeout) > time.Second {
					return time.Second
				}
				return time.Duration(demands[i].Timeout)
			}(),
		}

		res = append(res, &demands[i])
	}
	return res
}
