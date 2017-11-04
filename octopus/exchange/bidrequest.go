package exchange

import (
	"net/url"
	"time"
)

// AuctionType whether is first price or second price
type AuctionType int

const (
	// AuctionTypeFirstPrice is when the price is the first price. not supported by us
	AuctionTypeFirstPrice AuctionType = iota + 1
	// AuctionTypeSecondPrice second biding pricing
	AuctionTypeSecondPrice
)

// BidRequest is the single impression object
// Design note:
// XXX: Currency is not supported due to changes in the fucking Rial. if anything other than rial is in the request,
// reject the request
// XXX: Source is not supported
// XXX: all imps is currently not supported so the open rtb layer must reject any request with allimps set to 1
// XXX: Regulations is not supported.
type BidRequest interface {
	// ID return the random id of this imp object
	ID() string
	// Imp is the slot for this request
	Imp() []Impression
	// Inventory return the current inventory, caller should cast it to other inventory types
	Inventory() Inventory
	// Device return the device object
	Device() Device
	// User return the current user
	User() User
	// Test indicate this request is a test request and must not count in the financial calculation
	Test() bool
	// AuctionType is the way the auction is done. *JUST* accept the AuctionTypeSecondPrice
	AuctionType() AuctionType
	// TMax return the time that may accept the bid
	TMax() time.Duration
	// WhiteList the accepted advertiser/agency, is the WSeat in open rtb
	WhiteList() []string
	// BlackList the black listed advertiser/agency, is the BSeat in open rtb
	BlackList() []string
	// AllowedLanguage return a list of all allowed language for this request. also empty means all
	AllowedLanguage() []string
	// BlockedCategories return the list of all categories that must not be in this request's response
	BlockedCategories() []string
	// BlockedAdvertiserDomain return the array of top level domain of advertisers to block from this request
	// this contain both web site and apps.
	BlockedAdvertiserDomain() []string

	// Time time of the impression (the input time)
	Time() time.Time

	// Attributes is all extra data from input request
	Attributes() map[string]interface{}
	// CID unique identifier for every request
	CID() string
	// URL return the current url (host)
	URL() *url.URL
}
