package exchange

// Inventory is the publisher interface
// Design notes :
// XXX: We do not support for section cat and page cat.
// XXX: PrivacyPolicy is not supported
// XXX: We do not support content of the inventory
// XXX: We do not support keyword
// Not supporting, means we do not have a dedicated method for them. but
// all of them are available in attributes field and must pass in career
// which support them (i.e open-rtb layer)
type Inventory interface {
	// ID of the inventory in 3rdad
	ID() string
	// Name of publisher
	Name() string
	// Domain the domain of the inventory (domain for site, package name for app)
	Domain() string
	// Cat return the category of this inventory
	Cat() []Category
	// Publisher return the publisher of the inventory or nil when there is no publisher
	// WARN : nil is meaningful. watch it
	Publisher() Publisher

	// Attributes is the generic attribute system. any key that is not supported is dumped here.
	// we do not use them in exchange itself but we pass them into requests
	Attributes() map[string]interface{}

	// Non rtb parameters, use the ext field in open-rtb spec to pass this values

	// FloorCPM is the floor cpm for publisher
	FloorCPM() float64
	// SoftFloorCPM is the soft version of floor cpm. if the publisher ahs it, then the system
	// try to use this as floor, but if this is not available, the FloorCPM is used
	SoftFloorCPM() float64
	// Supplier is for get this inventory supplier
	Supplier() Supplier
}
