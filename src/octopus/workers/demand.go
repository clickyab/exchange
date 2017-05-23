package workers

import (
	"time"

	"services/assert"
	"services/broker"
	"services/safe"
)

type Ad struct {
	Winner_cpm int64 `json:"winner_cpm"`
}

type demandModel struct {
	Impression struct {
		Source struct {
			Name     string `json:"name"`
			Supplier struct {
				Name string `json:"name"`
			} `json:"supplier"`
		} `json:"source"`
		Time time.Time `json:"time"`
	} `json:"impression"`
	Demand struct {
		Name string `json:"name"`
	} `json:"dem"`
	Ads map[string]Ad `json:"ads"`
}

// winnerConsumer asd
type demandConsumer struct {
}

func (*demandConsumer) Topic() string {
	return "winner"
}

func (*demandConsumer) Queue() string {
	return "winner_que"
}

func (s *demandConsumer) Consume() chan<- broker.Delivery {
	chn := make(chan broker.Delivery, 0)
	safe.GoRoutine(func() { s.fillChan(chn) })
	return chn

}
func (s *demandConsumer) fillChan(chn chan broker.Delivery) {
	for {
		select {
		case del := <-chn:
			obj := demandModel{}
			err := del.Decode(&obj)
			assert.Nil(err)
			dataChannel <- tableModel{
				Supplier:     obj.Impression.Source.Supplier.Name,
				Source:       obj.Impression.Source.Name,
				Time:         factTableID(obj.Impression.Time),
				Demand:       obj.Demand.Name,
				Win:          len(obj.Ads),
				Acknowledger: &del.(Acknowledger),
			}
		}
	}
}
