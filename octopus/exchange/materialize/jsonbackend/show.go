package jsonbackend

import (
	"encoding/json"

	"github.com/clickyab/services/broker"

	"github.com/sirupsen/logrus"
)

// ShowWorker struct handle encode
type ShowWorker struct {
	Supplier  string  `json:"supplier"`
	Publisher string  `json:"publisher"`
	Demand    string  `json:"demand"`
	Winner    float64 `json:"winner"`
	Time      string  `json:"time"`
	Profit    float64 `json:"profit"`
}

type show struct {
	inner *ShowWorker
	key   string
	src   []byte
}

// Encode encode
func (s *show) Encode() ([]byte, error) {
	if s.src == nil {
		g, err := json.Marshal(s.inner)
		if err != nil {
			return nil, err
		}
		s.src = g
	}

	return s.src, nil
}

// Length return length
func (s *show) Length() int {
	x, _ := s.Encode()
	return len(x)
}

// Topic return topic
func (*show) Topic() string {
	return "show"
}

// Key return key
func (s *show) Key() string {
	return s.key
}

// Report report
func (*show) Report() func(error) {
	return func(err error) {
		if err != nil {
			logrus.Warn(err)
		}
	}
}

// ShowJob return a broker job
func ShowJob(demand string, IP string, winner float64, t string, supplier string, publisher string, profit float64) broker.Job {
	return &show{
		inner: fillShowJob(demand, winner, t, supplier, publisher, profit),
		key:   IP,
	}
}

func fillShowJob(demand string, winner float64, t string, supplier string, publisher string, profit float64) *ShowWorker {
	return &ShowWorker{
		Publisher: publisher,
		Time:      t,
		Supplier:  supplier,
		Profit:    profit,
		Winner:    winner,
		Demand:    demand,
	}
}
