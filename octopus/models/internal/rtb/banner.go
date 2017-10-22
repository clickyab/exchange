package rtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type banner struct {
	inner *openrtb.Banner
	bType []exchange.BannerType
	bAttr []exchange.CreativeAttribute
}

func (b *banner) ID() string {
	return b.inner.ID
}

func (b *banner) Width() int {
	return b.inner.W
}

func (b *banner) Height() int {
	return b.inner.H
}

func (b *banner) BlockedTypes() []exchange.BannerType {
	if b.bType == nil {
		for _, v := range b.inner.BType {
			b.bType = append(b.bType, exchange.BannerType(v))
		}
	}
	return b.bType
}

func (b *banner) BlockedAttributes() []exchange.CreativeAttribute {
	if b.bAttr == nil {
		for _, v := range b.inner.BAttr {
			b.bAttr = append(b.bAttr, exchange.CreativeAttribute(v))
		}
	}
	return b.bAttr
}

func (b *banner) Mimes() []string {
	return b.inner.Mimes
}

func (b *banner) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"WMax":     b.inner.WMax,
		"HMax":     b.inner.HMax,
		"WMin":     b.inner.WMin,
		"HMin":     b.inner.HMin,
		"BType":    b.inner.BType,
		"BAttr":    b.inner.BAttr,
		"Pos":      b.inner.Pos,
		"TopFrame": b.inner.TopFrame,
		"ExpDir":   b.inner.ExpDir,
		"Api":      b.inner.Api,
		"Ext":      b.inner.Ext,
	}
}
