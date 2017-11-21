package jsonbackend

import (
	"encoding/json"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/broker"

	"time"

	"github.com/sirupsen/logrus"
)

// WinnerWorker is model for winner job
type WinnerWorker struct {
	Publisher string    `json:"source"`
	Supplier  string    `json:"supplier"`
	Demand    string    `json:"demand"`
	Time      time.Time `json:"time"`
	WinnerCPM float64   `json:"winner_cpm"`
}

type winner struct {
	inner *WinnerWorker
	key   string

	src []byte
}

// Encode encode
func (w *winner) Encode() ([]byte, error) {
	if w.src == nil {
		w.src, _ = json.Marshal(w.inner)
	}

	return w.src, nil
}

// Length return length
func (w *winner) Length() int {
	x, _ := w.Encode()
	return len(x)
}

// Topic return topic
func (w *winner) Topic() string {
	return "winner"
}

// Key return key
func (w *winner) Key() string {
	return w.key
}

// Report report
func (w *winner) Report() func(error) {
	return func(err error) {
		if err != nil {
			logrus.Warn(err)
		}
	}
}

// WinnerJob return a broker job
func WinnerJob(bq exchange.BidRequest, bid exchange.Bid) broker.Job {
	return &winner{
		inner: fillWinJob(bq, bid),
		key:   bq.Device().IP(),
	}
}

func fillWinJob(bq exchange.BidRequest, bid exchange.Bid) *WinnerWorker {
	return &WinnerWorker{
		Publisher: bq.Inventory().Domain(),
		Supplier:  bq.Inventory().Supplier().Name(),
		Time:      bq.Time(),
		Demand:    bid.Demand().Name(),
		WinnerCPM: bid.Price(),
	}
}
