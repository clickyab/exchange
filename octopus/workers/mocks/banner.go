package mocks

import "clickyab.com/exchange/octopus/exchange"

type Banner struct {
	SID string
	SWidth int
	SHeight int
	SBlockedTypes []exchange.BannerType
	SBlockedAttr []exchange.CreativeAttribute
	SMimes []string
	SAttributes map[string]interface{}
}

func (b *Banner) ID() string {
	return b.SID
}

func (b *Banner) Width() int {
	return b.SWidth
}

func (b *Banner) Height() int {
	return b.SHeight
}

func (b *Banner) BlockedTypes() []exchange.BannerType {
	return b.SBlockedTypes
}

func (b *Banner) BlockedAttributes() []exchange.CreativeAttribute {
	return b.SBlockedAttr
}

func (b *Banner) Mimes() []string {
	return b.SMimes
}

func (b *Banner) Attributes() map[string]interface{} {
	return b.SAttributes
}

