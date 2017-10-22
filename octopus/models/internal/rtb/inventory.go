package rtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type app struct {
	inner *openrtb.App
	sup   *exchange.Supplier
	pub   exchange.Publisher
}

func (n *app) FloorCPM() int64 {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher}
	}
	return (*n.sup).FloorCPM()
}

func (n *app) SoftFloorCPM() int64 {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher}
	}
	return (*n.sup).SoftFloorCPM()
}

func (n *app) ID() string {
	return n.inner.ID
}

func (n *app) Name() string {
	return n.inner.Name
}

func (n *app) Domain() string {
	return n.inner.Domain
}

func (n *app) Cat() []exchange.Category {
	r := make([]exchange.Category, 0)
	for _, v := range n.inner.Cat {
		r = append(r, exchange.Category(v))
	}
	return r
}

func (n *app) Publisher() exchange.Publisher {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher}
	}
	return n.pub
}

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

func (n *app) Supplier() exchange.Supplier {
	return *n.sup
}

func (n *app) Bundle() string {
	return n.inner.Bundle
}

type site struct {
	inner *openrtb.Site
	sup   *exchange.Supplier
	pub   exchange.Publisher
}

func (n *site) FloorCPM() int64 {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher}
	}
	return (*n.sup).FloorCPM()
}

func (n *site) SoftFloorCPM() int64 {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher}
	}
	return (*n.sup).SoftFloorCPM()
}

func (n *site) ID() string {
	return n.inner.ID
}

func (n *site) Name() string {
	return n.inner.Name
}

func (n *site) Domain() string {
	return n.inner.Domain
}

func (n *site) Cat() []exchange.Category {
	r := make([]exchange.Category, 0)
	for _, v := range n.inner.Cat {
		r = append(r, exchange.Category(v))
	}
	return r
}

func (n *site) Publisher() exchange.Publisher {
	if n.pub == nil {
		n.pub = &publisher{inner: n.inner.Publisher}
	}
	return n.pub
}

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

func (n *site) Supplier() exchange.Supplier {
	return *n.sup
}

func (n *site) Page() string {
	return n.inner.Page
}

func (n *site) Ref() string {
	return n.inner.Ref
}
