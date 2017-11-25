package base

import (
	"bytes"
	"io"
	"io/ioutil"
	"time"

	"context"
	"net/http"

	"clickyab.com/exchange/octopus/biding"
	"clickyab.com/exchange/octopus/demands/internal/ortb"
	"clickyab.com/exchange/octopus/demands/internal/srtb"
	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/xlog"
	"github.com/sirupsen/logrus"
)

// demand is a single structure to handle demand data from database.
// loaded on initialize and on signals
type demand struct {
	ID                  int64                 `db:"id" json:"id"`
	FName               string                `db:"name" json:"name"`
	FType               exchange.DemandType   `db:"type" json:"type"`
	GetURL              string                `db:"get_url" json:"get_url"`
	MinuteLimit         float64               `db:"minute_limit" json:"minute_limit"`
	HourLimit           float64               `db:"hour_limit" json:"hour_limit"`
	DayLimit            float64               `db:"day_limit" json:"day_limit"`
	WeekLimit           float64               `db:"week_limit" json:"week_limit"`
	MonthLimit          float64               `db:"month_limit" json:"month_limit"`
	IdleConnections     int                   `db:"idle_connection" json:"idle_connection"`
	Timeout             int64                 `db:"timeout" json:"timeout"`
	Active              int                   `db:"active" json:"active"`
	FHandicap           int64                 `json:"handicap" db:"handicap"`
	Share               int                   `json:"-" db:"share"`
	Rate                int                   `json:"-" db:"call_rate"`
	FWhiteListCountries mysql.StringJSONArray `json:"white_countries" db:"white_countries"`
	FExcludedSuppliers  mysql.StringJSONArray `json:"excluded_suppliers" db:"excluded_suppliers"`
	UserID              int64                 `json:"user_id" db:"user_id"`
	FCurrencies         mysql.StringJSONArray `json:"currencies" db:"currencies"`
	FTestMode           bool                  `json:"test_mode" db:"test_mode"`
	client              *http.Client
}

func (d *demand) Currencies() []string {
	return d.FCurrencies
}

func (d *demand) GetBidResponse(ctx context.Context, r *http.Response, s exchange.Supplier) (exchange.BidResponse, error) {
	switch d.Type() {
	case exchange.DemandTypeSrtb:
		return srtb.GetBidResponse(ctx, d, r, s)
	case exchange.DemandTypeOrtb:
		return ortb.GetBidResponse(ctx, d, r, s)
	default:
		logrus.Panicf("Not supported demand type : %s", d.Type())
		return nil, nil
	}
}

func (d *demand) RenderBidRequest(ctx context.Context, w io.Writer, bq exchange.BidRequest) http.Header {
	switch d.Type() {
	case exchange.DemandTypeSrtb:
		return srtb.RenderBidRequest(ctx, d, w, bq)
	case exchange.DemandTypeOrtb:
		return ortb.RenderBidRequest(ctx, d, w, bq)
	default:
		logrus.Panicf("Not supported demand type : %s", d.Type())
		return nil
	}
}

func (d *demand) Provide(ctx context.Context, bq exchange.BidRequest, ch chan exchange.BidResponse) {
	defer close(ch)
	if !d.HasLimits() {
		return
	}
	buf := &bytes.Buffer{}

	header := d.RenderBidRequest(ctx, buf, bq)
	req, err := http.NewRequest("POST", d.EndPoint(), bytes.NewBuffer(buf.Bytes()))
	if err != nil {
		xlog.GetWithField(ctx, "exchange to demand request rendering", err.Error()).Debug()
		return
	}

	req.Header = header
	xlog.GetWithField(ctx, "key", d.Name()).Debug("calling demand")
	resp, err := d.client.Do(req.WithContext(ctx))
	if err != nil {
		xlog.GetWithError(ctx, err).Debug()
		return
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		xlog.GetWithField(ctx, "status", resp.StatusCode).Debug(string(body))
		return
	}
	result, err := d.GetBidResponse(ctx, resp, bq.Inventory().Supplier())
	if err != nil {
		xlog.GetWithError(ctx, err).Debug()
		return
	}
	ch <- result
}

// HasLimits check demand limit
func (d *demand) HasLimits() bool {
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
func (d *demand) EndPoint() string {
	return d.GetURL
}

//Name demand name
func (d *demand) Name() string {
	return d.FName
}

// Win demand win action
func (d *demand) Win(ctx context.Context, price float64, url string) {
	biding.DoBillGetRequest(ctx, d.client, url)
}

// Bill demand bill action
func (d *demand) Bill(ctx context.Context, price float64, url string) {
	incCPM(d.Name(), price)
	biding.DoBillGetRequest(ctx, d.client, url)
}

//Status demand status
func (d *demand) Status(context.Context, http.ResponseWriter, *http.Request) {
	panic("implement me")
}

// Handicap demand handicap
func (d *demand) Handicap() int64 {
	return d.FHandicap
}

// CallRate demand callrate
func (d *demand) CallRate() int {
	if d.Rate > 1 {
		d.Rate = 1
	}
	if d.Rate > 100 {
		d.Rate = 100
	}
	return d.Rate
}

// WhiteListCountries  demand whiteListCountries
func (d *demand) WhiteListCountries() []string {
	return d.FWhiteListCountries
}

// ExcludedSuppliers demand excludedSuppliers
func (d *demand) ExcludedSuppliers() []string {
	return d.FExcludedSuppliers
}

// TestMode test mode
func (d *demand) TestMode() bool {
	return d.FTestMode
}

// Type demand type (srtb/ortb)
func (d *demand) Type() exchange.DemandType {
	return d.FType
}

// GetTimeout return the timeout for this demand
func (d *demand) GetTimeout() time.Duration {
	if time.Duration(d.Timeout) < 100*time.Millisecond {
		return 100 * time.Millisecond
	}
	if time.Duration(d.Timeout) > time.Second {
		return time.Second
	}
	return time.Duration(d.Timeout)
}

// ActiveDemands list all active demands
func (m *Manager) ActiveDemands() []exchange.Demand {
	var res []exchange.Demand
	var demands []demand
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
