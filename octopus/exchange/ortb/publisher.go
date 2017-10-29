package ortb

import "github.com/bsm/openrtb"

type Publisher struct {
	inner *openrtb.Publisher
}

func (p *Publisher) ID() string {
	return p.inner.ID
}

func (p *Publisher) Name() string {
	return p.inner.Name
}

func (p *Publisher) Cat() []string {
	return p.inner.Cat
}

func (p *Publisher) Domain() string {
	return p.inner.Domain
}

func (p *Publisher) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Ext": p.inner.Ext,
	}
}
