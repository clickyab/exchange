package exchange

// BannerType which contains (XHTML, XHTMLText, TypeJS, TypeFrame)
type BannerType int

const (
	// BannerTypeXHTMLText XHTMLText banner type
	BannerTypeXHTMLText BannerType = iota + 1
	// BannerTypeXHTML TypeXHTML banner type
	BannerTypeXHTML
	// BannerTypeJS TypeJS banner type
	BannerTypeJS
	// BannerTypeFrame TypeFrame banner type
	BannerTypeFrame
)

// Banner part of bid request
type Banner interface {
	// ID represent unique id for banner
	ID() string
	// Width is the width of the ad (max width bypassed cuz of being deprecated)
	Width() int
	// Height is the height of the ad (max height bypassed cuz of being deprecated)
	Height() int
	// BlockedTypes returns blocked creative types
	BlockedTypes() []BannerType
	// BlockedAttributes returns blocked creative attributes
	BlockedAttributes() []CreativeAttribute
	// Mimes returns the mime range of a banner
	Mimes() []string
	// Attributes returns ext and other useless stuff from request as a map
	Attributes() map[string]interface{}
}
