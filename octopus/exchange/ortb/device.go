package ortb

import (
	"strings"

	"reflect"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type Device struct {
	inner *openrtb.Device
	geo   exchange.Geo
}

func (d *Device) UserAgent() string {
	return d.inner.UA
}

func (d *Device) Geo() exchange.Geo {
	if d.geo == nil {
		d.geo = &Geo{inner: d.inner.Geo, ip: d.inner.IP}
	}
	return d.geo
}

func (d *Device) IP() string {
	return d.inner.IP
}

func (d *Device) DeviceType() exchange.DeviceType {
	return exchange.DeviceType(d.inner.DeviceType)
}

func (d *Device) Make() string {
	return d.inner.Make
}

func (d *Device) Model() string {
	return d.inner.Model
}

func (d *Device) OS() string {
	return d.inner.OS
}

func (d *Device) Language() string {
	return d.inner.Language
}

func (d *Device) Carrier() string {
	return d.inner.Carrier
}

func (d *Device) MCC() string {
	if x := strings.Split(d.inner.MCCMNC, "-"); len(x) == 2 {
		return x[0]
	}
	return ""
}

func (d *Device) MNC() string {
	if x := strings.Split(d.inner.MCCMNC, "-"); len(x) == 2 {
		return x[1]
	}
	return ""
}

func (d *Device) ConnType() exchange.ConnectionType {
	return exchange.ConnectionType(d.inner.ConnType)
}

func (d *Device) LAC() string {
	return reflect.StructTag(d.inner.Ext).Get("lac")

}

func (d *Device) CID() string {
	return reflect.StructTag(d.inner.Ext).Get("cid")

}

func (d *Device) Attributes() map[string]interface{} {
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
