package restful

import "entity"

type restPublisher struct {
	PubName         string `json:"name"`
	PubFloorCPM     int64  `json:"floor_cpm"`
	PubSoftFloorCPM int64  `json:"soft_floor_cpm"`

	sup entity.Supplier
}

func (rp restPublisher) Name() string {
	return rp.PubName
}

func (rp restPublisher) FloorCPM() int64 {
	return rp.PubFloorCPM
}

func (rp restPublisher) SoftFloorCPM() int64 {
	return rp.PubFloorCPM
}

func (restPublisher) Attributes(entity.PublisherAttributes) interface{} {
	return nil
}

func (rp restPublisher) Supplier() entity.Supplier {
	return rp.sup
}