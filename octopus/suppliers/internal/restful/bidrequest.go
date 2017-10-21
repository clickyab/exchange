package restful

import (
	"net/http"

	"time"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/rest"
)

type bidRequest struct {
	FID     string                 `json:"id"`
	FImp    []imp                  `json:"imp"`
	FSite   inventory              `json:"inventory"`
	FType   rest.BidType           `json:"type"`
	FDevice device                 `json:"device"`
	FUser   user                   `json:"user"`
	FTest   bool                   `json:"test"`
	FTMax   time.Duration          `json:"-"`
	FTime   time.Time              `json:"time"`
	FAttr   map[string]interface{} `json:"attr"`
}

func (b bidRequest) ID() string {
	return b.FID
}

func (b bidRequest) Imp() []exchange.Impression {
	var res = make([]exchange.Impression, 0)
	for i := range b.FImp {
		res = append(res, b.FImp[i])
	}
	return res
}

func (b bidRequest) Inventory() exchange.Inventory {
	return b.FSite
}

func (b bidRequest) Device() exchange.Device {
	return b.FDevice
}

func (b bidRequest) User() exchange.User {
	return b.FUser
}

func (b bidRequest) Test() bool {
	return b.FTest
}

func (b bidRequest) AuctionType() exchange.AuctionType {
	return exchange.AuctionTypeSecondPrice
}

func (b bidRequest) TMax() time.Duration {
	return b.FTMax
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
	return b.FTime
}

func (b bidRequest) Attributes() map[string]interface{} {
	return b.FAttr
}

// NewBidRequest get new bid request for rest clients
func NewBidRequest(r *http.Request) exchange.BidRequest {
	panic("not implemented yet")
}
