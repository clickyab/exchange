package srtb

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"

	"time"

	"clickyab.com/exchange/octopus/demands/internal/base"
	"clickyab.com/exchange/octopus/exchange/srtb"
	s "clickyab.com/exchange/octopus/srtb"
	"github.com/clickyab/services/xlog"
)

// Demand srtb demand structure
type Demand struct {
	exchange.DemandBase
}

// Provide method for demand
func (d *Demand) Provide(ctx context.Context, bq exchange.BidRequest, ch chan exchange.BidResponse) {
	base.Provide(ctx, d, bq, ch)
}

// GetBidResponse try to get bidresponse from response
func (d *Demand) GetBidResponse(ctx context.Context, resp *http.Response, sup exchange.Supplier) exchange.BidResponse {
	l := &bytes.Buffer{}
	k := &bytes.Buffer{}

	t := io.MultiWriter(l, k)
	_, err := io.Copy(t, resp.Body)
	assert.Nil(err)

	defer resp.Body.Close()

	p, err := l.ReadByte()
	assert.Nil(err)
	xlog.Get(ctx).WithField("key", d.Name()).WithField("result", string(p)).Debug("Call done")
	de := json.NewDecoder(k)

	res := &s.BidResponse{}
	err = de.Decode(res)
	assert.Nil(err)
	return srtb.NewBidResponse(d, sup, res)
}

// RenderBidRequest cast bid request to ortb
func (d *Demand) RenderBidRequest(ctx context.Context, w io.Writer, bq exchange.BidRequest) http.Header {
	j := json.NewEncoder(w)
	if bq.LayerType() == exchange.SupplierSRTB {
		err := j.Encode(bq)
		assert.Nil(err)
		return http.Header{}
	}
	// render in rtb style
	j.Encode(srtb.NewBidRequest(bq.Inventory().Supplier(), bidRequestRtbToRest(bq)))
	return http.Header{}
}

// bidRequestRtbToRest change bid-request to srtb
func bidRequestRtbToRest(bq exchange.BidRequest) *s.BidRequest {
	imps := []s.Impression{}
	for i := range bq.Imp() {
		imps = append(imps, s.Impression{
			ID: bq.Imp()[i].ID(),
			Banner: &s.Banner{
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
	res := &s.BidRequest{
		Imp: imps,
		ID:  bq.ID(),
		Device: &s.Device{
			UA:       bq.Device().UserAgent(),
			IP:       bq.Device().IP(),
			ConnType: int(bq.Device().ConnType()),
			Carrier:  bq.Device().Carrier(),
			Lang:     bq.Device().Language(),
			CID:      bq.Device().CID(),
			LAC:      bq.Device().LAC(),
			MNC:      bq.Device().MNC(),
			MCC:      bq.Device().MCC(),
			Geo: s.Geo{
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
		User: &s.User{
			ID: bq.User().ID(),
		},
		BCat: bq.BlockedCategories(),
		Test: bq.Test(),
		TMax: int(bq.TMax() / time.Millisecond),
	}

	if n, ok := bq.Inventory().(exchange.Site); ok {
		res.Site = &s.Site{
			Publisher: s.Publisher{
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
	} else if n, ok := bq.Inventory().(exchange.App); ok {
		res.App = &s.App{
			Publisher: s.Publisher{
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

	} else {
		panic("[BUG] not a valid inventory")
	}

	return res

}
