package entity

// Supplier is the ad-network interface
type Supplier interface {
	// Name of Supplier
	Name() string
	// CPMFloor is the floor for this network. the publisher must be greeter equal to this
	FloorCPM() int64
	// SoftFloorCPM is the soft version of floor cpm. if the publisher ahs it, then the system
	// try to use this as floor, but if this is not available, the FloorCPM is used
	SoftFloorCPM() int64
	// AcceptedTypes is the types that this network can request
	AcceptedTypes() []AdType
	// ExcludedNetwork is the black listed network for this.
	ExcludedDemands() []string
	// CountryWhiteList is the list of country accepted for this
	CountryWhiteList() []Country
	// CallRate is the rate of calling the object, 0 < X <= 100
	CallRate() int
	// Handicap is used to handle the handicap for play in the bid
	Handicap() int
}
