package exchange

type ImpressionType int

const (
	AdTypeBanner ImpressionType = iota
	AdTypeVideo
	AdTypeNative
)

type Impression interface {
	// ID represent unique id for impression
	ID() string
	// BidFloor returns the bid floor (and also a floor for biding 8-) )
	BidFloor() float64
	// Attributes returns ext and other useless stuff from request as a map
	Attributes() map[string]interface{}
	// Type returns the type of impression (e.g banner, video, natives)
	Type() ImpressionType
	// Secure tells if the ad required a https protocol
	Secure() bool
}
