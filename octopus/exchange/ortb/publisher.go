package ortb

import "github.com/bsm/openrtb"

// Publisher ortb
type Publisher struct {
	inner *openrtb.Publisher
}

// ID return ortb ID
func (p *Publisher) ID() string {
	return p.inner.ID
}

// Name return ortb Name
func (p *Publisher) Name() string {
	return p.inner.Name
}

// Cat return ortb Cat
func (p *Publisher) Cat() []string {
	return p.inner.Cat
}

// Domain return ortb Domain
func (p *Publisher) Domain() string {
	return p.inner.Domain
}

// Attributes return ortb Attributes
func (p *Publisher) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Ext": p.inner.Ext,
	}
}
