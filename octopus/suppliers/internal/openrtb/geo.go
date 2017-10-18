package openrtb

import "clickyab.com/exchange/octopus/exchange"

type geo struct {
	ICountry exchange.Country `json:"country"`
	IRegion  exchange.Region  `json:"region"`
	IIsp     exchange.ISP     `json:"isp"`
	ILatLon  exchange.LatLon  `json:"latlon"`
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

//TODO remove just for lint
func init() {
	if false {
		panic(geo{})
	}
}
