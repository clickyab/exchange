package restful

import "clickyab.com/exchange/octopus/exchange"

type App struct {
	inventory
}

type Site struct {
	inventory
}

type inventory struct {
	IID       string                 `json:"id"`
	IName     string                 `json:"name,omitempty"`
	IDomain   string                 `json:"domain,omitempty"`
	IPage     string                 `json:"page,omitempty"`
	IRef      string                 `json:"ref,omitempty"`
	ISupplier supplier               `json:"-"`
	ICat      []exchange.Category    `json:"cat"`
	IAttr     map[string]interface{} `json:"attr"`
}

func (inv inventory) ID() string {
	return inv.IID
}

func (inv inventory) Name() string {
	return inv.IName
}

func (inv inventory) Domain() string {
	return inv.IDomain
}

func (inv inventory) Cat() []exchange.Category {
	return inv.ICat
}

func (inv inventory) Publisher() exchange.Publisher {
	var res = publisher{
		IID:     inv.IID,
		ICat:    inv.ICat,
		IDomain: inv.IDomain,
		IName:   inv.IName,
	}
	return res
}

func (inv inventory) Attributes() map[string]interface{} {
	return inv.IAttr
}

func (inv inventory) FloorCPM() int64 {
	return inv.ISupplier.FloorCPM()
}

func (inv inventory) SoftFloorCPM() int64 {
	return inv.ISupplier.SoftFloorCPM()
}

func (inv inventory) Supplier() exchange.Supplier {
	return inv.ISupplier
}
