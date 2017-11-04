package biding

import "clickyab.com/exchange/octopus/exchange"

type bid struct {
	bid     exchange.Bid
	price   int64
	markup  string
	winurl  string
	billurl string
}

func (b bid) ImpID() string {
	return b.bid.ImpID()
}

func (b bid) Demand() exchange.Demand {
	return b.bid.Demand()
}

func (b bid) ID() string {
	return b.bid.ID()
}

func (b bid) Price() int64 {
	return b.price
}

func (b bid) AdID() string {
	return b.bid.AdID()
}

func (b bid) AdHeight() int {
	return b.bid.AdHeight()
}

func (b bid) AdWidth() int {
	return b.bid.AdWidth()
}

func (b bid) AdMarkup() string {
	return b.markup
}

func (b bid) AdDomains() []string {
	return b.bid.AdDomains()
}

func (b bid) WinURL() string {
	return b.winurl
}

func (b bid) BillURL() string {
	return b.billurl
}

func (b bid) Categories() []string {
	return b.bid.Categories()
}

func (b bid) Attributes() map[string]interface{} {
	return b.bid.Attributes()
}

type rsp struct {
	id         string
	bids       []exchange.Bid
	excuse     int
	attributes map[string]interface{}
	supplier   exchange.Supplier
}

func (r rsp) ID() string {
	return r.id
}

func (r rsp) Bids() []exchange.Bid {
	return r.bids
}

func (r rsp) Excuse() int {
	return r.excuse
}

func (r rsp) Attributes() map[string]interface{} {
	return r.attributes
}

func (r rsp) Supplier() exchange.Supplier {
	return r.supplier
}
