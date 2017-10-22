package restful

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/random"
)

type Imp struct {
	IID       string                  `json:"id"`
	clid      string                  `json:"-"`
	IBanner   *Banner                 `json:"Banner"`
	IBidFloor float64                 `json:"bid_floor"`
	IAttr     map[string]interface{}  `json:"attr"`
	ISecure   int                     `json:"secure"`
	IType     exchange.ImpressionType `json:"-"`
}

func (i Imp) CID() string {
	if i.clid == "" {
		return <-random.ID
	}
	return i.clid
}

func (i Imp) ID() string {
	return i.IID
}

func (i Imp) BidFloor() float64 {
	return i.IBidFloor
}

func (i Imp) Banner() exchange.Banner {
	return i.IBanner
}

func (i Imp) Video() exchange.Video {
	return nil
}

func (i Imp) Native() exchange.Native {
	return nil
}

func (i Imp) Attributes() map[string]interface{} {
	return i.IAttr
}

func (i Imp) Type() exchange.ImpressionType {
	return i.IType
}

func (i Imp) Secure() bool {
	if i.ISecure == 1 {
		return true
	}
	return false
}
