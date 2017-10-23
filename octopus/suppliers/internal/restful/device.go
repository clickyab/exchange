package restful

import (
	"encoding/json"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/srtb"
)

type Device struct {
	inner *srtb.Device
	geo   geo
}

func (d *Device) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.inner)
}

func (d *Device) UnmarshalJSON(a []byte) error {
	i := srtb.Device{}
	err := json.Unmarshal(a, &i)
	if err != nil {
		return err
	}

	//TODO check validation
	//validate device

	d.inner = &i
	return nil
}

func (d Device) UserAgent() string {
	return d.inner.UA
}

func (d Device) Geo() exchange.Geo {
	return d.geo
}

func (d Device) IP() string {
	return d.inner.IP
}

func (d Device) DeviceType() exchange.DeviceType {
	return exchange.DeviceType(d.inner.DeviceType)
}

func (d Device) Make() string {
	return d.inner.Make
}

func (d Device) Model() string {
	return d.inner.Model
}

func (d Device) OS() string {
	return d.inner.Os
}

func (d Device) Language() string {
	return d.inner.Lang
}

func (d Device) Carrier() string {
	return d.inner.Lang
}

func (d Device) MCC() string {
	return d.inner.MCC
}

func (d Device) MNC() string {
	return d.inner.MNC
}

func (d Device) ConnType() exchange.ConnectionType {
	return exchange.ConnectionType(d.inner.ConnType)
}

func (d Device) LAC() string {
	return d.inner.LAC
}

func (d Device) CID() string {
	return d.inner.CID
}
