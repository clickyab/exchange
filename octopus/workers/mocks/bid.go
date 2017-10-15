package mocks

import "clickyab.com/exchange/octopus/exchange"

type Bid struct {
	IPrice    int64
	IID       string
	IImpID    string
	IAdID     string
	IAdHeight int
	IAdWidth  int
	IAttr     map[string]interface{}
	IDemand   Demands
}

func (b Bid) ID() string {
	return b.IID
}

func (b Bid) ImpID() string {
	return b.IImpID
}

func (b Bid) Price() int64 {
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
	panic("implement me")
}

func (b Bid) WinURL() string {
	panic("implement me")
}

func (b Bid) Categories() []string {
	panic("implement me")
}

func (b Bid) Attributes() map[string]interface{} {
	return b.IAttr
}

func (b Bid) Win() {
	panic("implement me")
}

func (b Bid) Demand() exchange.Demand {
	return b.IDemand
}