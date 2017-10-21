package restful

import "clickyab.com/exchange/octopus/exchange"

type imp struct {
	FID       string                  `json:"id"`
	FBanner   *banner                 `json:"banner"`
	FBidFloor float64                 `json:"bid_floor"`
	FAttr     map[string]interface{}  `json:"attr"`
	FSecure   bool                    `json:"secure"`
	FType     exchange.ImpressionType `json:"-"`
}

func (i imp) ID() string {
	return i.FID
}

func (i imp) BidFloor() float64 {
	return i.FBidFloor
}

func (i imp) Banner() exchange.Banner {
	return i.FBanner
}

func (i imp) Video() exchange.Video {
	return nil
}

func (i imp) Native() exchange.Native {
	return nil
}

func (i imp) Attributes() map[string]interface{} {
	return i.FAttr
}

func (i imp) Type() exchange.ImpressionType {
	return i.FType
}

func (i imp) Secure() bool {
	return i.FSecure
}
