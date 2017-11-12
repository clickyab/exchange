package ortb

import (
	"encoding/json"
	"time"

	"errors"

	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/random"
)

// bidRequest bid request structure
type bidRequest struct {
	inner   *openrtb.BidRequest
	imps    []exchange.Impression
	sup     exchange.Supplier
	time    time.Time
	cid     string
	request *http.Request
}

func (b *bidRequest) Request() *http.Request {
	return b.request
}

// CID clickyab track id
func (b *bidRequest) CID() string {
	if b.cid == "" {
		b.cid = <-random.ID
	}
	return b.cid
}

// UnmarshalJSON json Unmarshaller
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

// MarshalJSON json Marshaller
func (b *bidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.inner)
}

// ID bidrequest id
func (b *bidRequest) ID() string {
	return b.inner.ID
}

// Imp bidrequest imps
func (b *bidRequest) Imp() []exchange.Impression {
	if b.imps == nil {
		for i := range b.inner.Imp {
			b.imps = append(b.imps, &impression{inner: &b.inner.Imp[i]})
		}
	}
	return b.imps
}

// Inventory bidrequest inventory
func (b *bidRequest) Inventory() exchange.Inventory {
	if b.inner.Site != nil {
		return &site{inner: b.inner.Site, sup: b.sup}
	}
	if b.inner.App != nil {
		return &app{inner: b.inner.App, sup: b.sup}
	}
	panic("[BUG] not valid inventory")
}

// device device entity
func (b *bidRequest) Device() exchange.Device {
	return &device{
		inner: b.inner.Device,
		geo: &geo{
			inner: b.inner.Device.Geo,
			ip:    b.inner.Device.IP,
		},
	}
}

// user user entity
func (b *bidRequest) User() exchange.User {
	return &user{inner: b.inner.User}
}

// Test test mode
func (b *bidRequest) Test() bool {
	if b.inner.Test == 1 {
		return true
	}
	return false
}

// AuctionType (second/first pricing)
func (b *bidRequest) AuctionType() exchange.AuctionType {
	return exchange.AuctionType(b.inner.AuctionType)
}

// TMax max timeout
func (b *bidRequest) TMax() time.Duration {
	return time.Duration(b.inner.TMax) * time.Millisecond
}

// WhiteList bq whitelist
func (b *bidRequest) WhiteList() []string {
	return b.inner.WSeat
}

// BlackList bq blacklist
func (b *bidRequest) BlackList() []string {
	return b.inner.BSeat
}

// AllowedLanguage bq allowed language
func (b *bidRequest) AllowedLanguage() []string {
	return b.inner.WLang
}

// BlockedCategories bq BlockedCategories
func (b *bidRequest) BlockedCategories() []string {
	return b.inner.Bcat
}

// BlockedAdvertiserDomain bq BlockedAdvertiserDomain
func (b *bidRequest) BlockedAdvertiserDomain() []string {
	return b.inner.BAdv
}

// Time time for bidrequest
func (b *bidRequest) Time() time.Time {
	if b.time.IsZero() {
		panic("[BUG] time is not set")
	}
	return b.time
}

// Attributes return extra attributes
func (b *bidRequest) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Ext":     b.inner.Ext,
		"AllImps": b.inner.AllImps,
		"BApp":    b.inner.BApp,
		"Bcat":    b.inner.Bcat,
		"Cur":     b.inner.Cur,
		"Regs":    b.inner.Regs,
	}
}
