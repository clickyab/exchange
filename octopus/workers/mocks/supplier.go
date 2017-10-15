package mocks

import (
	"clickyab.com/exchange/octopus/exchange"
)

type Supplier struct {
	SName         string
	SFloorCPM     int64
	SSoftFloorCPM int64
	SShare        int
	SRenderer     exchange.Renderer
}

func (s Supplier) Name() string {
	return s.SName
}

func (s Supplier) FloorCPM() int64 {
	return s.SFloorCPM
}

func (s Supplier) SoftFloorCPM() int64 {
	return s.SSoftFloorCPM
}

func (s Supplier) ExcludedDemands() []string {
	return []string{}
}

func (s Supplier) Share() int {
	return s.SShare
}

func (s Supplier) Renderer() exchange.Renderer {
	return s.SRenderer
}

func (s Supplier) TestMode() bool {
	panic("implement me")
}

func (s Supplier) ClickMode() exchange.SupplierClickMode {
	panic("implement me")
}

func (s Supplier) Type() string {
	panic("implement me")
}
