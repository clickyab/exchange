package mocks

import "clickyab.com/exchange/octopus/exchange"

type Location struct {
	LCountry  exchange.Country
	LProvince exchange.Region
	LLatLon   exchange.LatLon
	LISP      exchange.ISP
}

func (l Location) Region() exchange.Region {
	panic("implement me")
}

func (l Location) ISP() exchange.ISP {
	return l.LISP
}

func (l Location) Country() exchange.Country {
	return l.LCountry
}

func (l Location) Province() exchange.Region {
	return l.LProvince
}

func (l Location) LatLon() exchange.LatLon {
	return l.LLatLon
}
