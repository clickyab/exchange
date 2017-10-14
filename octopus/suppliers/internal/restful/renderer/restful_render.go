package renderer

import (
	"encoding/json"

	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type dumbAd struct {
	TrackID   string `json:"track_id" structs:"track_id"`
	AdTrackID string `json:"ad_track_id" structs:"ad_track_id"`
	Winner    int64  `json:"winner" structs:"winner"`
	Width     int    `json:"width" structs:"width"`
	Height    int    `json:"height" structs:"height"`
	Code      string `json:"code" structs:"code"`
	IsFilled  bool   `json:"is_filled" structs:"is_filled"`
	Landing   string `json:"landing" structs:"landing"`
}

type restful struct {
	pixelPattern string
	sup          exchange.Supplier
}

func (rf restful) Render(resp exchange.BidResponse, w http.ResponseWriter) error {
	response := openrtb.BidResponse{}

	for i := range resp.Bids() {
		bid := resp.Bids()[i]

		response.SeatBid = append(response.SeatBid, openrtb.SeatBid{
			Bid: []openrtb.Bid{{
				ID:        bid.ID(),
				ImpID:     bid.ImpID(),
				Price:     float64(bid.Price()),
				AdID:      bid.AdID(),
				NURL:      bid.WinURL(),
				AdMarkup:  bid.AdMarkup(),
				AdvDomain: bid.AdDomains(),
				Cat:       bid.Categories(),
				W:         bid.AdWidth(),
				H:         bid.AdHeight(),
			}},
		})
	}

	enc := json.NewEncoder(w)
	w.WriteHeader(http.StatusOK)
	return enc.Encode(response)
}

// NewRestfulRenderer return a restful renderer
func NewRestfulRenderer(sup exchange.Supplier, pixel string) exchange.Renderer {
	return &restful{
		pixelPattern: pixel,
		sup:          sup,
	}
}

// TODO just for lint
var _ = dumbAd{}
