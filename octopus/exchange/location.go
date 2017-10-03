package exchange

// Country is the country object
type Country struct {
	Valid bool   `json:"valid"`
	Name  string `json:"name"`
	// Country using ISO 3166-1 Alpha 3
	ISO string `json:"iso"`
}

// Province of the request
type Province struct {
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
	Region() Province
}

// Geo is the alias of Location for open rtb compatibility
type Geo = Location
