package srtb

import (
	"encoding/json"
	"time"

	"net/http"

	"errors"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/random"
	simple "github.com/clickyab/simple-rtb"
)

// NewSimpleRTBFromBidRequest generate a simple rtb instance from bid-request
func NewSimpleRTBFromBidRequest(in exchange.BidRequest) (exchange.BidRequest, error) {
	z, err := bidRequestToSRTB(in)
	if err != nil {
		return nil, err
	}

	return &bidRequest{inner: z, time: time.Now(), sup: in.Inventory().Supplier(), cid: in.CID(), request: in.Request()}, nil
}

// NewSimpleRTBFromRequest return make a bid-request from http request
func NewSimpleRTBFromRequest(s exchange.Supplier, r *http.Request) (exchange.BidRequest, error) {
	var r1 = &bidRequest{sup: s, time: time.Now(), cid: <-random.ID, request: r}
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := d.Decode(r1); err != nil {
		return nil, err
	}

	if !validateSupplier(r1, s) {
		return nil, errors.New("invalid impression currency")
	}

	for i := range r1.inner.Imp {
		r1.inner.Imp[i].BidFloor = exchange.IncShare(r1.inner.Imp[i].BidFloor, s.Share())
	}
	return r1, nil
}

// bidRequestToSRTB change bid-request to srtb
// TODO : Split it to multiple simpler function
func bidRequestToSRTB(bq exchange.BidRequest) (*simple.BidRequest, error) {
	var imps []simple.Impression
	for i := range bq.Imp() {
		imps = append(imps, simple.Impression{
			Currency: bq.Imp()[i].Currency(),
			ID:       bq.Imp()[i].ID(),
			Banner: &simple.Banner{
				ID:     bq.Imp()[i].Banner().ID(),
				Height: bq.Imp()[i].Banner().Height(),
				Width:  bq.Imp()[i].Banner().Width(),
			},
			BidFloor: func() float64 {
				if bq.Imp()[i].BidFloor() != 0 {
					return exchange.IncShare(bq.Imp()[i].BidFloor(), bq.Inventory().Supplier().Share())
				}
				return exchange.IncShare(float64(bq.Inventory().Supplier().FloorCPM()), bq.Inventory().Supplier().Share())
			}(),
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
				Country: simple.Country{
					Name:  bq.Device().Geo().Country().Name,
					ISO:   bq.Device().Geo().Country().Name,
					Valid: bq.Device().Geo().Country().Valid,
				},
				Region: simple.Region{
					Name:  bq.Device().Geo().Region().Name,
					ISO:   bq.Device().Geo().Region().ISO,
					Valid: bq.Device().Geo().Region().Valid,
				},
				LatLon: simple.LatLon{
					Lat:   bq.Device().Geo().LatLon().Lat,
					Lon:   bq.Device().Geo().LatLon().Lon,
					Valid: bq.Device().Geo().LatLon().Valid,
				},
				ISP: simple.ISP{},
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
					var res []string
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
					var res []string
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
		return nil, errors.New("[BUG] not a valid inventory")
	}

	return res, nil
}

func validateSupplier(br exchange.BidRequest, sup exchange.Supplier) bool {
	for _, imp := range br.Imp() {
		if imp.Currency() != sup.Currency() {
			return false
		}
	}
	return true
}
