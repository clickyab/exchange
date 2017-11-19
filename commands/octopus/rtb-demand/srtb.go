package main

import (
	"context"
	"encoding/json"
	"net/http"

	"math/rand"

	"fmt"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/random"
	"github.com/clickyab/simple-rtb"
	"github.com/rs/xmux"
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

	rj, err := json.Marshal(createSrtbResponse(ctx, o, r))
	assert.Nil(err)

	w.WriteHeader(http.StatusOK)
	w.Write(rj)

}

func createSrtbResponse(ctx context.Context, o *srtb.BidRequest, r *http.Request) srtb.BidResponse {

	return srtb.BidResponse{
		ID:   <-random.ID,
		Bids: createSrtbBid(ctx, o, r),
	}

}
func createSrtbBid(ctx context.Context, r *srtb.BidRequest, q *http.Request) []srtb.Bid {
	bs := make([]srtb.Bid, 0)
	for i := range r.Imp {
		if &r.Imp[i].Banner != nil {
			bs = append(bs, createSrtbBannerBid(ctx, r, &r.Imp[i], q))
		}
	}

	return bs
}

func createSrtbBannerBid(ctx context.Context, bq *srtb.BidRequest, m *srtb.Impression, r *http.Request) srtb.Bid {

	scheme := func() string {
		if m.Secure == 1 {
			return "https"
		}
		return "http"
	}()
	adid := <-random.ID
	p := router.MustPath("rtb-demand-show", map[string]string{"id": adid})
	return srtb.Bid{
		AdMarkup: fmt.Sprintf(`<iframe width="%d" height="%d" src="%s://%s%s?srtb=0&aid=${AUCTION_ID}&imp=${AUCTION_IMP_ID}&prc=${AUCTION_PRICE}&cur=${AUCTION_CURRENCY}&crl=${CLICK_URL:B64}&sho=${PIXEL_URL_JS:B64}&wi=%d&he=%d" frameborder="0"></iframe>`,
			m.Banner.Width, m.Banner.Height, scheme, host, p, m.Banner.Width, m.Banner.Height),
		ID:     fmt.Sprintf("%s-%s-%s", xmux.Param(ctx, "name"), xmux.Param(ctx, "mode"), adid),
		ImpID:  m.ID,
		Price:  int64(m.BidFloor) + rand.Int63n(250),
		Width:  m.Banner.Width,
		Height: m.Banner.Height,
	}
}
