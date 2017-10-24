package ortb

import (
	"encoding/json"
	"net/http"
	"time"

	"errors"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

// ORTB is type of this layer
const ORTB = "ortb"

// NewBidRequest generate internal bid-request from open-rtb
func NewBidRequest(s exchange.Supplier, rq *openrtb.BidRequest) exchange.BidRequest {
	return &bidRequest{sup: s, inner: rq}
}

type bidRequest struct {
	inner *openrtb.BidRequest
	imps  []exchange.Impression
	sup   exchange.Supplier
	time  time.Time
}

func (b *bidRequest) LayerType() string {
	return ORTB
}

func (b *bidRequest) UnmarshalJSON(d []byte) error {
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
		return errors.New("user ip (under device object) is required")
	}

	b.inner = &i
	return nil
}

func (b *bidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.inner)
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
		return &site{inner: b.inner.Site, sup: &b.sup}
	}
	if b.inner.App != nil {
		return &app{inner: b.inner.App, sup: &b.sup}
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
	if b.inner.Test == 1 {
		return true
	}
	return false
}

func (b *bidRequest) AuctionType() exchange.AuctionType {
	return exchange.AuctionType(b.inner.AuctionType)
}

func (b *bidRequest) TMax() time.Duration {
	return time.Duration(b.inner.TMax) * time.Millisecond
}

func (b *bidRequest) WhiteList() []string {
	return b.inner.WSeat
}

func (b *bidRequest) BlackList() []string {
	return b.inner.BSeat
}

func (b *bidRequest) AllowedLanguage() []string {
	return b.inner.WLang
}

func (b *bidRequest) BlockedCategories() []string {
	return b.inner.Bcat
}

func (b *bidRequest) BlockedAdvertiserDomain() []string {
	return b.BlockedAdvertiserDomain()
}

func (b *bidRequest) Time() time.Time {
	if b.time.IsZero() {
		panic("[BUG] time is not set")
	}
	return b.time
}

func (b *bidRequest) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Ext":     b.inner.Ext,
		"AllImps": b.inner.AllImps,
		"BAdv":    b.inner.BAdv,
		"BApp":    b.inner.BApp,
		"Bcat":    b.inner.Bcat,
		"BSeat":   b.inner.BSeat,
		"Cur":     b.inner.Cur,
		"Regs":    b.inner.Regs,
		"WLang":   b.inner.WLang,
		"WSeat":   b.inner.WSeat,
	}
}

// GetBidRequest will generate bid-request from http request
func GetBidRequest(supplier exchange.Supplier, q *http.Request) (exchange.BidRequest, error) {
	r := &bidRequest{}
	jd := json.NewDecoder(q.Body)
	defer q.Body.Close()
	err := jd.Decode(r)
	if err != nil {
		return nil, err
	}
	r.sup = supplier
	return r, nil

}
