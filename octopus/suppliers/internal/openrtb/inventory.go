package openrtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type inventory struct {
	id                     string
	name                   string
	domain                 string
	cat                    []exchange.Category
	publisher              exchange.Publisher
	attributes             map[string]interface{}
	supplier               exchange.Supplier
	floorCPM, softFloorCPM int64
}

func (v inventory) FloorCPM() int64 {
	return v.floorCPM
}

func (v inventory) SoftFloorCPM() int64 {
	return v.softFloorCPM
}

func (v inventory) ID() string {
	return v.id
}

func (v inventory) Name() string {
	return v.name
}

func (v inventory) Domain() string {
	return v.domain
}

func (v inventory) Cat() []exchange.Category {
	return v.cat
}

func (v inventory) Publisher() exchange.Publisher {
	return v.publisher
}

func (v inventory) Attributes() map[string]interface{} {
	return v.attributes
}

func (v inventory) Supplier() exchange.Supplier {
	return v.supplier
}

type app struct {
	inventory
	bundle string
}

func (a app) Bundle() string {
	return a.bundle
}

type site struct {
	inventory
	page string
	ref  string
}

func (s site) Page() string {
	return s.page
}

func (s site) Ref() string {
	return s.ref
}

func newInventory(r *openrtb.BidRequest, supplier exchange.Supplier) exchange.Inventory {
	catter := func(m []string) []exchange.Category {
		res := make([]exchange.Category, 0)
		for _, l := range m {
			res = append(res, exchange.Category(l))
		}
		return res
	}
	if r.Site != nil {
		return site{
			inventory: inventory{
				id:           r.Site.ID,
				attributes:   siteAttributes(r.Site),
				name:         r.Site.Name,
				cat:          catter(r.Site.Cat),
				publisher:    newPublisher(r.Site.Publisher),
				domain:       r.Site.Domain,
				supplier:     supplier,
				floorCPM:     supplier.FloorCPM(),
				softFloorCPM: supplier.SoftFloorCPM(),
			},
			ref:  r.Site.Ref,
			page: r.Site.Page,
		}
	} else if r.App != nil {
		return app{
			inventory: inventory{
				id:           r.App.ID,
				attributes:   appAttributes(r.App),
				name:         r.App.Name,
				cat:          catter(r.App.Cat),
				publisher:    newPublisher(r.App.Publisher),
				domain:       r.App.Domain,
				supplier:     supplier,
				floorCPM:     supplier.FloorCPM(),
				softFloorCPM: supplier.SoftFloorCPM(),
			},
			bundle: r.App.Bundle,
		}
	}
	panic("not a valid inventory")

}

func siteAttributes(s *openrtb.Site) map[string]interface{} {
	return map[string]interface{}{
		"Cat":           s.Cat,
		"SectionCat":    s.SectionCat,
		"PageCat":       s.PageCat,
		"PrivacyPolicy": s.PrivacyPolicy,
		"Content":       s.Content,
		"Keywords":      s.Keywords,
		"Ext":           s.Ext,
		"Search":        s.Search,
		"Mobile":        s.Mobile,
	}
}

func appAttributes(a *openrtb.App) map[string]interface{} {
	return map[string]interface{}{
		"SectionCat":    a.SectionCat,
		"PageCat":       a.PageCat,
		"PrivacyPolicy": a.PrivacyPolicy,
		"Content":       a.Content,
		"Keywords":      a.Keywords,
		"Ext":           a.Ext,
		"StoreURL":      a.StoreURL,
		"Ver":           a.Ver,
		"Paid":          a.Paid,
	}
}
