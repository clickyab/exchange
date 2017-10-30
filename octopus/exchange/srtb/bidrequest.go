package srtb

import (
	"errors"
	"time"

	"encoding/json"

	"io"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/srtb"
	"github.com/clickyab/services/assert"
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
	return &bidRequest{sup: s, inner: rq, time: time.Now()}
}

// RenderBidRequestRtbToRest change bidrequest rtb to rest
func RenderBidRequestRtbToRest(w io.Writer, bq exchange.BidRequest) {
	imps := []srtb.Impression{}
	for i := range bq.Imp() {
		imps = append(imps, srtb.Impression{
			ID: bq.Imp()[i].ID(),
			Banner: &srtb.Banner{
				ID:     bq.Imp()[i].Banner().ID(),
				Height: bq.Imp()[i].Banner().Height(),
				Width:  bq.Imp()[i].Banner().Width(),
			},
			BidFloor: bq.Imp()[i].BidFloor(),
			Secure: func() int {
				if bq.Imp()[i].Secure() {
					return 1
				}
				return 0
			}(),
		})
	}
	res := bidRequest{
		inner: &srtb.BidRequest{
			Imp: imps,
			ID:  bq.ID(),
			Device: &srtb.Device{
				UA:       bq.Device().UserAgent(),
				IP:       bq.Device().IP(),
				ConnType: int(bq.Device().ConnType()),
				Carrier:  bq.Device().Carrier(),
				Lang:     bq.Device().Language(),
				CID:      bq.Device().CID(),
				LAC:      bq.Device().LAC(),
				MNC:      bq.Device().MNC(),
				MCC:      bq.Device().MCC(),
				Geo: srtb.Geo{
					Country: exchange.Country{
						Name:  bq.Device().Geo().Country().Name,
						ISO:   bq.Device().Geo().Country().Name,
						Valid: bq.Device().Geo().Country().Valid,
					},
					Region: exchange.Region{
						Name:  bq.Device().Geo().Region().Name,
						ISO:   bq.Device().Geo().Region().ISO,
						Valid: bq.Device().Geo().Region().Valid,
					},
					LatLon: exchange.LatLon{
						Lat:   bq.Device().Geo().LatLon().Lat,
						Lon:   bq.Device().Geo().LatLon().Lon,
						Valid: bq.Device().Geo().LatLon().Valid,
					},
					ISP: exchange.ISP{},
				},
			},
			User: &srtb.User{
				ID: bq.User().ID(),
			},
			BCat: bq.BlockedCategories(),
			Test: bq.Test(),
			TMax: int(bq.TMax() / time.Millisecond),
		},
	}

	s, ok := bq.Inventory().(exchange.Site)
	if ok {
		res.inner.Site = &srtb.Site{
			Publisher: srtb.Publisher{
				ID:     s.ID(),
				Domain: s.Domain(),
				Cat: func() []string {
					res := []string{}
					for i := range s.Cat() {
						res = append(res, string(s.Cat()[i]))
					}
					return res
				}(),
				Name: s.Name(),
			},
			Ref:  s.Ref(),
			Page: s.Page(),
		}
	} else if s, ok := bq.Inventory().(exchange.App); ok {
		res.inner.App = &srtb.App{
			Publisher: srtb.Publisher{
				ID:     s.ID(),
				Domain: s.Domain(),
				Cat: func() []string {
					res := []string{}
					for i := range s.Cat() {
						res = append(res, string(s.Cat()[i]))
					}
					return res
				}(),
				Name: s.Name(),
			},
			Bundle: s.Bundle(),
		}

	} else {
		panic("[BUG]")
	}

	enc := json.NewEncoder(w)
	err := enc.Encode(res)
	assert.Nil(err)
}

// CID return srtb CID
func (b *bidRequest) CID() string {
	if b.cid == "" {
		b.cid = <-random.ID
	}
	return b.cid
}

// UnmarshalJSON return srtb UnmarshalJSON
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

// MarshalJSON return srtb MarshalJSON
func (b *bidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

// ID return srtb ID
func (b *bidRequest) ID() string {
	return b.inner.ID
}

// Imp return srtb Imp
func (b *bidRequest) Imp() []exchange.Impression {
	if b.imps == nil {
		for _, m := range b.inner.Imp {
			b.imps = append(b.imps, &impression{inner: &m, banner: &banner{inner: m.Banner}})
		}
	}
	return b.imps
}

// Inventory return srtb Inventory
func (b *bidRequest) Inventory() exchange.Inventory {
	if b.inner.Site != nil {
		return &site{inner: b.inner.Site, sup: b.sup}
	}
	if b.inner.App != nil {
		return &app{inner: b.inner.App, sup: b.sup}
	}
	panic("[BUG] not valid inventory")
}

// Device return srtb Device
func (b *bidRequest) Device() exchange.Device {
	return &device{inner: b.inner.Device}
}

// User return srtb User
func (b *bidRequest) User() exchange.User {
	return &user{inner: b.inner.User}
}

// Test return srtb Test
func (b *bidRequest) Test() bool {
	return b.inner.Test
}

// AuctionType return srtb AuctionType
func (b *bidRequest) AuctionType() exchange.AuctionType {
	return exchange.AuctionTypeSecondPrice
}

// TMax return srtb TMax
func (b *bidRequest) TMax() time.Duration {
	return time.Duration(b.inner.TMax) * time.Millisecond
}

// WhiteList return srtb WhiteList
func (b *bidRequest) WhiteList() []string {
	return []string{}
}

// BlackList return srtb BlackList
func (b *bidRequest) BlackList() []string {
	return []string{}
}

// AllowedLanguage return srtb AllowedLanguage
func (b *bidRequest) AllowedLanguage() []string {
	return []string{}
}

// BlockedCategories return srtb BlockedCategories
func (b *bidRequest) BlockedCategories() []string {
	return b.inner.BCat
}

// BlockedAdvertiserDomain return srtb BlockedAdvertiserDomain
func (b *bidRequest) BlockedAdvertiserDomain() []string {
	return []string{}
}

// Time return srtb Time
func (b *bidRequest) Time() time.Time {
	if b.time.IsZero() {
		panic("[BUG] time is not set")
	}
	return b.time
}

// Attributes return srtb Attributes
func (b *bidRequest) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}

// LayerType return bidrequest layer (srtb/ortb)
func (b *bidRequest) LayerType() string {
	return exchange.SupplierSRTB
}
