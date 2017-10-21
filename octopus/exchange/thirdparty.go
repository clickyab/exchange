package exchange

// ThirdParty is the third-party publisher, if any. if not available then the
// publisher is the inventory itself, normally another exchange system or
// ad network
type ThirdParty interface {
	// ID of publisher in our system
	ID() string
	// Name of publisher i.e clickyab
	Name() string
	// Cat is the category of publisher, empty means all
	Cat() []string
	// Domain of the publisher, optional
	Domain() string
	// Attributes return all unused fields of open rtb publisher
	Attributes() map[string]interface{}
}

// Publisher is the alias of third party.
type Publisher ThirdParty
