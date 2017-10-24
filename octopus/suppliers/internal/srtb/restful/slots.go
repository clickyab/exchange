package restful

import (
	"clickyab.com/exchange/octopus/exchange"
)

type impRest struct {
	W           int               `json:"width"`
	H           int               `json:"height"`
	TID         string            `json:"track_id"`
	FallbackURL string            `json:"fallback_url"`
	FAttribute  map[string]string `json:"attributes"`
}

func (sr *impRest) Attributes() map[string]interface{} {
	panic("implement me")
}

func (sr *impRest) ID() string {
	panic("implement me")
}

func (sr *impRest) BidFloor() float64 {
	panic("implement me")
}

func (sr *impRest) Banner() exchange.Banner {
	panic("implement me")
}

func (sr *impRest) Video() exchange.Video {
	panic("implement me")
}

func (sr *impRest) Native() exchange.Native {
	panic("implement me")
}

func (sr *impRest) Type() exchange.ImpressionType {
	panic("implement me")
}

func (sr *impRest) Secure() bool {
	panic("implement me")
}

func (sr impRest) Fallback() string {
	return sr.FallbackURL
}

func (sr impRest) Width() int {
	return sr.W
}

func (sr impRest) Height() int {
	return sr.H
}

func (sr impRest) TrackID() string {
	return sr.TID
}

func (sr *impRest) SetAttribute(att string, v string) {
	if sr.FAttribute == nil {
		sr.FAttribute = make(map[string]string)
	}
	sr.FAttribute[att] = v
}
