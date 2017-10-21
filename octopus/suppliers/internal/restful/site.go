package restful

import "clickyab.com/exchange/octopus/exchange"

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
	return i.FSupplier.FPublisher
}

func (i inventory) Attributes() map[string]interface{} {
	return i.FAttr
}

func (i inventory) FloorCPM() int64 {
	return i.FSupplier.FFloorCPM
}

func (i inventory) SoftFloorCPM() int64 {
	return i.FSupplier.FSoftFloorCpm
}

func (i inventory) Supplier() exchange.Supplier {
	return i.FSupplier
}
