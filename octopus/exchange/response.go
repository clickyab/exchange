package exchange

const (
	// ExcuseUnknown Unknown is iota
	ExcuseUnknown = iota
	// ExcuseTechnical is Technical
	ExcuseTechnical
	// ExcuseInvalid is Invalid
	ExcuseInvalid
	// ExcuseKnownWebSpider is KnownWebSpider
	ExcuseKnownWebSpider
	// ExcuseSuspectedNonHumanTraffic is SuspectedNonHumanTraffic
	ExcuseSuspectedNonHumanTraffic
	// ExcuseCloudDataCenterProxyIP is CloudDataCenterProxyIP
	ExcuseCloudDataCenterProxyIP
	// ExcuseUnsupportedDevice is UnsupportedDevice
	ExcuseUnsupportedDevice
	// ExcuseBlockedPublisherSite is BlockedPublisherSite
	ExcuseBlockedPublisherSite
	// ExcuseUnmatchedUser is UnmatchedUser
	ExcuseUnmatchedUser
	// ExcuseDailyReaderCapMet is DailyReaderCapMet
	ExcuseDailyReaderCapMet
	// ExcuseDailyDomainCapMet is DailyDomainCapMet
	ExcuseDailyDomainCapMet
)

// BidResponse is the response of bid request (demand to exchange, exchange to publisher)
type BidResponse interface {
	// ID returns the id of bid request made this response
	ID() string
	// Bids returns the first bid seats of response from demand
	Bids() []Bid
	// Excuse returns the reason we don't have an ad
	Excuse() int
	// Attribute data in ext and other spots
	Attributes() map[string]interface{}
	// Supplier return supplier
	Supplier() Supplier
}

// Bid is designed the way it returns the data about the first bid of the seat bid
type Bid interface {
	// ID returns the bid id
	ID() string
	// ImpID is the id of impression
	ImpID() string
	// Price if the bid
	Price() float64

	AdDetail
	// WinURL gives you the url to call if current bid won (NURL in openrtb)
	WinURL() string
	// BillURL gives you the url to call if current bid shown
	BillURL() string
	// Categories returns the category list of ad
	Categories() []string
	// Attributes returns the attribute about ad and bid
	Attributes() map[string]interface{}
	// Demand return demand
	Demand() Demand
}

// AdDetail is the ad detail that contains in each bid in response
type AdDetail interface {
	// AdID of the bid, needed if the bid won
	AdID() string
	// AdHeight returns the height of the ad
	AdHeight() int
	// AdHeight returns the width of the ad
	AdWidth() int
	// AdMarkup returns html markup for ad
	AdMarkup() string
	// AdDomains like asd.com/winthrough or asd.com (both are ok)
	AdDomains() []string
}
