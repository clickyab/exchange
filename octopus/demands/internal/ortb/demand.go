package ortb

import (
	"context"
	"net/http"

	"io"

	"encoding/json"

	"bytes"

	"clickyab.com/exchange/octopus/demands/internal/base"
	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/ortb"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/xlog"
)

type Demand struct {
	exchange.DemandBase
}

func (d *Demand) Provide(ctx context.Context, bq exchange.BidRequest, ch chan exchange.BidResponse) {
	base.Provide(ctx, d, bq, ch)
}

func (d *Demand) GetBidResponse(ctx context.Context, r *http.Response, s exchange.Supplier) exchange.BidResponse {
	l := &bytes.Buffer{}
	k := &bytes.Buffer{}

	t := io.MultiWriter(l, k)
	_, err := io.Copy(t, r.Body)
	assert.Nil(err)

	defer r.Body.Close()

	p, err := l.ReadByte()
	xlog.Get(ctx).WithField("key", d.Name()).WithField("result", string(p)).Debug("Call done")
	de := json.NewDecoder(k)

	res := ortb.NewBidResponse(d, s, &openrtb.BidResponse{})
	err = de.Decode(&res)
	assert.Nil(err)
	return res
}

func (d *Demand) RenderBidRequest(ctx context.Context, w io.Writer, bq exchange.BidRequest) http.Header {
	if bq.LayerType() == "ortb" {
		j := json.NewEncoder(w)
		err := j.Encode(bq)
		assert.Nil(err)
		return http.Header{}
	}

	o := openrtb.BidRequest{
		Imp:  impression(bq.Imp()),
		ID:   bq.ID(),
		BAdv: bq.BlockedAdvertiserDomain(),
	}

	if v, ok := bq.Inventory().(exchange.App); ok {
		o.App = app(v)
	} else if v, ok := bq.Inventory().(exchange.Site); ok {
		o.Site = site(v)
	} else {
		panic("[BUG] invalid inventory")
	}

	xx := json.NewEncoder(w)
	err := xx.Encode(o)
	assert.Nil(err)
	return http.Header{}

}
func site(s exchange.Site) *openrtb.Site {
	return &openrtb.Site{
		Inventory: inventory(s),
		Page:      s.Page(),
		Ref:       s.Ref(),
	}
}
func inventory(n exchange.Inventory) openrtb.Inventory {
	return openrtb.Inventory{
		ID: n.ID(),
		Cat: func() []string {
			s := make([]string, 0)
			for _, v := range n.Cat() {
				s = append(s, string(v))
			}
			return s
		}(),
		Name:      n.Name(),
		Domain:    n.Domain(),
		Publisher: publisher(n.Publisher()),
	}
}
func publisher(p exchange.Publisher) *openrtb.Publisher {
	return &openrtb.Publisher{
		Domain: p.Domain(),
		Name:   p.Name(),
		Cat:    p.Cat(),
		ID:     p.ID(),
	}
}
func app(a exchange.App) *openrtb.App {
	return &openrtb.App{
		Bundle:    a.Bundle(),
		Inventory: inventory(a),
	}
}

func impression(m []exchange.Impression) []openrtb.Impression {
	ms := make([]openrtb.Impression, 0)
	for _, v := range m {
		t := openrtb.Impression{
			ID:     v.ID(),
			Banner: banner(v.Banner()),
			Secure: func() openrtb.NumberOrString {
				if v.Secure() {
					return openrtb.NumberOrString(1)
				}
				return openrtb.NumberOrString(0)
			}(),
			BidFloor:         v.BidFloor(),
			BidFloorCurrency: "IRR",
			Video:            video(v.Video()),
			Native:           native(v.Native()),
		}
		switch v.Type() {
		case exchange.AdTypeBanner:
			t.Banner = banner(v.Banner())
		case exchange.AdTypeVideo:
			t.Video = video(v.Video())
		case exchange.AdTypeNative:
			t.Native = native(v.Native())

		}

		ms = append(ms, t)
	}
	return ms

}
func native(n exchange.Native) *openrtb.Native {
	panic("implement me")
}
func video(b exchange.Video) *openrtb.Video {
	return &openrtb.Video{
		Mimes: b.Mimes(),
		W:     b.Width(),
		H:     b.Height(),
		Linearity: func() int {
			if b.Linearity() {
				return 1
			}
			return 0
		}(),
		BAttr: func() []int {
			res := make([]int, 0)
			for _, v := range b.BlockedAttributes() {
				res = append(res, int(v))
			}
			return res
		}(),
	}
}
func banner(b exchange.Banner) *openrtb.Banner {
	return &openrtb.Banner{
		ID: b.ID(),
		BType: func() []int {
			res := make([]int, 0)
			for _, v := range b.BlockedTypes() {
				res = append(res, int(v))
			}
			return res
		}(),
		H:     b.Height(),
		W:     b.Width(),
		Mimes: b.Mimes(),
		BAttr: func() []int {
			res := make([]int, 0)
			for _, v := range b.BlockedAttributes() {
				res = append(res, int(v))
			}
			return res
		}(),
	}

}
