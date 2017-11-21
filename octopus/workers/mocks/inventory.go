package mocks

import "clickyab.com/exchange/octopus/exchange"

type Inventory struct {
	IID            string
	IName, IDomain string
	ISupplier      Supplier
	IAttr          map[string]interface{}
	IFloorCPM      float64
	ISoftFloorCPM  float64
}

func (i Inventory) ID() string {
	return i.IID
}

func (i Inventory) Name() string {
	return i.IName
}

func (i Inventory) Domain() string {
	return i.IDomain
}

func (Inventory) Cat() []exchange.Category {
	panic("implement me")
}

func (Inventory) Publisher() exchange.Publisher {
	panic("implement me")
}

func (i Inventory) Attributes() map[string]interface{} {
	return i.IAttr
}

func (i Inventory) FloorCPM() float64 {
	return i.IFloorCPM
}

func (i Inventory) SoftFloorCPM() float64 {
	return i.ISoftFloorCPM
}

func (i Inventory) Supplier() exchange.Supplier {
	return &i.ISupplier
}
