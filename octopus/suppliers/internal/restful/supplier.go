package restful

import "clickyab.com/exchange/octopus/exchange"

type supplier struct {
	FName         string    `json:"name"`
	FFloorCPM     int64     `json:"floor_cpm"`
	FSoftFloorCpm int64     `json:"soft_floor_cpm"`
	FPublisher    publisher `json:"publisher"`
	FShare        int       `json:"share"`
	FTest         bool      `json:"-"`
}

func (s supplier) Name() string {
	return s.FName
}

func (s supplier) FloorCPM() int64 {
	return s.FFloorCPM
}

func (s supplier) SoftFloorCPM() int64 {
	return s.FSoftFloorCpm
}

func (s supplier) ExcludedDemands() []string {
	panic("implement me")
}

func (s supplier) Share() int {
	return s.FShare
}

func (s supplier) Renderer() exchange.Renderer {
	panic("implement me")
}

func (s supplier) TestMode() bool {
	return s.FTest
}

func (s supplier) ClickMode() exchange.SupplierClickMode {
	panic("implement me")
}

func (s supplier) Type() string {
	return "rest"
}
