package main

import (
	"context"
	"encoding/json"
	"net/http"

	"math/rand"

	"fmt"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/xlog"
	"github.com/clickyab/simple-rtb"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

// srtbHandler for handling exam (test) account
func srtbHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	o := &srtb.BidRequest{}
	j := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := j.Decode(o)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	br := srtb.BidResponse{ID: o.ID}
	br.Bids = createSrtbBid(ctx, o, r)
	rj, err := json.Marshal(br)
	assert.Nil(err)

	w.WriteHeader(http.StatusOK)
	w.Write(rj)

	xlog.GetWithFields(ctx, logrus.Fields{
		"demand_name": xmux.Param(ctx, "name"),
		"test_mode":   xmux.Param(ctx, "mode"),
	}).Debug()
}

func createSrtbBid(ctx context.Context, r *srtb.BidRequest, q *http.Request) []srtb.Bid {
	bs := make([]srtb.Bid, 0)
	data := map[string]interface{}{}
	for i := range r.Imp {
		if &r.Imp[i].Banner != nil {
			bid := createSrtbBannerBid(ctx, r, &r.Imp[i], q)
			bs = append(bs, bid)
			data[r.Imp[i].ID] = bid.Price
		}
	}
	xlog.SetField(ctx, "amount_bid", data)

	return bs
}

func createSrtbBannerBid(ctx context.Context, bq *srtb.BidRequest, m *srtb.Impression, r *http.Request) srtb.Bid {

	scheme := func() string {
		if m.Secure == 1 {
			return "https"
		}
		return "http"
	}()

	return srtb.Bid{
		AdMarkup: fmt.Sprintf(`<iframe width="%d" height="%d" src="%s://%s/api/ad/0?srtb=0&aid=${AUCTION_ID}&imp=${AUCTION_IMP_ID}&prc=${AUCTION_PRICE}&cur=${AUCTION_CURRENCY}&crl=${CLICK_URL:B64}&sho=${PIXEL_URL_JS:B64}&wi=%d&he=%d" frameborder="0"></iframe>`,
			m.Banner.Width, m.Banner.Height, scheme, host, m.Banner.Width, m.Banner.Height),
		ID:     fmt.Sprintf("%s-%s-%s", xmux.Param(ctx, "name"), xmux.Param(ctx, "mode"), <-random.ID),
		ImpID:  m.ID,
		Price:  int64(m.BidFloor) + rand.Int63n(250),
		Width:  m.Banner.Width,
		Height: m.Banner.Height,
	}
}
