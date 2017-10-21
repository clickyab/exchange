package restful

import (
	"clickyab.com/exchange/octopus/exchange"
)

type device struct {
	FUA         string                  `json:"ua"`
	FIP         string                  `json:"ip"`
	FGeo        geo                     `json:"geo,omitempty"`
	FDeviceType exchange.DeviceType     `json:"device_type"`
	FMake       string                  `json:"make"`
	FModel      string                  `json:"model"`
	FConnType   exchange.ConnectionType `json:"conn_type"`
	FCarrier    string                  `json:"carrier"`
	FOs         string                  `json:"os"`
	FLang       string                  `json:"lang"`
}

func (d device) UserAgent() string {
	return d.FUA
}

func (d device) Geo() exchange.Geo {
	return d.FGeo
}

func (d device) IP() string {
	return d.FIP
}

func (d device) DeviceType() exchange.DeviceType {
	return d.FDeviceType
}

func (d device) Make() string {
	return d.FMake
}

func (d device) Model() string {
	return d.FModel
}

func (d device) OS() string {
	return d.FOs
}

func (d device) Language() string {
	panic("implement me")
}

func (d device) Carrier() string {
	return d.FLang
}

func (d device) MCC() string {
	return ""
}

func (d device) MNC() string {
	return ""
}

func (d device) ConnType() exchange.ConnectionType {
	return d.FConnType
}

func (d device) LAC() string {
	return ""
}

func (d device) CID() string {
	return ""
}
