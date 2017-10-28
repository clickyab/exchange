package static

import (
	"context"
	"encoding/json"
	"net/http"

	"math/rand"

	"clickyab.com/exchange/octopus/srtb"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/random"
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

	rj, err := json.Marshal(createSrtbResponse(o))
	assert.Nil(err)
	w.Write(rj)
	w.WriteHeader(http.StatusOK)

}

func createSrtbResponse(o *srtb.BidRequest) srtb.BidResponse {

	return srtb.BidResponse{
		ID:   <-random.ID,
		Bids: createSrtbBid(o),
	}

}
func createSrtbBid(r *srtb.BidRequest) []srtb.Bid {
	bs := make([]srtb.Bid, 0)
	for i := range r.Imp {
		if &r.Imp[i].Banner != nil {
			bs = append(bs, createSrtbBannerBid(r, &r.Imp[i]))
		}
	}
	return bs
}

func createSrtbBannerBid(request *srtb.BidRequest, m *srtb.Impression) srtb.Bid {
	return srtb.Bid{
		AdMarkup: `<div>TEST</div>`,
		ID:       <-random.ID,
		ImpID:    m.ID,
		Price:    int64(m.BidFloor) + rand.Int63n(250),
		Width:    m.Banner.Width,
		Height:   m.Banner.Height,
	}
}
