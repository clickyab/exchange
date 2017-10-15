package exchange

// ImpressionType which could contain (banner, video, native)
type ImpressionType int

const (
	// AdTypeBanner  is the type of ad
	AdTypeBanner ImpressionType = iota
	// AdTypeVideo is the type of ad
	AdTypeVideo
	// AdTypeNative is the type of ad
	AdTypeNative
)

// Impression is the imp part of bid request
type Impression interface {
	// ID represent unique id for impression
	ID() string
	// BidFloor returns the bid floor (and also a floor for biding 8-) )
	BidFloor() float64
	// Banner return banner of current imp
	Banner() Banner
	// Video return video of current imp
	Video() Video
	// Native return native of current imp
	Native() Native
	// Attributes returns ext and other useless stuff from request as a map
	Attributes() map[string]interface{}
	// Type returns the type of impression (e.g banner, video, natives)
	Type() ImpressionType
	// Secure tells if the ad required a https protocol
	Secure() bool
}
