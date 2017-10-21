package renderer

import (
	"encoding/json"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type restful struct {
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

// NewrtbRenderer return a restful renderer
func NewrtbRenderer() exchange.Renderer {
	return &restful{}
}
