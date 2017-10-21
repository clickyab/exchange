package rest

type BidType string

const (
	WebBidType  BidType = "web"
	VastBidType BidType = "vast"
	AppBidType  BidType = "app"
)
