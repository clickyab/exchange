package restful

import (
	"clickyab.com/exchange/octopus/exchange"
)

type Device struct {
	IUA         string                  `json:"ua"`
	IIP         string                  `json:"ip"`
	IGeo        Geo                     `json:"-"`
	IDeviceType exchange.DeviceType     `json:"device_type"`
	IMake       string                  `json:"make"`
	IModel      string                  `json:"model"`
	IConnType   exchange.ConnectionType `json:"conn_type"`
	ICarrier    string                  `json:"carrier"`
	IOs         string                  `json:"os"`
	ILang       string                  `json:"lang"`
	ILAC        string                  `json:"lac"`
	IMNC        string                  `json:"mnc"`
	IMCC        string                  `json:"mcc"`
	ICID        string                  `json:"cid"`
}

func (d Device) UserAgent() string {
	return d.IUA
}

func (d Device) Geo() exchange.Geo {
	return d.IGeo
}

func (d Device) IP() string {
	return d.IIP
}

func (d Device) DeviceType() exchange.DeviceType {
	return d.IDeviceType
}

func (d Device) Make() string {
	return d.IMake
}

func (d Device) Model() string {
	return d.IModel
}

func (d Device) OS() string {
	return d.IOs
}

func (d Device) Language() string {
	return d.ILang
}

func (d Device) Carrier() string {
	return d.ILang
}

func (d Device) MCC() string {
	return d.IMCC
}

func (d Device) MNC() string {
	return d.IMNC
}

func (d Device) ConnType() exchange.ConnectionType {
	return d.IConnType
}

func (d Device) LAC() string {
	return d.ILAC
}

func (d Device) CID() string {
	return d.ICID
}
