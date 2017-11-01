package ortb

import (
	"encoding/json"
	"time"

	"errors"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/ip2location"
	"github.com/clickyab/services/random"
)

// NewBidRequest generate internal bid-request from open-rtb
func NewBidRequest(s exchange.Supplier, rq *BidRequest) exchange.BidRequest {
	record := ip2location.GetAll(rq.inner.Device.IP)
	rq.inner.Device.Geo = &openrtb.Geo{
		Region:  record.Region,
		Lon:     float64(record.Longitude),
		Lat:     float64(record.Latitude),
		City:    record.City,
		Country: record.CountryLong,
	}
	return &BidRequest{sup: s, inner: rq.inner, time: time.Now()}
}

// BidRequest bid request structure
type BidRequest struct {
	inner  *openrtb.BidRequest
	imps   []exchange.Impression
	sup    exchange.Supplier
	device exchange.Device
	time   time.Time
	cid    string
}

// CID clickyab track id
func (b *BidRequest) CID() string {
	if b.cid == "" {
		b.cid = <-random.ID
	}
	return b.cid
}

// LayerType (srtb/ortb)
func (b *BidRequest) LayerType() string {
	return exchange.SupplierORTB
}

// UnmarshalJSON json Unmarshaller
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
		return errors.New("user ip (under Device object) is required")
	}

	b.inner = &i
	return nil
}

// MarshalJSON json Marshaller
func (b *BidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.inner)
}

// ID bidrequest id
func (b *BidRequest) ID() string {
	return b.inner.ID
}

// Imp bidrequest imps
func (b *BidRequest) Imp() []exchange.Impression {
	if b.imps == nil {
		for _, m := range b.inner.Imp {
			b.imps = append(b.imps, &Impression{inner: &m})
		}
	}
	return b.imps
}

// Inventory bidrequest inventory
func (b *BidRequest) Inventory() exchange.Inventory {
	if b.inner.Site != nil {
		return &Site{inner: b.inner.Site, sup: b.sup}
	}
	if b.inner.App != nil {
		return &App{inner: b.inner.App, sup: b.sup}
	}
	panic("[BUG] not valid inventory")
}

// Device device entity
func (b *BidRequest) Device() exchange.Device {
	return &Device{
		inner: b.inner.Device,
		geo: &Geo{
			inner: b.inner.Device.Geo,
			ip:    b.inner.Device.IP,
		},
	}
}

// User user entity
func (b *BidRequest) User() exchange.User {
	return &User{inner: b.inner.User}
}

// Test test mode
func (b *BidRequest) Test() bool {
	if b.inner.Test == 1 {
		return true
	}
	return false
}

// AuctionType (second/first pricing)
func (b *BidRequest) AuctionType() exchange.AuctionType {
	return exchange.AuctionType(b.inner.AuctionType)
}

// TMax max timeout
func (b *BidRequest) TMax() time.Duration {
	return time.Duration(b.inner.TMax) * time.Millisecond
}

// WhiteList bq whitelist
func (b *BidRequest) WhiteList() []string {
	return b.inner.WSeat
}

// BlackList bq blacklist
func (b *BidRequest) BlackList() []string {
	return b.inner.BSeat
}

// AllowedLanguage bq allowed language
func (b *BidRequest) AllowedLanguage() []string {
	return b.inner.WLang
}

// BlockedCategories bq BlockedCategories
func (b *BidRequest) BlockedCategories() []string {
	return b.inner.Bcat
}

// BlockedAdvertiserDomain bq BlockedAdvertiserDomain
func (b *BidRequest) BlockedAdvertiserDomain() []string {
	return b.inner.BAdv
}

// Time time for bidrequest
func (b *BidRequest) Time() time.Time {
	if b.time.IsZero() {
		panic("[BUG] time is not set")
	}
	return b.time
}

// Attributes return extra attributes
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
