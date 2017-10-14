package restful

import "clickyab.com/exchange/octopus/exchange"

type bidResponse struct {
	bids []bid
}

func (*bidResponse) ID() string {
	panic("implement me")
}

func (*bidResponse) Bids() []exchange.Bid {
	panic("implement me")
}

func (*bidResponse) Excuse() int {
	panic("implement me")
}

func (*bidResponse) Attributes() map[string]interface{} {
	panic("implement me")
}

func (*bidResponse) Supplier() exchange.Supplier {
	panic("implement me")
}

type bid struct {
	demand exchange.Demand
}

func (*bid) ID() string {
	panic("implement me")
}

func (*bid) ImpID() string {
	panic("implement me")
}

func (*bid) Price() int64 {
	panic("implement me")
}

func (*bid) AdID() string {
	panic("implement me")
}

func (*bid) AdHeight() int {
	panic("implement me")
}

func (*bid) AdWidth() int {
	panic("implement me")
}

func (*bid) AdMarkup() string {
	panic("implement me")
}

func (*bid) AdDomains() []string {
	panic("implement me")
}

func (*bid) WinURL() string {
	panic("implement me")
}

func (*bid) Categories() []string {
	panic("implement me")
}

func (*bid) Attributes() map[string]interface{} {
	panic("implement me")
}

func (*bid) Win() {
	panic("implement me")
}

func (*bid) Demand() exchange.Demand {
	panic("implement me")
}
