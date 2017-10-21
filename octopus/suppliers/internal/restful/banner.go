package restful

import "clickyab.com/exchange/octopus/exchange"

type banner struct {
	FID       string                 `json:"id"`
	FW        int                    `json:"w"`
	FH        int                    `json:"h"`
	FBidFloor float64                `json:"bid_floor"`
	FSecure   bool                   `json:"secure"`
	FAttr     map[string]interface{} `json:"attr"`
}

func (b banner) ID() string {
	return b.FID
}

func (b banner) Width() int {
	return b.FW
}

func (b banner) Height() int {
	return b.FH
}

func (b banner) BlockedTypes() []exchange.BannerType {
	panic("implement me")
}

func (b banner) BlockedAttributes() []exchange.CreativeAttribute {
	panic("implement me")
}

func (b banner) Mimes() []string {
	panic("implement me")
}

func (b banner) Attributes() map[string]interface{} {
	return b.FAttr
}
