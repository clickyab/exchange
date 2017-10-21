package openrtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type publisher struct {
	id, name, domain string
	cat              []string
	attributes       map[string]interface{}
}

func (p publisher) Attributes() map[string]interface{} {
	return p.attributes
}

func (p publisher) ID() string {
	return p.id
}

func (p publisher) Name() string {
	return p.name
}

func (p publisher) Cat() []string {
	return p.cat
}

func (p publisher) Domain() string {
	return p.domain
}

func newPublisher(p *openrtb.Publisher) exchange.Publisher {
	return publisher{
		attributes: publisherAttributes(p),
		domain:     p.Domain,
		cat:        p.Cat,
		name:       p.Name,
		id:         p.ID,
	}
}

func publisherAttributes(p *openrtb.Publisher) map[string]interface{} {
	return map[string]interface{}{
		"Ext": p.Ext,
	}
}
