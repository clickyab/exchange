package mocks

import (
	"net"
	"time"

	"net/url"

	"clickyab.com/exchange/octopus/exchange"
)

type BidRequest struct {
	IID          string
	IIP          net.IP
	ISchema      string
	ITest        bool
	IAuctionType exchange.AuctionType
	IUserTrackID string
	IPageTrackID string
	IUserAgent   string
	IInventory   Inventory

	IAttributes map[string]interface{}
	IImps       []Imp
	ICategory   []exchange.Category
	IPlatform   exchange.DeviceType
	ITime       time.Time
	ITMax       time.Duration
	IDevice     Device
}

func (b *BidRequest) URL() *url.URL {
	panic("implement me")
}

func (b *BidRequest) CID() string {
	panic("implement me")
}

func (b *BidRequest) ID() string {
	return b.IID
}

func (b *BidRequest) Imp() []exchange.Impression {
	var res = make([]exchange.Impression, 0)
	for _, val := range b.IImps {
		res = append(res, val)
	}
	return res
}

func (b *BidRequest) Inventory() exchange.Inventory {
	return b.IInventory
}

func (b *BidRequest) Device() exchange.Device {
	return &b.IDevice
}

func (b *BidRequest) User() exchange.User {
	panic("implement me")
}

func (b *BidRequest) Test() bool {
	return b.ITest
}

func (b *BidRequest) AuctionType() exchange.AuctionType {
	return b.IAuctionType
}

func (b *BidRequest) TMax() time.Duration {
	return b.ITMax
}

func (b *BidRequest) WhiteList() []string {
	return []string{}
}

func (b *BidRequest) BlackList() []string {
	return []string{}
}

func (b *BidRequest) AllowedLanguage() []string {
	return []string{}
}

func (b *BidRequest) BlockedCategories() []string {
	return []string{}
}

func (b *BidRequest) BlockedAdvertiserDomain() []string {
	return []string{}
}

func (b *BidRequest) Time() time.Time {
	return b.ITime
}

func (b *BidRequest) Attributes() map[string]interface{} {
	return b.IAttributes
}
