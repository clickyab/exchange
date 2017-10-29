package ortb

import (
	"encoding/json"
	"time"

	"errors"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/random"
)

// NewBidRequest generate internal bid-request from open-rtb
func NewBidRequest(s exchange.Supplier, rq *openrtb.BidRequest) exchange.BidRequest {
	return &BidRequest{sup: s, inner: rq}
}

type BidRequest struct {
	inner *openrtb.BidRequest
	imps  []exchange.Impression
	sup   exchange.Supplier
	time  time.Time
	cid   string
}

func (b *BidRequest) CID() string {
	if b.cid == "" {
		b.cid = <-random.ID
	}
	return b.cid
}

func (b *BidRequest) LayerType() string {
	return exchange.SupplierORTB
}

func (b *BidRequest) UnmarshalJSON(d []byte) error {
	i := openrtb.BidRequest{}
	err := json.Unmarshal(d, &i)
	if err != nil {
		return err
	}

	if err = i.Validate(); err != nil {
		return err
	}

	// TODO: extra validate
	if i.Device == nil || i.Device.IP == "" {
		return errors.New("User ip (under Device object) is required")
	}

	b.inner = &i
	return nil
}

func (b *BidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.inner)
}

func (b *BidRequest) ID() string {
	return b.inner.ID
}

func (b *BidRequest) Imp() []exchange.Impression {
	if b.imps == nil {
		for _, m := range b.inner.Imp {
			b.imps = append(b.imps, &Impression{inner: &m})
		}
	}
	return b.imps
}

func (b *BidRequest) Inventory() exchange.Inventory {
	if b.inner.Site != nil {
		return &Site{inner: b.inner.Site, sup: b.sup}
	}
	if b.inner.App != nil {
		return &App{inner: b.inner.App, sup: b.sup}
	}
	panic("[BUG] not valid inventory")
}

func (b *BidRequest) Device() exchange.Device {
	return &Device{inner: b.inner.Device}
}

func (b *BidRequest) User() exchange.User {
	return &User{inner: b.inner.User}
}

func (b *BidRequest) Test() bool {
	if b.inner.Test == 1 {
		return true
	}
	return false
}

func (b *BidRequest) AuctionType() exchange.AuctionType {
	return exchange.AuctionType(b.inner.AuctionType)
}

func (b *BidRequest) TMax() time.Duration {
	return time.Duration(b.inner.TMax) * time.Millisecond
}

func (b *BidRequest) WhiteList() []string {
	return b.inner.WSeat
}

func (b *BidRequest) BlackList() []string {
	return b.inner.BSeat
}

func (b *BidRequest) AllowedLanguage() []string {
	return b.inner.WLang
}

func (b *BidRequest) BlockedCategories() []string {
	return b.inner.Bcat
}

func (b *BidRequest) BlockedAdvertiserDomain() []string {
	return b.inner.BAdv
}

func (b *BidRequest) Time() time.Time {
	if b.time.IsZero() {
		panic("[BUG] time is not set")
	}
	return b.time
}

func (b *BidRequest) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Ext":     b.inner.Ext,
		"AllImps": b.inner.AllImps,
		"BApp":    b.inner.BApp,
		"Bcat":    b.inner.Bcat,
		"Cur":     b.inner.Cur,
		"Regs":    b.inner.Regs,
	}
}
