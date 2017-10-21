package openrtb

import (
	"time"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type bidRequest struct {
	id          string
	impressions []exchange.Impression
	inventory   exchange.Inventory
	device      exchange.Device
	user        exchange.User
	test        bool
	auctionType exchange.AuctionType
	tMax        time.Duration
	whiteList, blackList, allowedLanguage, blockedCategories,
	blockedAdvertiseDomain []string
	time       time.Time
	attributes map[string]interface{}
}

func (b bidRequest) ID() string {
	return b.id
}

func (b bidRequest) Imp() []exchange.Impression {
	return b.impressions
}

func (b bidRequest) Inventory() exchange.Inventory {
	return b.inventory
}

func (b bidRequest) Device() exchange.Device {
	return b.device
}

func (b bidRequest) User() exchange.User {
	return b.user
}

func (b bidRequest) Test() bool {
	return b.test
}

func (b bidRequest) AuctionType() exchange.AuctionType {
	return b.auctionType
}

func (b bidRequest) TMax() time.Duration {
	return b.tMax
}

func (b bidRequest) WhiteList() []string {
	return b.whiteList
}

func (b bidRequest) BlackList() []string {
	return b.blackList
}

func (b bidRequest) AllowedLanguage() []string {
	return b.allowedLanguage
}

func (b bidRequest) BlockedCategories() []string {
	return b.blockedAdvertiseDomain
}

func (b bidRequest) BlockedAdvertiserDomain() []string {
	return b.blockedAdvertiseDomain
}

func (b bidRequest) Time() time.Time {
	return b.time
}

func (b bidRequest) Attributes() map[string]interface{} {
	return b.attributes
}

func newBidRequest(r *openrtb.BidRequest, supplier exchange.Supplier) exchange.BidRequest {
	return bidRequest{
		impressions: newImpressions(r.Imp),
		attributes:  requestAttributes(r),
		device:      newDevice(r.Device),
		inventory:   newInventory(r, supplier),
		time:        time.Now(),
		blockedAdvertiseDomain: r.BAdv,
		allowedLanguage:        r.WLang,
		blackList:              r.BSeat,
		whiteList:              r.WSeat,
		auctionType:            exchange.AuctionType(r.AuctionType),
		user:                   newUser(r.User),
		id:                     r.ID,
		blockedCategories:      r.Bcat,
		test: func() bool {
			if r.Test == 0 {
				return false
			}
			return true
		}(),
		tMax: time.Millisecond * time.Duration(r.TMax),
	}

}

func requestAttributes(r *openrtb.BidRequest) map[string]interface{} {
	return map[string]interface{}{
		"Ext":     r.Ext,
		"AllImps": r.AllImps,
		"BAdv":    r.BAdv,
		"BApp":    r.BApp,
		"Bcat":    r.Bcat,
		"BSeat":   r.BSeat,
		"Cur":     r.Cur,
		"Regs":    r.Regs,
		"WLang":   r.WLang,
		"WSeat":   r.WSeat,
	}
}
