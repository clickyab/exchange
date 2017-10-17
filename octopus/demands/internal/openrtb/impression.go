package openrtb

import "clickyab.com/exchange/octopus/exchange"

type impression struct {
	IID        string                  `json:"id"`
	IBidFloor  float64                 `json:"bid_floor"`
	IBanner    exchange.Banner         `json:"banner"`
	IVideo     exchange.Video          `json:"video"`
	INative    exchange.Native         `json:"native"`
	IType      exchange.ImpressionType `json:"type"`
	ISecure    bool                    `json:"secure"`
	attributes map[string]interface{}
}

func (i impression) ID() string {
	return i.IID
}

func (i impression) BidFloor() float64 {
	return i.IBidFloor
}

func (i impression) Banner() exchange.Banner {
	return i.IBanner
}

func (i impression) Video() exchange.Video {
	return i.IVideo
}

func (i impression) Native() exchange.Native {
	return i.INative
}

func (i impression) Attributes() map[string]interface{} {
	return i.attributes
}

func (i impression) Type() exchange.ImpressionType {
	return i.IType
}

func (i impression) Secure() bool {
	return i.ISecure
}

// lint bypass
func init() {
	if false {
		panic(impression{})
	}
}
