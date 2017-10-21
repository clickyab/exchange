package restful

import "clickyab.com/exchange/octopus/exchange"

type publisher struct {
	FID     string   `json:"id"`
	FName   string   `json:"name"`
	FCat    []exchange.Category `json:"cat"`
	FDomain string   `json:"domain"`
}

func (p publisher) ID() string {
	return p.FID
}

func (p publisher) Name() string {
	return p.FName
}

func (p publisher) Cat() []exchange.Category {
	return p.FCat
}

func (p publisher) Domain() string {
	return p.FDomain
}
