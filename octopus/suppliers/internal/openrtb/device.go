package openrtb

import "clickyab.com/exchange/octopus/exchange"

type device struct {
	IUA         string                  `json:"ua"`
	IIP         string                  `json:"ip"`
	IGeo        geo                     `json:"geo"`
	IDeviceType exchange.DeviceType     `json:"device_type"`
	IMake       string                  `json:"make"`
	IModel      string                  `json:"model"`
	IOs         string                  `json:"os"`
	ILanguage   string                  `json:"language"`
	ICarrier    string                  `json:"carrier"`
	IMNC        string                  `json:"mnc"`
	IMCC        string                  `json:"mcc"`
	ILAC        string                  `json:"lac"`
	ICID        string                  `json:"cid"`
	IConnType   exchange.ConnectionType `json:"conn_type"`
}

func (d device) UserAgent() string {
	return d.IUA
}

func (d device) Geo() exchange.Geo {
	return d.IGeo
}

func (d device) IP() string {
	return d.IIP
}

func (d device) DeviceType() exchange.DeviceType {
	return d.IDeviceType
}

func (d device) Make() string {
	return d.IMake
}

func (d device) Model() string {
	return d.IModel
}

func (d device) OS() string {
	return d.IOs
}

func (d device) Language() string {
	return d.ILanguage
}

func (d device) Carrier() string {
	return d.ICarrier
}

func (d device) MCC() string {
	return d.IMCC
}

func (d device) MNC() string {
	return d.IMNC
}

func (d device) ConnType() exchange.ConnectionType {
	return d.IConnType
}

func (d device) LAC() string {
	return d.ILAC
}

func (d device) CID() string {
	return d.ICID
}

//TODO remove just for lint
func init() {
	if false {
		panic(device{})
	}
}
