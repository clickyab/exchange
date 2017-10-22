package restful

import "clickyab.com/exchange/octopus/exchange"

type Banner struct {
	IID   string                 `json:"id"`
	IW    int                    `json:"w"`
	IH    int                    `json:"h"`
	IAttr map[string]interface{} `json:"attr"`
}

func (b Banner) ID() string {
	return b.IID
}

func (b Banner) Width() int {
	return b.IW
}

func (b Banner) Height() int {
	return b.IH
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
	return b.IAttr
}
