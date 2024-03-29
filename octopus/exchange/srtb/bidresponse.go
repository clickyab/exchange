package srtb

import (
	"encoding/json"
	"errors"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/simple-rtb"
)

type bidResponse struct {
	inner    *srtb.BidResponse
	bids     []exchange.Bid
	demand   exchange.Demand
	supplier exchange.Supplier
}

// NewBidResponse generate internal bid-response from srtb
func NewBidResponse(d exchange.Demand, s exchange.Supplier, resp *srtb.BidResponse) exchange.BidResponse {
	return &bidResponse{inner: resp, supplier: s, demand: d}
}

func (b bidResponse) ID() string {
	return b.inner.ID
}

func (b bidResponse) Bids() []exchange.Bid {
	if b.bids == nil {
		for i := range b.inner.Bids {
			b.bids = append(b.bids, &bid{
				inner:  &b.inner.Bids[i],
				demand: b.demand,
			})
		}
	}
	return b.bids
}

func (b bidResponse) Excuse() int {
	return 0
}

func (b bidResponse) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}

func (b bidResponse) Supplier() exchange.Supplier {
	return b.supplier
}

func (b bidResponse) UnmarshalJSON(a []byte) error {
	i := srtb.BidResponse{}
	err := json.Unmarshal(a, &i)
	if err != nil {
		return err
	}
	if i.ID == "" {
		return errors.New("bid response id is required")
	}
	if len(i.Bids) == 0 {
		return errors.New("your bid response has no bids object")
	}

	b.inner = &i
	return nil
}

func (b bidResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

type bid struct {
	inner  *srtb.Bid
	demand exchange.Demand
}

func (b *bid) ID() string {
	return b.inner.ID
}

func (b *bid) ImpID() string {
	return b.inner.ImpID
}

func (b *bid) Price() float64 {
	return b.inner.Price
}

func (b *bid) AdID() string {
	return b.inner.AdID
}

func (b *bid) AdHeight() int {
	return b.inner.Height
}

func (b *bid) AdWidth() int {
	return b.inner.Width
}

func (b *bid) AdMarkup() string {
	return b.inner.AdMarkup
}

func (b *bid) AdDomains() []string {
	return []string{}
}

func (b *bid) WinURL() string {
	return b.inner.WinURL
}

func (b *bid) Categories() []string {
	return b.inner.Cat
}

func (b *bid) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}

func (b *bid) Demand() exchange.Demand {
	return b.demand
}

func (b *bid) BillURL() string {
	return b.inner.BillURL
}
