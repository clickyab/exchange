package srtb

import "clickyab.com/exchange/octopus/srtb"

type publisher struct {
	inner *srtb.Publisher
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
	return make(map[string]interface{})
}
