package srtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/srtb"
	"github.com/clickyab/services/random"
)

type impression struct {
	inner  *srtb.Impression
	banner exchange.Banner
	cid    string
}

func (i *impression) CID() string {
	if i.cid == "" {
		i.cid = <-random.ID
	}
	return i.cid
}

func (i *impression) ID() string {
	return i.inner.ID
}

func (i *impression) BidFloor() float64 {
	return i.inner.BidFloor
}

func (i *impression) Banner() exchange.Banner {
	if i.Type() != exchange.AdTypeBanner {
		return nil
	}
	if i.banner == nil {
		i.banner = &banner{inner: i.inner.Banner}
	}
	return i.banner
}

func (i *impression) Video() exchange.Video {
	panic("implement me")
}

func (i *impression) Native() exchange.Native {
	panic("implement me")
}

func (i *impression) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}

func (i *impression) Type() exchange.ImpressionType {
	if i.inner.Banner != nil {
		return exchange.AdTypeBanner
	}
	panic("not valid ad type only banner supported")
}

func (i *impression) Secure() bool {
	if i.inner.Secure == 1 {
		return true
	}
	return false
}
