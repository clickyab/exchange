package exchange

// Native part of bid request
type Native interface {
	// Extension returns marshaled data from request
	Extension() []byte
	// IsExtValid tells if we can unmarshal bytes passed to us
	IsExtValid() bool
	// AdLength returns number of ads needed from request
	AdLength() int
	// Attributes returns ext and other useless stuff from request as a map
	Attributes() map[string]interface{}
}
