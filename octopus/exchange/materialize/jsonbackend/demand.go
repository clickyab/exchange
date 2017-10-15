package jsonbackend

import (
	"encoding/json"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/broker"

	"github.com/sirupsen/logrus"
)

type demand struct {
	dmn  exchange.Demand
	resp exchange.BidResponse

	src []byte
}

// Encode encode
func (d demand) Encode() ([]byte, error) {
	if d.src == nil {
		themap := make(map[string]interface{})
		themap["bids"] = bidsToMap(d.resp)
		d.src, _ = json.Marshal(themap)
	}

	return d.src, nil
}

// Length return length
func (d demand) Length() int {
	x, _ := d.Encode()
	return len(x)
}

// Topic return topic
func (d demand) Topic() string {
	return "demand"
}

// Key return key
func (d demand) Key() string {
	return ""
}

// Report report
func (d demand) Report() func(error) {
	return func(err error) {
		if err != nil {
			logrus.Warn(err)
		}
	}
}

// DemandJob returns a job for demand
// TODO : add a duration to this. for better view this is important
func DemandJob(resp exchange.BidResponse) broker.Job {
	return &demand{
		resp: resp,
	}
}
