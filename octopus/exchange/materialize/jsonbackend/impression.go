package jsonbackend

import (
	"encoding/json"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/broker"

	"time"

	"github.com/sirupsen/logrus"
)

// ImpressionWorker is model for impression job
type ImpressionWorker struct {
	Publisher string    `json:"source"`
	Supplier  string    `json:"supplier"`
	Imps      []string  `json:"imps"`
	Time      time.Time `json:"time"`
}

type impression struct {
	inner *ImpressionWorker
	key   string
	src   []byte
}

// Encode encode
func (i impression) Encode() ([]byte, error) {
	if i.src == nil {
		i.src, _ = json.Marshal(i.inner)
	}

	return i.src, nil

}

// Length return length
func (i impression) Length() int {
	x, _ := i.Encode()
	return len(x)
}

// Topic return topic
func (i impression) Topic() string {
	return "impression"
}

// Key return key
func (i impression) Key() string {
	return i.key
}

// Report report
func (i impression) Report() func(error) {
	return func(err error) {
		if err != nil {
			logrus.Warn(err)
		}
	}
}

// ImpressionJob return a broker job
func ImpressionJob(req exchange.BidRequest) broker.Job {
	return impression{
		inner: fillImpJob(req),
		key:   req.Device().IP(),
	}
}

// try to fill imp prepare for job
func fillImpJob(req exchange.BidRequest) *ImpressionWorker {
	var impRes []string
	for i := range req.Imp() {
		impRes = append(impRes, req.Imp()[i].ID())
	}
	return &ImpressionWorker{
		Time:      req.Time(),
		Supplier:  req.Inventory().Supplier().Name(),
		Publisher: req.Inventory().Domain(),
		Imps:      impRes,
	}
}
