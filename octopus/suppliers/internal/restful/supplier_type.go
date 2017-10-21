package restful

import "clickyab.com/exchange/octopus/exchange"

type app struct {
	inventory
}

type site struct {
	inventory
}

type inventory struct {
	FID       string                 `json:"id"`
	FName     string                 `json:"name,omitempty"`
	FDomain   string                 `json:"domain"`
	FPage     string                 `json:"page,omitempty"`
	FRef      string                 `json:"ref,omitempty"`
	FSupplier supplier               `json:"supplier"`
	FCat      []exchange.Category    `json:"cat"`
	FAttr     map[string]interface{} `json:"attr"`
}

func (i inventory) ID() string {
	return i.FID
}

func (i inventory) Name() string {
	return i.FName
}

func (i inventory) Domain() string {
	return i.FDomain
}

func (i inventory) Cat() []exchange.Category {
	return i.FCat
}

func (i inventory) Publisher() exchange.Publisher {
	var res = publisher{
		FID:     i.FID,
		FCat:    i.FCat,
		FDomain: i.FDomain,
		FName:   i.FName,
	}
	return res
}

func (i inventory) Attributes() map[string]interface{} {
	return i.FAttr
}

func (i inventory) FloorCPM() int64 {
	return i.FSupplier.FloorCPM()
}

func (i inventory) SoftFloorCPM() int64 {
	return i.FSupplier.SoftFloorCPM()
}

func (i inventory) Supplier() exchange.Supplier {
	return i.FSupplier
}
