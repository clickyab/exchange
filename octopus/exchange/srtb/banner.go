package srtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/simple-rtb"
)

type banner struct {
	inner *srtb.Banner
}

func (b *banner) ID() string {
	return b.inner.ID
}

func (b *banner) Width() int {
	return b.inner.Width
}

func (b *banner) Height() int {
	return b.inner.Height
}

func (b *banner) BlockedTypes() []exchange.BannerType {
	return make([]exchange.BannerType, 0)
}

func (b *banner) BlockedAttributes() []exchange.CreativeAttribute {
	return make([]exchange.CreativeAttribute, 0)
}

func (b *banner) Mimes() []string {
	return []string{}
}

func (b *banner) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}
