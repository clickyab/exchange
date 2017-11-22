package jsonbackend

import (
	"encoding/json"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/broker"

	"time"

	"github.com/sirupsen/logrus"
)

// Demand demand struct
type Demand struct {
	Supplier   string    `json:"supplier"`
	Source     string    `json:"source"`
	Demand     string    `json:"demand"`
	Time       time.Time `json:"time"`
	TotalPrice float64   `json:"total_price"`
	BidLen     int       `json:"bid_len"`

	src []byte
}

// Encode encode the data
func (d Demand) Encode() ([]byte, error) {
	if len(d.src) == 0 {
		var err error
		d.src, err = json.Marshal(d)
		if err != nil {
			return nil, err
		}
	}
	return d.src, nil
}

// Length get length
func (d Demand) Length() int {
	if len(d.src) == 0 {
		d.Encode()
	}
	return len(d.src)
}

// Topic return topic
func (Demand) Topic() string {
	return "demand"
}

// Key return key
func (Demand) Key() string {
	return "demand_aggregate"
}

// Report stuff
func (Demand) Report() func(error) {
	return func(err error) {
		if err != nil {
			logrus.Warn(err)
		}
	}
}

// DemandJob returns a job for Demand
// TODO : add a duration to this. for better view this is important
func DemandJob(rq exchange.BidRequest, resp exchange.BidResponse, demand string) broker.Job {
	var total float64
	for _, i := range resp.Bids() {
		total += i.Price()
	}
	return &Demand{
		Source:     rq.Inventory().Domain(),
		Supplier:   rq.Inventory().Supplier().Name(),
		Demand:     demand,
		BidLen:     len(resp.Bids()),
		TotalPrice: total,
	}
}
