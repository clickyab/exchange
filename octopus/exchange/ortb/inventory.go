package ortb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type App struct {
	inner *openrtb.App
	sup   exchange.Supplier
	pub   exchange.Publisher
}

func (n *App) FloorCPM() int64 {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.sup.FloorCPM()
}

func (n *App) SoftFloorCPM() int64 {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.sup.SoftFloorCPM()
}

func (n *App) ID() string {
	return n.inner.ID
}

func (n *App) Name() string {
	return n.inner.Name
}

func (n *App) Domain() string {
	return n.inner.Domain
}

func (n *App) Cat() []exchange.Category {
	r := make([]exchange.Category, 0)
	for _, v := range n.inner.Cat {
		r = append(r, exchange.Category(v))
	}
	return r
}

func (n *App) Publisher() exchange.Publisher {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.pub
}

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

func (n *App) Supplier() exchange.Supplier {
	return n.sup
}

func (n *App) Bundle() string {
	return n.inner.Bundle
}

type Site struct {
	inner *openrtb.Site
	sup   exchange.Supplier
	pub   exchange.Publisher
}

func (n *Site) FloorCPM() int64 {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.sup.FloorCPM()
}

func (n *Site) SoftFloorCPM() int64 {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.sup.SoftFloorCPM()
}

func (n *Site) ID() string {
	return n.inner.ID
}

func (n *Site) Name() string {
	return n.inner.Name
}

func (n *Site) Domain() string {
	return n.inner.Domain
}

func (n *Site) Cat() []exchange.Category {
	r := make([]exchange.Category, 0)
	for _, v := range n.inner.Cat {
		r = append(r, exchange.Category(v))
	}
	return r
}

func (n *Site) Publisher() exchange.Publisher {
	if n.pub == nil {
		n.pub = &Publisher{inner: n.inner.Publisher}
	}
	return n.pub
}

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

func (n *Site) Supplier() exchange.Supplier {
	return n.sup
}

func (n *Site) Page() string {
	return n.inner.Page
}

func (n *Site) Ref() string {
	return n.inner.Ref
}
