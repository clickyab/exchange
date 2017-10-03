package exchange

// Impression is the slot of the app
type Impression interface {
	// Size return the primary size of this slot
	Width() int
	Height() int
	// ID is an string for this slot, its a random at first but the value is not changed at all other calls
	TrackID() string
	// Fallback returns slots fallback url
	Fallback() string
	Attributes() map[string]string
	// SetAttribute is a way to set attributes on the slot
	// TODO : this is a workaround for click url backup. think it again.
	SetAttribute(string, string)
}
