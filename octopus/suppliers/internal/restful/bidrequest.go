package restful

import (
	"net/http"

	"time"

	"clickyab.com/exchange/octopus/exchange"
)

type bidRequest struct {
	IID     string                 `json:"id"`
	IImp    []imp                  `json:"imp"`
	ISite   *site                  `json:"site,omitempty"`
	IApp    *app                   `json:"app,omitempty"`
	IDevice device                 `json:"device"`
	IUser   user                   `json:"user"`
	ITest   bool                   `json:"test"`
	ITMax   time.Duration          `json:"-"`
	ITime   time.Time              `json:"time"`
	IAttr   map[string]interface{} `json:"attr"`
}

func (b bidRequest) ID() string {
	return b.IID
}

func (b bidRequest) Imp() []exchange.Impression {
	var res = make([]exchange.Impression, 0)
	for i := range b.IImp {
		res = append(res, b.IImp[i])
	}
	return res
}

func (b bidRequest) Inventory() exchange.Inventory {
	if b.ISite != nil {
		return b.ISite
	}
	return b.IApp
}

func (b bidRequest) Device() exchange.Device {
	return b.IDevice
}

func (b bidRequest) User() exchange.User {
	return b.IUser
}

func (b bidRequest) Test() bool {
	return b.ITest
}

func (b bidRequest) AuctionType() exchange.AuctionType {
	return exchange.AuctionTypeSecondPrice
}

func (b bidRequest) TMax() time.Duration {
	return b.ITMax
}

func (b bidRequest) WhiteList() []string {
	return []string{}
}

func (b bidRequest) BlackList() []string {
	return []string{}
}

func (b bidRequest) AllowedLanguage() []string {
	return []string{}
}

func (b bidRequest) BlockedCategories() []string {
	return []string{}
}

func (b bidRequest) BlockedAdvertiserDomain() []string {
	return []string{}
}

func (b bidRequest) Time() time.Time {
	return b.ITime
}

func (b bidRequest) Attributes() map[string]interface{} {
	return b.IAttr
}

// NewBidRequest get new bid request for rest clients
func NewBidRequest(r *http.Request) exchange.BidRequest {
	panic("not implemented yet")
}
