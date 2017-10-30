package restful

import (
	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/srtb"
)

type Banner struct {
	inner *srtb.Banner
}

func (b Banner) ID() string {
	return b.inner.ID
}

func (b Banner) Width() int {
	return b.inner.Width
}

func (b Banner) Height() int {
	return b.inner.Height
}

func (b Banner) BlockedTypes() []exchange.BannerType {
	return make([]exchange.BannerType, 0)
}

func (b Banner) BlockedAttributes() []exchange.CreativeAttribute {
	return make([]exchange.CreativeAttribute, 0)
}

func (b Banner) Mimes() []string {
	return []string{}
}

func (b Banner) Attributes() map[string]interface{} {
	return map[string]interface{}{}
}
