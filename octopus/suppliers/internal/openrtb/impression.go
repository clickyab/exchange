package openrtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
	"github.com/sirupsen/logrus"
)

type impression struct {
	openrtb.Impression
	attributes map[string]interface{}
}

func (imp impression) FID() string {
	return imp.ID
}

func (imp impression) FBidFloor() float64 {
	return imp.BidFloor
}

func (imp impression) Attributes() map[string]interface{} {
	return imp.attributes
}

func (imp impression) Type() exchange.ImpressionType {
	if imp.Banner != nil {
		return exchange.AdTypeBanner
	} else if imp.Native != nil {
		return exchange.AdTypeNative
	} else if imp.Audio != nil {
		logrus.Panicln("audio ad not supported")
	}

	return exchange.AdTypeVideo
}

func (imp impression) FSecure() bool {
	return int(imp.Secure) == 1

}
