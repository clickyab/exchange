package restful

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"

	"clickyab.com/exchange/octopus/demands/internal/models"

	"io/ioutil"

	"github.com/sirupsen/logrus"
)

type demand struct {
	client             *http.Client
	callRate           int
	dayLimit           int64
	encoder            func(exchange.Impression) interface{}
	endPoint           string
	handicap           int64
	hourLimit          int64
	key                string
	maxIdleConnections int
	minuteLimit        int64
	monthLimit         int64
	requestTimeout     time.Duration
	weekLimit          int64
	winPoint           *url.URL
	country            []string
	supplier           []string
	testMode           bool
}

func log(imp exchange.Impression) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"track_id": imp.TrackID(),
		"type":     "provider",
	})
}

func (d *demand) WhiteListCountries() []string {
	return d.country
}

func (d *demand) ExcludedSuppliers() []string {
	return d.supplier
}

func (*demand) Status(context.Context, http.ResponseWriter, *http.Request) {
	panic("implement me")
}

func (d *demand) Name() string {
	return d.key
}

func (d *demand) Provide(ctx context.Context, imp exchange.Impression, ch chan exchange.Advertise) {
	defer close(ch)
	if !d.hasLimits() {
		return
	}
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	if err := enc.Encode(d.encoder(imp)); err != nil {
		logrus.Debug(err)
		return
	}
	req, err := http.NewRequest("POST", d.endPoint, bytes.NewBuffer(buf.Bytes()))
	if err != nil {
		logrus.Debug(err)
		return
	}
	log(imp).WithField("key", d.key).Debug("calling demand")
	resp, err := d.client.Do(req.WithContext(ctx))
	if err != nil {
		logrus.Debug(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		log(imp).WithField("status", resp.StatusCode).Debug(string(body))
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)
	reader := bytes.NewReader(data)
	log(imp).WithField("key", d.key).WithField("result", string(data)).Debug("Call done")

	ads := []*restAd{}
	dec := json.NewDecoder(reader)
	defer resp.Body.Close()
	if err := dec.Decode(&ads); err != nil {
		logrus.Debug(err)
		return
	}

	log(imp).WithField("count", len(ads)).Debug("selected ad from pool")
	for i := range ads {
		ads[i].demand = d
		ch <- ads[i]
	}
}

func (d *demand) Win(ctx context.Context, id string, cpm int64) {
	incCPM(d.key, cpm)
	if d.winPoint == nil {
		return
	}
	u := *d.winPoint
	tmp := u.Query()
	tmp.Add("win", id)
	tmp.Add("cpm", fmt.Sprint(cpm))
	u.RawQuery = tmp.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		logrus.Debug(err)
		return
	}
	resp, err := d.client.Do(req.WithContext(ctx))
	if err != nil {
		logrus.Debug(err)
		return
	}

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted || resp.StatusCode == http.StatusNoContent {
		return
	}

	logrus.WithField("status", resp.StatusCode).Debug("winner call status code")
}

func (d demand) CallRate() int {
	return d.callRate
}

func (d *demand) Handicap() int64 {
	return d.handicap
}

func (d *demand) TestMode() bool {
	return d.testMode
}

func (d *demand) hasLimits() bool {
	if d.minuteLimit == 0 &&
		d.hourLimit == 0 &&
		d.dayLimit == 0 &&
		d.weekLimit == 0 &&
		d.monthLimit == 0 {
		return true
	}
	mo, we, da, ho, mi := getCPM(d.key)
	if mo > 0 && mo >= d.monthLimit {
		return false
	}
	if we > 0 && we >= d.weekLimit {
		return false
	}
	if da > 0 && da >= d.dayLimit {
		return false
	}
	if ho > 0 && ho >= d.hourLimit {
		return false
	}
	if mi > 0 && mi >= d.minuteLimit {
		return false
	}
	return true
}

func (d *demand) createConnection() {
	d.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: d.maxIdleConnections,
		},
		Timeout: d.requestTimeout,
	}
}

// NewRestfulClient return a new client for restful call
func NewRestfulClient(d models.Demand, encoder func(exchange.Impression) interface{}) exchange.Demand {
	var win *url.URL
	var err error
	if d.WinURL.Valid {
		win, err = url.Parse(d.WinURL.String)
		assert.Nil(err)
	}
	dm := &demand{
		endPoint:           d.GetURL,
		winPoint:           win,
		maxIdleConnections: d.IdleConnections,
		requestTimeout:     d.GetTimeout(),
		key:                d.Name,
		minuteLimit:        d.MinuteLimit,
		hourLimit:          d.HourLimit,
		dayLimit:           d.DayLimit,
		weekLimit:          d.WeekLimit,
		monthLimit:         d.MonthLimit,
		handicap:           d.Handicap,
		country:            d.WhiteListCountries,
		supplier:           d.ExcludedSuppliers,
		testMode:           d.TestMode != 0,
		encoder:            encoder,
	}
	dm.callRate = d.Rate
	if dm.callRate > 1 {
		dm.callRate = 1
	}
	if dm.callRate > 100 {
		dm.callRate = 100
	}
	dm.createConnection()
	return dm
}
