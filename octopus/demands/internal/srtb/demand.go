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
	simple "clickyab.com/exchange/octopus/srtb"
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
func (d *Demand) GetBidResponse(ctx context.Context, resp *http.Response, sup exchange.Supplier) (exchange.BidResponse, error) {
	t := &bytes.Buffer{}
	_, err := io.Copy(t, resp.Body)
	assert.Nil(err)

	defer resp.Body.Close()

	p := t.Bytes()
	xlog.Get(ctx).WithField("key", d.Name()).WithField("result", string(p)).Debug("Call done")

	res := &simple.BidResponse{}
	err = json.Unmarshal(p, res)
	if err != nil {
		return nil, err
	}
	return srtb.NewBidResponse(d, sup, res), nil
}

// RenderBidRequest cast bid request to ortb
func (d *Demand) RenderBidRequest(ctx context.Context, w io.Writer, bq exchange.BidRequest) http.Header {
	j := json.NewEncoder(w)
	// render in rtb style
	err := j.Encode(bidRequestToSRTB(ctx, bq))
	assert.Nil(err)
	// TODO : simple rtb headers
	return http.Header{}
}

// bidRequestToSRTB change bid-request to srtb
// TODO : Split it to multiple simpler function
func bidRequestToSRTB(ctx context.Context, bq exchange.BidRequest) *simple.BidRequest {
	imps := []simple.Impression{}
	for i := range bq.Imp() {
		imps = append(imps, simple.Impression{
			ID: bq.Imp()[i].ID(),
			Banner: &simple.Banner{
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
