package srtb

import (
	"clickyab.com/exchange/octopus/exchange"
)

type geo struct {
	ip      string
	latlon  exchange.LatLon
	country exchange.Country
	region  exchange.Region
	isp     exchange.ISP
}

func (g *geo) LatLon() exchange.LatLon {
	return g.latlon
}

func (g *geo) Country() exchange.Country {
	return g.country
}

func (g *geo) Region() exchange.Region {
	return g.region
}

func (g *geo) ISP() exchange.ISP {
	return g.isp
}

func (g *geo) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}
