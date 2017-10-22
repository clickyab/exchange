package restful

import "clickyab.com/exchange/octopus/exchange"

type Geo struct {
	ILatLon  exchange.LatLon  `json:"lat_lon"`
	ICountry exchange.Country `json:"country"`
	IRegion  exchange.Region  `json:"region"`
	IIsp     exchange.ISP     `json:"isp"`
}

func (g Geo) LatLon() exchange.LatLon {
	return g.ILatLon
}

func (g Geo) Country() exchange.Country {
	return g.ICountry
}

func (g Geo) Region() exchange.Region {
	return g.IRegion
}

func (g Geo) ISP() exchange.ISP {
	return g.IIsp
}
