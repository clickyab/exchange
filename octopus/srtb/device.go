package srtb

// Device shows the clients device details
type Device struct {
	UA         string `json:"ua"`
	IP         string `json:"ip"`
	DeviceType int    `json:"device_type"`
	Make       string `json:"make"`
	Model      string `json:"model"`
	ConnType   int    `json:"conn_type"`
	Carrier    string `json:"carrier"`
	Os         string `json:"os"`
	Lang       string `json:"lang"`
	LAC        string `json:"lac"`
	MNC        string `json:"mnc"`
	MCC        string `json:"mcc"`
	CID        string `json:"cid"`
}
