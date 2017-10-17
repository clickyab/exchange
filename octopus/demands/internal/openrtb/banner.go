package openrtb

import "clickyab.com/exchange/octopus/exchange"

type banner struct {
	IID           string                       `json:"id"`
	IWidth        int                          `json:"width"`
	IHeight       int                          `json:"height"`
	IBlockedTypes []exchange.BannerType        `json:"blocked_types"`
	BlockedAttrs  []exchange.CreativeAttribute `json:"blocked_attrs"`
	IMimes        []string                     `json:"mimes"`
	attributes    map[string]interface{}
}

func (b banner) ID() string {
	return b.IID
}

func (b banner) Width() int {
	return b.IWidth
}

func (b banner) Height() int {
	return b.IHeight
}

func (b banner) BlockedTypes() []exchange.BannerType {
	return b.IBlockedTypes
}

func (b banner) BlockedAttributes() []exchange.CreativeAttribute {
	return b.BlockedAttrs
}

func (b banner) Mimes() []string {
	return b.IMimes
}

func (b banner) Attributes() map[string]interface{} {
	return b.attributes
}

//lint bypass
func init() {
	if false {
		panic(banner{})
	}
}
