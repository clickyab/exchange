package ortb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

// publisher ortb
type publisher struct {
	inner     *openrtb.Publisher
	inventory exchange.Inventory
}

// ID return ortb ID
func (p *publisher) ID() string {
	if p.inner == nil {
		return ""
	}
	return p.inner.ID
}

// Name return ortb Name
func (p *publisher) Name() string {
	if p.inner == nil {
		return p.inventory.Supplier().Name()
	}
	return p.inner.Name
}

// Cat return ortb Cat
func (p *publisher) Cat() []string {
	if p.inner == nil {
		return []string{}
	}
	return p.inner.Cat
}

// Domain return ortb Domain
func (p *publisher) Domain() string {
	if p.inner == nil {
		return p.inventory.Domain()
	}
	return p.inner.Domain
}

// Attributes return ortb Attributes
func (p *publisher) Attributes() map[string]interface{} {
	if p.inner == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"Ext": p.inner.Ext,
	}
}
