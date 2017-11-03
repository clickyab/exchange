package ortb

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/xlog"
)

// NewOpenRTBFromBidRequest generate a open rtb instance from bid-request
func NewOpenRTBFromBidRequest(ctx context.Context, in exchange.BidRequest) exchange.BidRequest {
	o := &openrtb.BidRequest{
		Imp:  newImpression(in.Imp(), in.Inventory().Supplier().Share()),
		ID:   in.ID(),
		BAdv: in.BlockedAdvertiserDomain(),
	}

	switch v := in.Inventory().(type) {
	case exchange.App:
		o.App = newApp(v)
	case exchange.Site:
		o.Site = newSite(v)
	default:
		xlog.Get(ctx).Panic("[BUG] invalid inventory")
	}

	return &bidRequest{sup: in.Inventory().Supplier(), time: time.Now(), cid: in.CID(), inner: o}
}

// NewOpenRTBFromRequest generate internal bid-request from open-rtb
func NewOpenRTBFromRequest(s exchange.Supplier, r *http.Request) (exchange.BidRequest, error) {
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	rq := &bidRequest{sup: s, time: time.Now()}
	if err := d.Decode(rq); err != nil {
		return nil, err
	}
	return rq, nil
}
func newSite(s exchange.Site) *openrtb.Site {
	return &openrtb.Site{
		Inventory: newInventory(s),
		Page:      s.Page(),
		Ref:       s.Ref(),
	}
}

func newInventory(n exchange.Inventory) openrtb.Inventory {
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
		Publisher: newPublisher(n.Publisher()),
	}
}

func newPublisher(p exchange.Publisher) *openrtb.Publisher {
	return &openrtb.Publisher{
		Domain: p.Domain(),
		Name:   p.Name(),
		Cat:    p.Cat(),
		ID:     p.ID(),
	}
}
func newApp(a exchange.App) *openrtb.App {
	return &openrtb.App{
		Bundle:    a.Bundle(),
		Inventory: newInventory(a),
	}
}

func newImpression(m []exchange.Impression, share int) []openrtb.Impression {
	sh := float64(share + 100)
	ms := make([]openrtb.Impression, 0)
	for _, v := range m {
		t := openrtb.Impression{
			ID: v.ID(),
			Secure: func() openrtb.NumberOrString {
				if v.Secure() {
					return openrtb.NumberOrString(1)
				}
				return openrtb.NumberOrString(0)
			}(),
			BidFloor:         (sh * v.BidFloor()) / 100,
			BidFloorCurrency: "IRR",
		}
		switch v.Type() {
		case exchange.AdTypeBanner:
			t.Banner = newBanner(v.Banner())
		case exchange.AdTypeVideo:
			t.Video = newVideo(v.Video())
		case exchange.AdTypeNative:
			t.Native = newNative(v.Native())

		}

		ms = append(ms, t)
	}
	return ms

}

func newNative(n exchange.Native) *openrtb.Native {
	panic("implement me")
}

func newVideo(b exchange.Video) *openrtb.Video {
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

func newBanner(b exchange.Banner) *openrtb.Banner {
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
