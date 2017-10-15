package mocks

import "clickyab.com/exchange/octopus/exchange"

type Inventory struct {
	IID string
	IName string
	ISupplier Supplier
	IAttr map[string]interface{}
	IFloorCPM int64
	ISoftFloorCPM int64
}

func (i Inventory) ID() string {
	return i.IID
}

func (i Inventory) Name() string {
	return i.IName
}

func (Inventory) Domain() string {
	panic("implement me")
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

func (i Inventory) FloorCPM() int64 {
	return i.IFloorCPM
}

func (i Inventory) SoftFloorCPM() int64 {
	return i.ISoftFloorCPM
}

func (i Inventory) Supplier() exchange.Supplier {
	return &i.ISupplier
}

