package exchange

// Video video part of bid request
type Video interface {
	// Width is the width of the ad (max width bypassed cuz of being deprecated)
	Width() int
	// Height is the height of the ad (max height bypassed cuz of being deprecated)
	Height() int
	// Mimes returns the mime range of a banner
	Mimes() []string
	// Linearity shows whether video is linear or not
	Linearity() bool
	// BlockedAttributes returns blocked creative attributes
	BlockedAttributes() []CreativeAttribute
	// Attributes returns ext and other useless stuff from request as a map
	Attributes() map[string]interface{}
}
