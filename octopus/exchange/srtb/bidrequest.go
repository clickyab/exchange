package srtb

import (
	"errors"
	"time"

	"encoding/json"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/srtb"
	"github.com/clickyab/services/random"
)

type bidRequest struct {
	inner *srtb.BidRequest
	imps  []exchange.Impression
	sup   exchange.Supplier
	time  time.Time
	cid   string
}

// NewBidRequest generate internal bid-request from simple rtb
func NewBidRequest(s exchange.Supplier, rq *srtb.BidRequest) exchange.BidRequest {
	return &bidRequest{sup: s, inner: rq}
}

func (b *bidRequest) CID() string {
	if b.cid == "" {
		b.cid = <-random.ID
	}
	return b.cid
}

func (b *bidRequest) UnmarshalJSON(a []byte) error {
	i := srtb.BidRequest{}
	err := json.Unmarshal(a, &i)
	if err != nil {
		return err
	}
	if i.Device == nil || i.Device.IP == "" {
		return errors.New("user ip (under device object) is required")
	}
	if len(i.Imp) == 0 {
		return errors.New("your bid request has no imp object")
	}
	for _, j := range i.Imp {
		if j.Banner == nil {
			return errors.New("imp object has no banner in it")
		}
	}
	if i.Site == nil && i.App == nil {
		return errors.New("there is no site or app object")
	}
	b.inner = &i
	return nil
}

func (b *bidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b *bidRequest) ID() string {
	return b.inner.ID
}

func (b *bidRequest) Imp() []exchange.Impression {
	if b.imps == nil {
		for _, m := range b.inner.Imp {
			b.imps = append(b.imps, &impression{inner: &m})
		}
	}
	return b.imps
}

func (b *bidRequest) Inventory() exchange.Inventory {
	if b.inner.Site != nil {
		return &site{inner: b.inner.Site}
	}
	if b.inner.App != nil {
		return &app{inner: b.inner.App}
	}
	panic("[BUG] not valid inventory")
}

func (b *bidRequest) Device() exchange.Device {
	return &device{inner: b.inner.Device}
}

func (b *bidRequest) User() exchange.User {
	return &user{inner: b.inner.User}
}

func (b *bidRequest) Test() bool {
	return b.inner.Test
}

func (b *bidRequest) AuctionType() exchange.AuctionType {
	return exchange.AuctionTypeSecondPrice
}

func (b *bidRequest) TMax() time.Duration {
	return time.Duration(b.inner.TMax) * time.Millisecond
}

func (b *bidRequest) WhiteList() []string {
	return []string{}
}

func (b *bidRequest) BlackList() []string {
	return []string{}
}

func (b *bidRequest) AllowedLanguage() []string {
	return []string{}
}

func (b *bidRequest) BlockedCategories() []string {
	return b.inner.BCat
}

func (b *bidRequest) BlockedAdvertiserDomain() []string {
	return []string{}
}

func (b *bidRequest) Time() time.Time {
	if b.time.IsZero() {
		panic("[BUG] time is not set")
	}
	return b.time
}

func (b *bidRequest) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}

func (b *bidRequest) LayerType() string {
	return exchange.SupplierSRTB
}
