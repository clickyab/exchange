package main

import (
	"context"
	"encoding/json"
	"net/http"

	"math/rand"

	"github.com/bsm/openrtb"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/random"
)

// ortbHandler for handling exam (test) account
func ortbHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	o := &openrtb.BidRequest{}
	j := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := j.Decode(o)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}

	err = o.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	rj, err := json.Marshal(createOrtbResponse(o))
	assert.Nil(err)
	w.WriteHeader(http.StatusOK)
	w.Write(rj)

}

func createOrtbResponse(o *openrtb.BidRequest) openrtb.BidResponse {
	seat := make([]openrtb.SeatBid, 0)
	for _, v := range o.Imp {
		if v.Banner != nil {
			seat = append(seat, createOrtbBannerBide(o, &v))
		}
	}
	return openrtb.BidResponse{
		ID:       <-random.ID,
		Currency: "IRR",
		SeatBid:  seat,
	}

}

var advertisers = []string{
	"clickyab.com",
	"adro.com",
	"technica.com",
	"google.com",
	"facebook.il",
}
var cats = []string{
	"iab-sport",
	"iab-news",
	"iab-art",
}

func createOrtbBannerBide(o *openrtb.BidRequest, m *openrtb.Impression) openrtb.SeatBid {
	return openrtb.SeatBid{
		Bid: []openrtb.Bid{
			{
				AdvDomain:  stringSlicer(advertisers),
				W:          m.Banner.W,
				H:          m.Banner.H,
				AdID:       <-random.ID,
				Cat:        stringSlicer(cats),
				Price:      float64(rand.Int63n(500)) + m.BidFloor,
				ImpID:      m.ID,
				ID:         <-random.ID,
				Protocol:   0,
				BURL:       "http://changeme/",
				NURL:       "http://changeme/",
				CampaignID: openrtb.StringOrNumber(<-random.ID),
				AdMarkup:   `<div>TEST DEMAND</div>`,
				LURL:       "http://changeme/",
			},
		},
	}
}

func stringSlicer(m []string) []string {
	s, e := slicer(len(m))
	return m[s:e]
}

func slicer(m int) (s, e int32) {
	i := rand.Int31n(int32(m))
	return rand.Int31n(i), i
}
