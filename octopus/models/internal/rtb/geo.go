package rtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/ip2location"
)

type geo struct {
	inner *openrtb.Geo
	ip    string
}

func (g *geo) LatLon() exchange.LatLon {
	t := ip2location.GetAll(g.ip)
	return exchange.LatLon{
		Valid: true,
		Lon:   float64(t.Longitude),
		Lat:   float64(t.Latitude),
	}
}

func (g *geo) Country() exchange.Country {
	t := ip2location.GetAll(g.ip)
	return exchange.Country{
		Valid: t.CountryLong != "",
		Name:  t.CountryLong,
		ISO:   t.CountryShort,
	}
}

func (g *geo) Region() exchange.Region {
	t := ip2location.GetAll(g.ip)
	return exchange.Region{
		Valid: t.Region != "",
		Name:  t.Region,
		ISO:   t.Region,
	}
}

func (g *geo) ISP() exchange.ISP {
	t := ip2location.GetAll(g.ip)
	return exchange.ISP{
		Name:  t.Isp,
		Valid: t.Isp != "",
	}
}

func (g *geo) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Type":          g.inner.Type,
		"Accuracy":      g.inner.Accuracy,
		"LastFix":       g.inner.LastFix,
		"IPService":     g.inner.IPService,
		"RegionFIPS104": g.inner.RegionFIPS104,
		"Metro":         g.inner.Metro,
		"City":          g.inner.City,
		"Zip":           g.inner.Zip,
		"UTCOffset":     g.inner.UTCOffset,
		"Ext":           g.inner.Ext,
	}
}
