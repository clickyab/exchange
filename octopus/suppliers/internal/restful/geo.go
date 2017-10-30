package restful

import (
	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/srtb"
)

type geo struct {
	inner *srtb.Geo
	ip string
}

func (g geo) LatLon() exchange.LatLon {
	return g.ILatLon
}

func (g geo) Country() exchange.Country {
	return g.ICountry
}

func (g geo) Region() exchange.Region {
	return g.IRegion
}

func (g geo) ISP() exchange.ISP {
	return g.IIsp
}
