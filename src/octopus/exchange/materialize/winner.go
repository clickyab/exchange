package materialize

import (
	"encoding/json"
	"octopus/exchange"
	"services/broker"
)

type winner struct {
	data map[string]interface{}
	key  string

	src []byte
}

// Encode encode
func (w *winner) Encode() ([]byte, error) {
	if w.src == nil {
		w.src, _ = json.Marshal(w.data)
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
	return func(error) {}
}

// WinnerJob return a broker job
func WinnerJob(imp exchange.Impression, dmn exchange.Demand, ad exchange.Advertise, slotID string) broker.Job {
	return &winner{
		data: winnerToMap(imp, dmn, ad, slotID),
		key:  imp.IP().String(),
	}
}
