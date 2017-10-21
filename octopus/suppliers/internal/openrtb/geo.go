package openrtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/ip2location"
)

type geo struct {
	country    exchange.Country
	region     exchange.Region
	isp        exchange.ISP
	latLon     exchange.LatLon
	attributes map[string]interface{}
}

func (g geo) Attributes() map[string]interface{} {
	return g.attributes
}

func (g geo) LatLon() exchange.LatLon {
	return g.latLon
}

func (g geo) Country() exchange.Country {
	return g.country
}

func (g geo) Region() exchange.Region {
	return g.region
}

func (g geo) ISP() exchange.ISP {
	return g.isp
}

func geoAttributes(r *openrtb.Geo) map[string]interface{} {
	return map[string]interface{}{
		"Type":          r.Type,
		"Accuracy":      r.Accuracy,
		"LastFix":       r.LastFix,
		"IPService":     r.IPService,
		"RegionFIPS104": r.RegionFIPS104,
		"Metro":         r.Metro,
		"City":          r.City,
		"Zip":           r.Zip,
		"UTCOffset":     r.UTCOffset,
		"Ext":           r.Ext,
	}
}

func newGeo(ip string, g *openrtb.Geo) exchange.Geo {
	x := ip2location.GetAll(ip)
	return geo{
		attributes: geoAttributes(g),
		country: exchange.Country{
			Valid: true,
			Name:  x.CountryLong,
			ISO:   x.CountryShort,
		},
		isp: exchange.ISP{
			Valid: x.Isp != "",
			Name:  x.Isp,
		},
		latLon: exchange.LatLon{
			Valid: true,
			Lat:   float64(x.Latitude),
			Lon:   float64(x.Longitude),
		},
		region: exchange.Region{
			Valid: x.Region != "",
			Name:  x.Region,
		},
	}
}
