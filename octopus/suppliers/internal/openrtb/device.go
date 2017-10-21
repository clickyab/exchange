package openrtb

import (
	"strings"

	"reflect"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type device struct {
	ua         string
	ip         string
	geo        exchange.Geo
	deviceType exchange.DeviceType
	make       string
	model      string
	os         string
	language   string
	carrier    string
	mnc        string
	mcc        string
	lac        string
	cid        string
	connType   exchange.ConnectionType
	attributes map[string]interface{}
}

func (d device) Attributes() map[string]interface{} {
	return d.attributes
}

func (d device) UserAgent() string {
	return d.ua
}

func (d device) Geo() exchange.Geo {
	return d.geo
}

func (d device) IP() string {
	return d.ip
}

func (d device) DeviceType() exchange.DeviceType {
	return d.deviceType
}

func (d device) Make() string {
	return d.make
}

func (d device) Model() string {
	return d.model
}

func (d device) OS() string {
	return d.os
}

func (d device) Language() string {
	return d.language
}

func (d device) Carrier() string {
	return d.carrier
}

func (d device) MCC() string {
	return d.mcc
}

func (d device) MNC() string {
	return d.mnc
}

func (d device) ConnType() exchange.ConnectionType {
	return d.connType
}

func (d device) LAC() string {
	return d.lac
}

func (d device) CID() string {
	return d.cid
}

func deviceAttributes(d *openrtb.Device) map[string]interface{} {
	return map[string]interface{}{
		"IPv6":       d.IPv6,
		"DeviceType": d.DeviceType,
		"Make":       d.Make,
		"Model":      d.Model,
		"OS":         d.OS,
		"Language":   d.Language,
		"Carrier":    d.Carrier,
		"ConnType":   d.ConnType,
		"Ext":        d.Ext,
	}
}

func newDevice(d *openrtb.Device) exchange.Device {
	mcc, mnc := func() (string, string) {
		t := strings.Split(d.MCCMNC, "-")
		if len(t) == 2 {
			return t[0], t[1]
		}
		return "", ""
	}()

	return device{
		geo:        newGeo(d.IP, d.Geo),
		attributes: deviceAttributes(d),
		ua:         d.UA,
		ip:         d.IP,
		deviceType: exchange.DeviceType(d.DeviceType),
		make:       d.Make,
		model:      d.Model,
		os:         d.OS,
		language:   d.Language,
		carrier:    d.Carrier,
		mnc:        mnc,
		mcc:        mcc,
		connType:   exchange.ConnectionType(d.ConnType),
		lac:        reflect.StructTag(d.Ext).Get("lac"),
		cid:        reflect.StructTag(d.Ext).Get("cid"),
	}
}
