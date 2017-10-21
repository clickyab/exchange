package openrtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type banner struct {
	id           string
	width        int
	height       int
	blockedTypes []exchange.BannerType
	blockedAttrs []exchange.CreativeAttribute
	mimes        []string
	attributes   map[string]interface{}
}

func (b banner) ID() string {
	return b.id
}

func (b banner) Width() int {
	return b.width
}

func (b banner) Height() int {
	return b.height
}

func (b banner) BlockedTypes() []exchange.BannerType {
	return b.blockedTypes
}

func (b banner) BlockedAttributes() []exchange.CreativeAttribute {
	return b.blockedAttrs
}

func (b banner) Mimes() []string {
	return b.mimes
}

func (b banner) Attributes() map[string]interface{} {
	return b.attributes
}

func newBanner(b *openrtb.Banner) exchange.Banner {
	bt := func(x []int) []exchange.BannerType {
		r := make([]exchange.BannerType, 0)
		for _, v := range x {
			r = append(r, exchange.BannerType(v))
		}
		return r
	}(b.BType)
	ba := func(x []int) []exchange.CreativeAttribute {
		r := make([]exchange.CreativeAttribute, 0)
		for _, v := range x {
			r = append(r, exchange.CreativeAttribute(v))
		}
		return r
	}(b.BAttr)
	return banner{
		id:           b.ID,
		width:        b.W,
		height:       b.H,
		blockedTypes: bt,
		blockedAttrs: ba,
		mimes:        b.Mimes,
		attributes:   bannerAttributes(b),
	}
}

func bannerAttributes(b *openrtb.Banner) map[string]interface{} {
	return map[string]interface{}{
		"WMax":     b.WMax,
		"HMax":     b.HMax,
		"WMin":     b.WMin,
		"HMin":     b.HMin,
		"BType":    b.BType,
		"BAttr":    b.BAttr,
		"Pos":      b.Pos,
		"TopFrame": b.TopFrame,
		"ExpDir":   b.ExpDir,
		"Api":      b.Api,
		"Ext":      b.Ext,
	}
}
