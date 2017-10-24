package rtbobj

import (
	"strings"

	"reflect"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type device struct {
	inner *openrtb.Device
	geo   exchange.Geo
}

func (d *device) UserAgent() string {
	return d.inner.UA
}

func (d *device) Geo() exchange.Geo {
	if d.geo == nil {
		d.geo = &geo{inner: d.inner.Geo, ip: d.inner.IP}
	}
	return d.geo
}

func (d *device) IP() string {
	return d.inner.IP
}

func (d *device) DeviceType() exchange.DeviceType {
	return exchange.DeviceType(d.inner.DeviceType)
}

func (d *device) Make() string {
	return d.inner.Make
}

func (d *device) Model() string {
	return d.inner.Model
}

func (d *device) OS() string {
	return d.inner.OS
}

func (d *device) Language() string {
	return d.inner.Language
}

func (d *device) Carrier() string {
	return d.inner.Carrier
}

func (d *device) MCC() string {
	if x := strings.Split(d.inner.MCCMNC, "-"); len(x) == 2 {
		return x[0]
	}
	return ""
}

func (d *device) MNC() string {
	if x := strings.Split(d.inner.MCCMNC, "-"); len(x) == 2 {
		return x[1]
	}
	return ""
}

func (d *device) ConnType() exchange.ConnectionType {
	return exchange.ConnectionType(d.inner.ConnType)
}

func (d *device) LAC() string {
	return reflect.StructTag(d.inner.Ext).Get("lac")

}

func (d *device) CID() string {
	return reflect.StructTag(d.inner.Ext).Get("cid")

}

func (d *device) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"IPv6":       d.inner.IPv6,
		"DeviceType": d.inner.DeviceType,
		"Make":       d.inner.Make,
		"Model":      d.inner.Model,
		"OS":         d.inner.OS,
		"Language":   d.inner.Language,
		"Carrier":    d.inner.Carrier,
		"ConnType":   d.inner.ConnType,
		"Ext":        d.inner.Ext,
	}
}
