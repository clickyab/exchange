package restful

import "clickyab.com/exchange/octopus/exchange"

type supplier struct {
	IName         string
	IFloorCPM     int64
	ISoftFloorCpm int64
	IPublisher    publisher
	IShare        int
	ITest         bool
}

func (s supplier) Name() string {
	return s.IName
}

func (s supplier) FloorCPM() int64 {
	return s.IFloorCPM
}

func (s supplier) SoftFloorCPM() int64 {
	return s.ISoftFloorCpm
}

func (s supplier) ExcludedDemands() []string {
	panic("implement me")
}

func (s supplier) Share() int {
	return s.IShare
}

func (s supplier) Renderer() exchange.Renderer {
	panic("implement me")
}

func (s supplier) TestMode() bool {
	return s.ITest
}

func (s supplier) ClickMode() exchange.SupplierClickMode {
	panic("implement me")
}

func (s supplier) Type() string {
	return "rest"
}

// setSupplier profit is here. the floor and soft floor are rising here
func (inv *inventory) setSupplier(sup exchange.Supplier) {

	share := int64(100 + sup.Share())
	floorCPM := (sup.FloorCPM() * share) / 100
	if floorCPM == 0 {
		floorCPM = (sup.FloorCPM() * share) / 100
	}
	softFloorCPM := (sup.SoftFloorCPM() * share) / 100
	if softFloorCPM == 0 {
		softFloorCPM = (sup.SoftFloorCPM() * share) / 100
	}
	inv.ISupplier = supplier{
		ITest:         sup.TestMode(),
		IShare:        int(share),
		IFloorCPM:     floorCPM,
		ISoftFloorCpm: softFloorCPM,
		IName:         sup.Name(),
	}
}
