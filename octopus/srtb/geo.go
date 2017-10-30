package srtb

// Geo consist of location based data of a client
type Geo struct {
	LatLon  LatLon  `json:"lat_lon"`
	Country Country `json:"country"`
	Region  Region  `json:"region"`
	Isp     string  `json:"isp"`
}

// LatLon consist of lat and lon of the client device
type LatLon struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Country is the country data of client
type Country struct {
	Name string `json:"name"`
	ISO  string `json:"iso"`
}

// Region is the region part of client's geo
type Region struct {
	Name string `json:"name"`
	ISO  string `json:"iso"`
}
