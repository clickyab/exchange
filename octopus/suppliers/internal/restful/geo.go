package restful

import "clickyab.com/exchange/octopus/exchange"

type geo struct {
	FLatLon  exchange.LatLon  `json:"lat_lon"`
	FCountry exchange.Country `json:"country"`
	FRegion  exchange.Region  `json:"region"`
	FIsp     exchange.ISP     `json:"isp"`
}

func (g geo) LatLon() exchange.LatLon {
	return g.FLatLon
}

func (g geo) Country() exchange.Country {
	return g.FCountry
}

func (g geo) Region() exchange.Region {
	return g.FRegion
}

func (g geo) ISP() exchange.ISP {
	return g.FIsp
}
