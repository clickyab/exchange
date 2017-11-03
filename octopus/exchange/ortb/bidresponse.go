package ortb

import (
	"encoding/json"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

// NewBidResponse new bid response
func NewBidResponse(d exchange.Demand, s exchange.Supplier) exchange.BidResponse {
	return &bidResponse{
		demand:   d,
		supplier: s,
	}
}

type bidResponse struct {
	inner    *openrtb.BidResponse
	bids     []exchange.Bid
	demand   exchange.Demand
	supplier exchange.Supplier
}

func (b *bidResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.inner)
}

func (b *bidResponse) UnmarshalJSON(d []byte) error {
	o := &openrtb.BidResponse{}
	err := json.Unmarshal(d, o)
	if err != nil {
		return err
	}
	err = o.Validate()
	if err != nil {
		return err
	}
	b.inner = o
	return nil
}

func (b *bidResponse) ID() string {
	return b.inner.ID
}

func (b *bidResponse) Bids() []exchange.Bid {
	if b.bids == nil {
		for _, v := range b.inner.SeatBid {
			for _, x := range v.Bid {
				b.bids = append(b.bids, &bid{
					inner:  &x,
					demand: b.demand,
				})
			}
		}
	}
	return b.bids
}

func (b *bidResponse) Excuse() int {
	return b.Excuse()
}

func (b *bidResponse) Attributes() map[string]interface{} {
	return map[string]interface{}{}
}

func (b *bidResponse) Supplier() exchange.Supplier {
	return b.supplier
}

type bid struct {
	inner  *openrtb.Bid
	demand exchange.Demand
}

func (b *bid) ID() string {
	return b.inner.ID
}

func (b *bid) ImpID() string {
	return b.inner.ImpID
}

func (b *bid) Price() int64 {
	return int64(b.inner.Price)
}

func (b *bid) AdID() string {
	return b.inner.AdID
}

func (b *bid) AdHeight() int {
	return b.inner.H
}

func (b *bid) AdWidth() int {
	return b.inner.W
}

func (b *bid) AdMarkup() string {
	return b.inner.AdMarkup
}

func (b *bid) AdDomains() []string {
	return b.inner.AdvDomain
}

func (b *bid) WinURL() string {
	return b.inner.NURL
}

func (b *bid) Categories() []string {
	return b.inner.Cat
}

func (b *bid) Attributes() map[string]interface{} {
	return map[string]interface{}{}
}

func (b *bid) Demand() exchange.Demand {
	return b.demand
}
