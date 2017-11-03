package srtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/ip2location"
)

type geo struct {
	ip      string
	latlon  exchange.LatLon
	country exchange.Country
	region  exchange.Region
	isp     exchange.ISP
}

// LatLon return LatLon
func (g *geo) LatLon() exchange.LatLon {
	t := ip2location.GetAll(g.ip)
	return exchange.LatLon{
		Valid: true,
		Lon:   float64(t.Longitude),
		Lat:   float64(t.Latitude),
	}
}

// Country return Country
func (g *geo) Country() exchange.Country {
	t := ip2location.GetAll(g.ip)
	return exchange.Country{
		Valid: t.CountryLong != "",
		Name:  t.CountryLong,
		ISO:   t.CountryShort,
	}
}

// Region return Region
func (g *geo) Region() exchange.Region {
	t := ip2location.GetAll(g.ip)
	return exchange.Region{
		Valid: t.Region != "",
		Name:  t.Region,
		ISO:   t.Region,
	}
}

// ISP return ISP
func (g *geo) ISP() exchange.ISP {
	t := ip2location.GetAll(g.ip)
	return exchange.ISP{
		Name:  t.Isp,
		Valid: t.Isp != "",
	}
}
func (g *geo) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}
