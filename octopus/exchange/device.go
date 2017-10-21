package exchange

// DeviceType based on openrtb section 5.21
type DeviceType int

// ConnectionType based on the openrtb 5.22
type ConnectionType int

const (
	// DeviceTypeMobileTablet means mobile tablet
	DeviceTypeMobileTablet DeviceType = iota + 1
	// DeviceTypePC means personal computer
	DeviceTypePC
	// DeviceTypeTV is the connected TV
	DeviceTypeTV
	// DeviceTypePhone is the phone
	DeviceTypePhone
	// DeviceTypeTablet is the tablet, if not sure if this is tablet or not use 1
	DeviceTypeTablet
	// DeviceTypeConnectedDevice any unknown connected device
	DeviceTypeConnectedDevice
	// DeviceTypeSetTopBox SetTopBox
	DeviceTypeSetTopBox
)

const (
	// ConnectionTypeUnknown unknown
	ConnectionTypeUnknown ConnectionType = iota
	// ConnectionTypeEthernet is the ethernet
	ConnectionTypeEthernet
	// ConnectionTypeWIFI is the wifi
	ConnectionTypeWIFI
	// ConnectionTypeCellularUnknown cellular but unknown
	ConnectionTypeCellularUnknown
	// ConnectionTypeCellular2G 2g
	ConnectionTypeCellular2G
	// ConnectionTypeCellular3G 3g
	ConnectionTypeCellular3G
	// ConnectionTypeCellular4G 4g
	ConnectionTypeCellular4G
)

// Device is where the app is shown
// Design note:
// XXX : DNT, LMT,IPv6, OSVer,HwVer, H, W, PPI,PxRatio,JS ,GeoFetch,FlashVer,IFA,IDSHA1,IDMD5,PIDSHA1,PIDMD5,MacSHA1,MacMD5 are not supported
// LAC and CID are clickyab addition
type Device interface {
	// User agent
	UserAgent() string
	// Location of the device assumed to be the user’s current location (open rtb Geo)
	Geo() Geo
	// IP of device (v4)
	IP() string
	// The general type of device
	DeviceType() DeviceType
	// Device make
	Make() string
	// Device model
	Model() string
	// Device OS
	OS() string
	// Browser language
	Language() string
	// Carrier or ISP derived from the IP address
	Carrier() string
	// Mobile carrier as the concatenated MCC-MNC code (e.g., "310-005" identifies Verizon Wireless CDMA in the USA). MCC part
	MCC() string
	// Mobile carrier as the concatenated MCC-MNC code (e.g., "310-005" identifies Verizon Wireless CDMA in the USA). MNC part
	MNC() string
	// Network connection type.
	ConnType() ConnectionType
	// LAC from mobile data
	LAC() string
	// CID from mobile data
	CID() string
	// Attributes return all unused fields from open rtb device
	Attributes() map[string]interface{}
}
