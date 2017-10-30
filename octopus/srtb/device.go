package srtb

import "clickyab.com/exchange/octopus/exchange"

// Device shows the clients device details
type Device struct {
	UA       string `json:"ua"`
	IP       string `json:"ip"`
	Geo      Geo    `json:"geo,omitempty"`
	ConnType int    `json:"connectiontype,omitempty"`
	Carrier  string `json:"carrier,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LAC      string `json:"lac,omitempty"`
	MNC      string `json:"mnc,omitempty"`
	MCC      string `json:"mcc,omitempty"`
	CID      string `json:"cid,omitempty"`
}

// Geo srtb geo type
type Geo struct {
	LatLon  exchange.LatLon  `json:"latlon,omitempty"`
	Region  exchange.Region  `json:"region,omitempty"`
	Country exchange.Country `json:"country,omitempty"`
	ISP     exchange.ISP     `json:"isp,omitempty"`
}
