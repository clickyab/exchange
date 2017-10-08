package exchange

type Banner interface {
	// ID represent unique id for banner
	ID() string
	// Width is the width of the ad (max width bypassed cuz of being deprecated)
	Width() int
	// Height is the height of the ad (max height bypassed cuz of being deprecated)
	Height() int
	// Type of the banner (e.g BannerTypeXHTMLText, BannerTypeXHTML, BannerTypeJS, BannerTypeFrame)
	Type() []int
	// BlockedTypes returns blocked creative types
	BlockedTypes() []CreativeType
	// BlockedAttributes returns blocked creative attributes
	BlockedAttributes() []CreativeAttribute
	// Mimes returns the mime range of a banner
	Mimes() []string
	// Attributes returns ext and other useless stuff from request as a map
	Attributes() map[string]interface{}
}
