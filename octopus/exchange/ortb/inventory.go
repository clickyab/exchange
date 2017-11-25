package ortb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

// app ortb structure
type app struct {
	inner *openrtb.App
	sup   exchange.Supplier
	pub   exchange.Publisher
}

// FloorCPM return ortb FloorCPM
func (n *app) FloorCPM() float64 {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher}
	}
	return n.sup.FloorCPM()
}

// SoftFloorCPM return ortb SoftFloorCPM
func (n *app) SoftFloorCPM() float64 {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher}
	}
	return n.sup.SoftFloorCPM()
}

// ID return ortb ID
func (n *app) ID() string {
	return n.inner.ID
}

// Name return ortb Name
func (n *app) Name() string {
	return n.inner.Name
}

// Domain return ortb Domain
func (n *app) Domain() string {
	return n.inner.Domain
}

// Cat return ortb Cat
func (n *app) Cat() []exchange.Category {
	r := make([]exchange.Category, 0)
	for _, v := range n.inner.Cat {
		r = append(r, exchange.Category(v))
	}
	return r
}

// publisher return ortb publisher
func (n *app) Publisher() exchange.Publisher {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher, inventory: n}
	}
	return n.pub
}

// Attributes return ortb Attributes
func (n *app) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"SectionCat":    n.inner.SectionCat,
		"PageCat":       n.inner.PageCat,
		"PrivacyPolicy": n.inner.PrivacyPolicy,
		"Content":       n.inner.Content,
		"Keywords":      n.inner.Keywords,
		"Ext":           n.inner.Ext,
		"StoreURL":      n.inner.StoreURL,
		"Ver":           n.inner.Ver,
		"Paid":          n.inner.Paid,
	}
}

// Supplier return ortb Supplier
func (n *app) Supplier() exchange.Supplier {
	return n.sup
}

// Bundle return ortb Bundle
func (n *app) Bundle() string {
	return n.inner.Bundle
}

// site ortb site
type site struct {
	inner *openrtb.Site
	sup   exchange.Supplier
	pub   exchange.Publisher
}

// FloorCPM return ortb FloorCPM
func (n *site) FloorCPM() float64 {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher}
	}
	return n.sup.FloorCPM()
}

// SoftFloorCPM return ortb SoftFloorCPM
func (n *site) SoftFloorCPM() float64 {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher}
	}
	return n.sup.SoftFloorCPM()
}

// ID return ortb ID
func (n *site) ID() string {
	return n.inner.ID
}

// Name return ortb Name
func (n *site) Name() string {
	return n.inner.Name
}

// Domain return ortb Domain
func (n *site) Domain() string {
	return n.inner.Domain
}

// Cat return ortb Cat
func (n *site) Cat() []exchange.Category {
	r := make([]exchange.Category, 0)
	for _, v := range n.inner.Cat {
		r = append(r, exchange.Category(v))
	}
	return r
}

// publisher return ortb publisher
func (n *site) Publisher() exchange.Publisher {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher, inventory: n}
	}
	return n.pub
}

// Attributes return ortb Attributes
func (n *site) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Cat":           n.inner.Cat,
		"SectionCat":    n.inner.SectionCat,
		"PageCat":       n.inner.PageCat,
		"PrivacyPolicy": n.inner.PrivacyPolicy,
		"Content":       n.inner.Content,
		"Keywords":      n.inner.Keywords,
		"Ext":           n.inner.Ext,
		"Search":        n.inner.Search,
		"Mobile":        n.inner.Mobile,
	}
}

// Supplier return ortb Supplier
func (n *site) Supplier() exchange.Supplier {
	return n.sup
}

// Page return ortb Page
func (n *site) Page() string {
	return n.inner.Page
}

// Ref return ortb Ref
func (n *site) Ref() string {
	return n.inner.Ref
}
