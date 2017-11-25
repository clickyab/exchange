package srtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/simple-rtb"
)

type publisher struct {
	inner     *srtb.Publisher
	inventory exchange.Inventory
}

func (p *publisher) ID() string {
	if p.inner == nil {
		return p.inventory.ID()
	}
	return p.inner.ID
}

func (p *publisher) Name() string {
	if p.inner == nil {
		return p.inventory.Name()
	}
	return p.inner.Name
}

func (p *publisher) Cat() []string {
	if p.inner == nil {
		return []string{}
	}
	return p.inner.Cat
}

func (p *publisher) Domain() string {
	if p.inner == nil {
		return p.inventory.Domain()
	}
	return p.inner.Domain
}

func (p *publisher) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}
