package srtb

// Device shows the clients device details
type Device struct {
	UA       string `json:"ua"`
	IP       string `json:"ip"`
	ConnType int    `json:"connectiontype,omitempty"`
	Carrier  string `json:"carrier,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LAC      string `json:"lac,omitempty"`
	MNC      string `json:"mnc,omitempty"`
	MCC      string `json:"mcc,omitempty"`
	CID      string `json:"cid,omitempty"`
}