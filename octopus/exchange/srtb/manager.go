package srtb

import (
	"context"
	"encoding/json"
	"time"

	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	simple "clickyab.com/exchange/octopus/srtb"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/xlog"
)

// NewSimpleRTBFromBidRequest generate a simple rtb instance from bid-request
func NewSimpleRTBFromBidRequest(ctx context.Context, in exchange.BidRequest) exchange.BidRequest {
	return &bidRequest{inner: bidRequestToSRTB(ctx, in), time: time.Now(), sup: in.Inventory().Supplier(), cid: in.CID()}
}

// NewSimpleRTBFromRequest return make a bid-request from http request
func NewSimpleRTBFromRequest(s exchange.Supplier, r *http.Request) (exchange.BidRequest, error) {
	var r1 = &bidRequest{sup: s, time: time.Now(), cid: <-random.ID}
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := d.Decode(r1); err != nil {
		return nil, err
	}
	return r1, nil
}

// bidRequestToSRTB change bid-request to srtb
// TODO : Split it to multiple simpler function
func bidRequestToSRTB(ctx context.Context, bq exchange.BidRequest) *simple.BidRequest {
	sh := float64(bq.Inventory().Supplier().Share() + 100)
	imps := []simple.Impression{}
	for i := range bq.Imp() {
		imps = append(imps, simple.Impression{
			ID: bq.Imp()[i].ID(),
			Banner: &simple.Banner{
				ID:     bq.Imp()[i].Banner().ID(),
				Height: bq.Imp()[i].Banner().Height(),
				Width:  bq.Imp()[i].Banner().Width(),
			},
			BidFloor: (sh + bq.Imp()[i].BidFloor()) / 100,
			Secure: func() int {
				if bq.Imp()[i].Secure() {
					return 1
				}
				return 0
			}(),
		})
	}
	res := &simple.BidRequest{
		Imp: imps,
		ID:  bq.ID(),
		Device: &simple.Device{
			UA:       bq.Device().UserAgent(),
			IP:       bq.Device().IP(),
			ConnType: int(bq.Device().ConnType()),
			Carrier:  bq.Device().Carrier(),
			Lang:     bq.Device().Language(),
			CID:      bq.Device().CID(),
			LAC:      bq.Device().LAC(),
			MNC:      bq.Device().MNC(),
			MCC:      bq.Device().MCC(),
			Geo: simple.Geo{
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
		User: &simple.User{
			ID: bq.User().ID(),
		},
		BCat: bq.BlockedCategories(),
		Test: func() int {
			if bq.Test() {
				return 1
			}
			return 0
		}(),
		TMax: int(bq.TMax() / time.Millisecond),
	}

	switch n := bq.Inventory().(type) {
	case exchange.Site:
		res.Site = &simple.Site{
			Publisher: simple.Publisher{
				ID:     n.ID(),
				Domain: n.Domain(),
				Cat: func() []string {
					res := []string{}
					for i := range n.Cat() {
						res = append(res, string(n.Cat()[i]))
					}
					return res
				}(),
				Name: n.Name(),
			},
			Ref:  n.Ref(),
			Page: n.Page(),
		}
	case exchange.App:
		res.App = &simple.App{
			Publisher: simple.Publisher{
				ID:     n.ID(),
				Domain: n.Domain(),
				Cat: func() []string {
					res := []string{}
					for i := range n.Cat() {
						res = append(res, string(n.Cat()[i]))
					}
					return res
				}(),
				Name: n.Name(),
			},
			Bundle: n.Bundle(),
		}

	default:
		xlog.Get(ctx).Panic("[BUG] not a valid inventory")
	}

	return res
}
