package ortb

import (
	"strings"

	"reflect"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

// Device ortb device structure
type Device struct {
	inner *openrtb.Device
	geo   exchange.Geo
}

// UserAgent return user agent
func (d *Device) UserAgent() string {
	return d.inner.UA
}

// Geo return geo
func (d *Device) Geo() exchange.Geo {
	if d.geo == nil {
		d.geo = &Geo{inner: d.inner.Geo, ip: d.inner.IP}
	}
	return d.geo
}

// IP return ip
func (d *Device) IP() string {
	return d.inner.IP
}

// DeviceType return device type (tv,pc,...)
func (d *Device) DeviceType() exchange.DeviceType {
	return exchange.DeviceType(d.inner.DeviceType)
}

// Make return manufacturer
func (d *Device) Make() string {
	return d.inner.Make
}

// Model return model
func (d *Device) Model() string {
	return d.inner.Model
}

// OS return os
func (d *Device) OS() string {
	return d.inner.OS
}

// Language return language
func (d *Device) Language() string {
	return d.inner.Language
}

// Carrier return carrier
func (d *Device) Carrier() string {
	return d.inner.Carrier
}

// MCC return mcc
func (d *Device) MCC() string {
	if x := strings.Split(d.inner.MCCMNC, "-"); len(x) == 2 {
		return x[0]
	}
	return ""
}

// MNC return mnc
func (d *Device) MNC() string {
	if x := strings.Split(d.inner.MCCMNC, "-"); len(x) == 2 {
		return x[1]
	}
	return ""
}

// ConnType return connection type
func (d *Device) ConnType() exchange.ConnectionType {
	return exchange.ConnectionType(d.inner.ConnType)
}

// LAC return lac
func (d *Device) LAC() string {
	return reflect.StructTag(d.inner.Ext).Get("lac")

}

// CID return cid
func (d *Device) CID() string {
	return reflect.StructTag(d.inner.Ext).Get("cid")

}

// Attributes return device extra attributes
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
