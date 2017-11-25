package srtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/simple-rtb"
)

type site struct {
	inner *srtb.Site
	sup   exchange.Supplier
	pub   exchange.Publisher
}

func (s *site) ID() string {
	return s.inner.ID
}

func (s *site) Name() string {
	return s.inner.Name
}

func (s *site) Domain() string {
	return s.inner.Domain
}

func (s *site) Cat() []exchange.Category {
	r := make([]exchange.Category, 0)
	for _, i := range s.inner.Cat {
		r = append(r, exchange.Category(i))
	}
	return r
}

func (s *site) Publisher() exchange.Publisher {
	if s.pub == nil {
		s.pub = &publisher{inner: &s.inner.Publisher, inventory: s}
	}
	return s.pub
}

func (s *site) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}

func (s *site) FloorCPM() float64 {
	return s.sup.FloorCPM()
}

func (s *site) SoftFloorCPM() float64 {
	return s.sup.SoftFloorCPM()
}

func (s *site) Supplier() exchange.Supplier {
	return s.sup
}

func (s *site) Page() string {
	return s.inner.Page
}

func (s *site) Ref() string {
	return s.inner.Ref
}

type app struct {
	inner *srtb.App
	sup   exchange.Supplier
	pub   exchange.Publisher
}

func (a *app) ID() string {
	return a.inner.ID
}

func (a *app) Name() string {
	return a.inner.Name
}

func (a *app) Domain() string {
	return a.inner.Domain
}

func (a *app) Cat() []exchange.Category {
	r := make([]exchange.Category, 0)
	for _, v := range a.inner.Cat {
		r = append(r, exchange.Category(v))
	}
	return r
}

func (a *app) Publisher() exchange.Publisher {
	if a.pub == nil {
		a.pub = &publisher{inner: &a.inner.Publisher, inventory: a}
	}
	return a.pub
}

func (a *app) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}

func (a *app) FloorCPM() float64 {
	return a.sup.FloorCPM()
}

func (a *app) SoftFloorCPM() float64 {
	return a.sup.SoftFloorCPM()
}

func (a *app) Supplier() exchange.Supplier {
	return a.sup
}

func (a *app) Bundle() string {
	return a.inner.Bundle
}
