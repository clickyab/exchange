package mocks

import (
	"time"

	"clickyab.com/exchange/octopus/exchange"
)

type BidResponse struct {
	IID       string
	IBids     []Bid
	IAttr     map[string]interface{}
	ISupplier Supplier
	ITime     time.Time
}

func (b BidResponse) Time() time.Time {
	return b.ITime
}

func (b BidResponse) ID() string {
	return b.IID
}

func (b BidResponse) Bids() []exchange.Bid {
	res := make([]exchange.Bid, 0)
	for _, val := range b.IBids {
		res = append(res, val)
	}
	return res
}

func (b BidResponse) Excuse() int {
	panic("implement me")
}

func (b BidResponse) Attributes() map[string]interface{} {
	return b.IAttr
}

func (b BidResponse) Supplier() exchange.Supplier {
	return b.ISupplier
}
