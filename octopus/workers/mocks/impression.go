package mocks

import (
	"net"
	"time"

	"clickyab.com/exchange/octopus/exchange"
)

type BidRequest struct {
	ITrackID     string
	IIP          net.IP
	ISchema      string
	IUserTrackID string
	IPageTrackID string
	IUserAgent   string
	ISource      exchange.Inventory
	ILocation    Location
	IAttributes  map[string]interface{}
	IImps        []*Imp
	ICategory    []exchange.Category
	IPlatform    exchange.DeviceType
	IUnderFloor  bool
	ITime        time.Time
}

func (*BidRequest) ID() string {
	panic("implement me")
}

func (*BidRequest) Imp() []exchange.Impression {
	panic("implement me")
}

func (*BidRequest) Inventory() exchange.Inventory {
	panic("implement me")
}

func (*BidRequest) Device() exchange.Device {
	panic("implement me")
}

func (*BidRequest) User() exchange.User {
	panic("implement me")
}

func (*BidRequest) Test() bool {
	panic("implement me")
}

func (*BidRequest) AuctionType() exchange.AuctionType {
	panic("implement me")
}

func (*BidRequest) TMax() time.Duration {
	panic("implement me")
}

func (*BidRequest) WhiteList() []string {
	panic("implement me")
}

func (*BidRequest) BlackList() []string {
	panic("implement me")
}

func (*BidRequest) AllowedLanguage() []string {
	panic("implement me")
}

func (*BidRequest) BlockedCategories() []string {
	panic("implement me")
}

func (*BidRequest) BlockedAdvertiserDomain() []string {
	panic("implement me")
}

func (*BidRequest) Time() time.Time {
	panic("implement me")
}

func (*BidRequest) Attributes() map[string]interface{} {
	panic("implement me")
}
