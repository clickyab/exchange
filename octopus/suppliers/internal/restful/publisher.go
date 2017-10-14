package restful

import "clickyab.com/exchange/octopus/exchange"

type restPublisher struct {
	PubName         string `json:"name"`
	PubFloorCPM     int64  `json:"floor_cpm"`
	PubSoftFloorCPM int64  `json:"soft_floor_cpm"`

	sup   exchange.Supplier
}

func (rp *restPublisher) ID() string {
	panic("implement me")
}

func (rp *restPublisher) Domain() {
	panic("implement me")
}

func (rp *restPublisher) Cat() []exchange.Category {
	panic("implement me")
}

func (rp *restPublisher) Publisher() exchange.Publisher {
	panic("implement me")
}

func (rp restPublisher) Name() string {
	return rp.PubName
}

func (rp restPublisher) FloorCPM() int64 {
	return rp.PubFloorCPM
}

func (rp restPublisher) SoftFloorCPM() int64 {
	return rp.PubSoftFloorCPM
}

func (restPublisher) Attributes() map[string]interface{} {
	return nil
}

func (rp restPublisher) Supplier() exchange.Supplier {
	return rp.sup
}
