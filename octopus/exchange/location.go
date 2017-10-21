package exchange

// Country is the country object
type Country struct {
	Valid bool   `json:"valid"`
	Name  string `json:"name"`
	// Country using ISO 3166-1 Alpha 3
	ISO string `json:"iso"`
}

// Region of the request
type Region struct {
	Valid bool   `json:"valid"`
	Name  string `json:"name"`
	// Region using ISO 3166-2
	ISO string `json:"iso"`
}

// LatLon is the latitude longitude
type LatLon struct {
	Valid bool    `json:"valid"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}

// ISP network holder
type ISP struct {
	Valid bool   `json:"valid"`
	Name  string `json:"name"`
}

// Location is the location provider
// Design notes:
// XXX : Type, Accuracy,LastFix,IPService,RegionFIPS104,Metro, City, Zip, UTCOffset   are not supported
// TODO : Support for city
type Location interface {
	// The latitude and longitude of device
	LatLon() LatLon
	// The country of the device
	Country() Country
	// Region of the device
	Region() Region
	ISP() ISP
	// Attributes return all unused fields from open rtb geo
	Attributes() map[string]interface{}
}

// Geo is the alias of Location for open rtb compatibility
type Geo Location
