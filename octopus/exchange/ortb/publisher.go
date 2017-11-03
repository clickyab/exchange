package ortb

import "github.com/bsm/openrtb"

// publisher ortb
type publisher struct {
	inner *openrtb.Publisher
}

// ID return ortb ID
func (p *publisher) ID() string {
	return p.inner.ID
}

// Name return ortb Name
func (p *publisher) Name() string {
	return p.inner.Name
}

// Cat return ortb Cat
func (p *publisher) Cat() []string {
	return p.inner.Cat
}

// Domain return ortb Domain
func (p *publisher) Domain() string {
	return p.inner.Domain
}

// Attributes return ortb Attributes
func (p *publisher) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Ext": p.inner.Ext,
	}
}
