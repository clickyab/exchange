package rtb

import "github.com/bsm/openrtb"

type publisher struct {
	inner *openrtb.Publisher
}

func (p *publisher) ID() string {
	return p.inner.ID
}

func (p *publisher) Name() string {
	return p.inner.Name
}

func (p *publisher) Cat() []string {
	return p.inner.Cat
}

func (p *publisher) Domain() string {
	return p.inner.Domain
}

func (p *publisher) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Ext": p.inner.Ext,
	}
}
