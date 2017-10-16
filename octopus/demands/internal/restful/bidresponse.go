package restful

import "clickyab.com/exchange/octopus/exchange"

type bidResponse struct {
	FBids []bid  `json:"bids"`
	FID   string `json:"id"`
}

func (b bidResponse) ID() string {
	return b.FID
}

func (b bidResponse) Bids() []exchange.Bid {
	res := make([]exchange.Bid, 0)
	for _, v := range b.FBids {
		res = append(res, v)
	}
	return res
}

func (bidResponse) Excuse() int {
	panic("implement me")
}

func (bidResponse) Attributes() map[string]interface{} {
	panic("implement me")
}

func (bidResponse) Supplier() exchange.Supplier {
	panic("implement me")
}

type bid struct {
	FID     string `json:"id"`
	FDemand exchange.Demand
}

func (b bid) ID() string {
	return b.FID
}

func (bid) ImpID() string {
	panic("implement me")
}

func (bid) Price() int64 {
	panic("implement me")
}

func (bid) AdID() string {
	panic("implement me")
}

func (bid) AdHeight() int {
	panic("implement me")
}

func (bid) AdWidth() int {
	panic("implement me")
}

func (bid) AdMarkup() string {
	panic("implement me")
}

func (bid) AdDomains() []string {
	panic("implement me")
}

func (bid) WinURL() string {
	panic("implement me")
}

func (bid) Categories() []string {
	panic("implement me")
}

func (bid) Attributes() map[string]interface{} {
	panic("implement me")
}

func (bid) Win() {
	panic("implement me")
}

func (bid) Demand() exchange.Demand {
	panic("implement me")
}
