package ortb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

// App ortb structure
type App struct {
	inner *openrtb.App
	sup   exchange.Supplier
	pub   exchange.Publisher
}

// FloorCPM return ortb FloorCPM
func (n *App) FloorCPM() int64 {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.sup.FloorCPM()
}

// SoftFloorCPM return ortb SoftFloorCPM
func (n *App) SoftFloorCPM() int64 {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.sup.SoftFloorCPM()
}

// ID return ortb ID
func (n *App) ID() string {
	return n.inner.ID
}

// Name return ortb Name
func (n *App) Name() string {
	return n.inner.Name
}

// Domain return ortb Domain
func (n *App) Domain() string {
	return n.inner.Domain
}

// Cat return ortb Cat
func (n *App) Cat() []exchange.Category {
	r := make([]exchange.Category, 0)
	for _, v := range n.inner.Cat {
		r = append(r, exchange.Category(v))
	}
	return r
}

// Publisher return ortb Publisher
func (n *App) Publisher() exchange.Publisher {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.pub
}

// Attributes return ortb Attributes
func (n *App) Attributes() map[string]interface{} {
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
func (n *App) Supplier() exchange.Supplier {
	return n.sup
}

// Bundle return ortb Bundle
func (n *App) Bundle() string {
	return n.inner.Bundle
}

// Site ortb site
type Site struct {
	inner *openrtb.Site
	sup   exchange.Supplier
	pub   exchange.Publisher
}

// FloorCPM return ortb FloorCPM
func (n *Site) FloorCPM() int64 {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.sup.FloorCPM()
}

// SoftFloorCPM return ortb SoftFloorCPM
func (n *Site) SoftFloorCPM() int64 {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.sup.SoftFloorCPM()
}

// ID return ortb ID
func (n *Site) ID() string {
	return n.inner.ID
}

// Name return ortb Name
func (n *Site) Name() string {
	return n.inner.Name
}

// Domain return ortb Domain
func (n *Site) Domain() string {
	return n.inner.Domain
}

// Cat return ortb Cat
func (n *Site) Cat() []exchange.Category {
	r := make([]exchange.Category, 0)
	for _, v := range n.inner.Cat {
		r = append(r, exchange.Category(v))
	}
	return r
}

// Publisher return ortb Publisher
func (n *Site) Publisher() exchange.Publisher {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.pub
}

// Attributes return ortb Attributes
func (n *Site) Attributes() map[string]interface{} {
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
func (n *Site) Supplier() exchange.Supplier {
	return n.sup
}

// Page return ortb Page
func (n *Site) Page() string {
	return n.inner.Page
}

// Ref return ortb Ref
func (n *Site) Ref() string {
	return n.inner.Ref
}
