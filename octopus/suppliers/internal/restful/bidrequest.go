package restful

import (
	"time"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/random"
)

type BidRequest struct {
	IID     string                 `json:"id"`
	clid    string                 `json:"-"`
	IImp    []Imp                  `json:"Imp"`
	ISite   *Site                  `json:"Site,omitempty"`
	IApp    *App                   `json:"App,omitempty"`
	IDevice Device                 `json:"Device"`
	IUser   User                   `json:"User"`
	ITest   bool                   `json:"test"`
	ITMax   time.Duration          `json:"tmax,omitempty"`
	time    time.Time              `json:"-"`
	IWLang  []string               `json:"wlang,omitempty"`
	IBCat   []string               `json:"bcat,omitempty"`
	IBAdv   []string               `json:"badv,omitempty"`
	ILayer  string                 `json:"-"`
	IAttr   map[string]interface{} `json:"attr"`
}

func (b BidRequest) Layer() string {
	return b.ILayer
}

func (b BidRequest) ID() string {
	return b.IID
}

func (b BidRequest) Imp() []exchange.Impression {
	var res = make([]exchange.Impression, 0)
	for i := range b.IImp {
		res = append(res, b.IImp[i])
	}
	return res
}

func (b BidRequest) Inventory() exchange.Inventory {
	if b.ISite != nil {
		return b.ISite
	}
	return b.IApp
}

func (b BidRequest) Device() exchange.Device {
	return b.IDevice
}

func (b BidRequest) User() exchange.User {
	return b.IUser
}

func (b BidRequest) Test() bool {
	return b.ITest
}

func (b BidRequest) AuctionType() exchange.AuctionType {
	return exchange.AuctionTypeSecondPrice
}

func (b BidRequest) TMax() time.Duration {
	return b.ITMax
}

func (b BidRequest) WhiteList() []string {
	return []string{}
}

func (b BidRequest) BlackList() []string {
	return []string{}
}

func (b BidRequest) AllowedLanguage() []string {
	return b.IWLang
}

func (b BidRequest) BlockedCategories() []string {
	return b.IBCat
}

func (b BidRequest) BlockedAdvertiserDomain() []string {
	return b.IBAdv
}

func (b BidRequest) Time() time.Time {
	return b.time
}

func (b BidRequest) Attributes() map[string]interface{} {
	return b.IAttr
}

func (b BidRequest) CID() string {
	if b.clid == "" {
		return <-random.ID
	}
	return b.clid
}
