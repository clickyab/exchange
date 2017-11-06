package main

import (
	"context"
	"encoding/json"
	"net/http"

	"math/rand"

	"fmt"

	"github.com/bsm/openrtb"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/random"
	"github.com/sirupsen/logrus"
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

	rj, err := json.Marshal(createOrtbResponse(o, r))
	assert.Nil(err)
	logrus.Warn(rj)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rj)

}

func createOrtbResponse(o *openrtb.BidRequest, r *http.Request) openrtb.BidResponse {

	seat := make([]openrtb.SeatBid, 0)
	for _, v := range o.Imp {
		if v.Banner != nil {
			seat = append(seat, createOrtbBannerBide(o, &v, r))
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

func createOrtbBannerBide(o *openrtb.BidRequest, m *openrtb.Impression, r *http.Request) openrtb.SeatBid {
	scheme := func() string {
		if m.Secure == 1 {
			return "https"
		}
		return "http"
	}()

	return openrtb.SeatBid{
		Bid: []openrtb.Bid{
			{
				AdvDomain:  stringSlicer(advertisers),
				W:          m.Banner.W,
				H:          m.Banner.H,
				AdID:       <-random.ID,
				Cat:        stringSlicer(cats),
				Price:      float64(rand.Int63n(250)) + m.BidFloor,
				ImpID:      m.ID,
				ID:         <-random.ID,
				Protocol:   0,
				BURL:       fmt.Sprintf("%s://%s/api/burl/%s", scheme, host, m.ID),
				NURL:       fmt.Sprintf("%s://%s/api/nurl/%s", scheme, host, m.ID),
				CampaignID: openrtb.StringOrNumber(<-random.ID),
				AdMarkup: fmt.Sprintf(`<iframe width="%d" height="%d" src="%s://%s/api/ad/0?ortb=1&aid=${AUCTION_ID}&imp=${AUCTION_IMP_ID}&prc=${AUCTION_PRICE}&cur=${AUCTION_CURRENCY}&crl=${CLICK_URL:B64}&sho=${PIXEL_URL_JS:B64}&wi=%d&he=%d" frameborder="0"></iframe>`,
					m.Banner.W, m.Banner.H, scheme, host, m.Banner.W, m.Banner.H),
				LURL: fmt.Sprintf("%s://%s/api/lurl/%s", scheme, host, m.ID),
			},
		},
	}
}

func stringSlicer(m []string) []string {
	s, e := slicer(len(m))
	return m[s:e]
}

func slicer(m int) (s, e int32) {
	var i int32
	for {

		i = rand.Int31n(int32(m))
		if i != 0 {
			break
		}
	}

	return rand.Int31n(i), i
}
