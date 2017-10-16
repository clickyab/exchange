package mocks

import "clickyab.com/exchange/octopus/exchange"

type Device struct {
	IUA       string
	IGeo      exchange.Geo
	IIP       string
	ILocation Location
	IDType exchange.DeviceType
}

func (d *Device) UserAgent() string {
	return d.IUA
}

func (d *Device) Geo() exchange.Geo {
	return d.ILocation
}

func (d *Device) IP() string {
	return d.IIP
}

func (d *Device) DeviceType() exchange.DeviceType {
	return d.IDType
}

func (d *Device) Make() string {
	panic("implement me")
}

func (d *Device) Model() string {
	panic("implement me")
}

func (d *Device) OS() string {
	panic("implement me")
}

func (d *Device) Language() string {
	panic("implement me")
}

func (d *Device) Carrier() string {
	panic("implement me")
}

func (d *Device) MCC() string {
	panic("implement me")
}

func (d *Device) MNC() string {
	panic("implement me")
}

func (d *Device) ConnType() exchange.ConnectionType {
	panic("implement me")
}

func (d *Device) LAC() string {
	panic("implement me")
}

func (d *Device) CID() string {
	panic("implement me")
}
