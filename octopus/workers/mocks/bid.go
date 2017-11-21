package mocks

import "clickyab.com/exchange/octopus/exchange"

type Bid struct {
	IPrice    float64
	IID       string
	IImpID    string
	IAdID     string
	IAdHeight int
	IAdWidth  int
	IAttr     map[string]interface{}
	IDemand   Demands
	IWinURL   string
	IBillURL  string
	ICat      []string
	IDomains  []string
}

func (b Bid) ID() string {
	return b.IID
}

func (b Bid) ImpID() string {
	return b.IImpID
}

func (b Bid) Price() float64 {
	return b.IPrice
}

func (b Bid) AdID() string {
	return b.IAdID
}

func (b Bid) AdHeight() int {
	return b.IAdHeight
}

func (b Bid) AdWidth() int {
	return b.IAdWidth
}

func (b Bid) AdMarkup() string {
	panic("implement me")
}

func (b Bid) AdDomains() []string {

	return b.IDomains
}

func (b Bid) WinURL() string {
	return b.IWinURL
}

func (b Bid) Categories() []string {
	return b.ICat
}

func (b Bid) Attributes() map[string]interface{} {
	return b.IAttr
}

func (b Bid) Demand() exchange.Demand {
	return b.IDemand
}

func (b Bid) BillURL() string {
	return b.IBillURL
}
