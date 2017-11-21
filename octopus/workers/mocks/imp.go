package mocks

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/random"
)

type Imp struct {
	SBanner   Banner
	SNative   Native
	SVideo    Video
	SID       string
	SType     exchange.ImpressionType
	SBidFloor float64
	SSecure   bool
	SCid      string
	SCurrency string
	Attribute map[string]interface{}
}

func (i Imp) Currency() string {
	return i.SCurrency
}

func (i Imp) CID() string {
	if i.SCid == "" {
		i.SCid = <-random.ID
	}
	return i.SCid
}

func (i Imp) ID() string {
	return i.SID
}

func (i Imp) BidFloor() float64 {
	return i.SBidFloor
}

func (i Imp) Banner() exchange.Banner {
	return &i.SBanner
}

func (i Imp) Video() exchange.Video {
	return &i.SVideo
}

func (i Imp) Native() exchange.Native {
	return &i.SNative
}

func (i Imp) Attributes() map[string]interface{} {
	return i.Attribute
}

func (i Imp) Type() exchange.ImpressionType {
	return i.SType
}

func (i Imp) Secure() bool {
	return i.SSecure
}
