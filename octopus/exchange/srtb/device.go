package srtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/srtb"
	"github.com/clickyab/services/ip2location"
)

type device struct {
	inner *srtb.Device
	geo   exchange.Geo
}

func (d *device) UserAgent() string {
	return d.inner.UA
}

func (d *device) Geo() exchange.Geo {
	if d.geo == nil {
		d.geo = extractGeoFromIP(d.inner.IP)
	}
	return d.geo
}

func (d *device) IP() string {
	return d.inner.IP
}

func (d *device) DeviceType() exchange.DeviceType {
	panic("implement me")
}

func (d *device) Make() string {
	return ""
}

func (d *device) Model() string {
	return ""
}

func (d *device) OS() string {
	return ""
}

func (d *device) Language() string {
	return d.inner.Lang
}

func (d *device) Carrier() string {
	return d.inner.Carrier
}

func (d *device) MCC() string {
	return d.inner.MCC
}

func (d *device) MNC() string {
	return d.inner.MNC
}

func (d *device) ConnType() exchange.ConnectionType {
	return exchange.ConnectionType(d.inner.ConnType)
}

func (d *device) LAC() string {
	return d.inner.LAC
}

func (d *device) CID() string {
	return d.inner.CID
}

func (d *device) Attributes() map[string]interface{} {
	return make(map[string]interface{}, 0)
}

func extractGeoFromIP(ip string) exchange.Geo {
	record := ip2location.GetAll(ip)
	return &geo{
		isp: exchange.ISP{
			Name:  record.Isp,
			Valid: record.Isp != "",
		},
		region: exchange.Region{
			Name:  record.Region,
			Valid: record.Region != "",
		},
		country: exchange.Country{
			Name:  record.CountryLong,
			ISO:   record.CountryShort,
			Valid: record.CountryLong != "",
		},
		latlon: exchange.LatLon{
			Valid: true,
			Lon:   float64(record.Longitude),
			Lat:   float64(record.Latitude),
		},
	}
}
