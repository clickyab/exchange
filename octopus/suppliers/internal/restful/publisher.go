package restful

import "clickyab.com/exchange/octopus/exchange"

type publisher struct {
	IID     string              `json:"id"`
	IName   string              `json:"name"`
	ICat    []exchange.Category `json:"cat"`
	IDomain string              `json:"domain"`
}

func (p publisher) ID() string {
	return p.IID
}

func (p publisher) Name() string {
	return p.IName
}

func (p publisher) Cat() []exchange.Category {
	return p.ICat
}

func (p publisher) Domain() string {
	return p.IDomain
}
