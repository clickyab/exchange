package ortb

import (
	"context"
	"io"
	"net/http"

	"encoding/json"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/assert"
)

// RenderBidResponse for rendering open-rtb
func RenderBidResponse(ctx context.Context, s exchange.Supplier, w io.Writer, b exchange.BidResponse) http.Header {
	bids := func() []openrtb.SeatBid {
		x := make([]openrtb.SeatBid, 0)
		for _, b := range b.Bids() {
			x = append(x, openrtb.SeatBid{
				Bid: []openrtb.Bid{
					{
						ID:        b.ID(),
						ImpID:     b.ImpID(),
						Price:     float64(b.Price()),
						Cat:       b.Categories(),
						NURL:      b.WinURL(),
						AdID:      b.AdID(),
						H:         b.AdHeight(),
						W:         b.AdWidth(),
						AdMarkup:  b.AdMarkup(),
						AdvDomain: b.AdDomains(),
					},
				},
			})
		}
		return x
	}()

	r, err := json.Marshal(openrtb.BidResponse{
		NBR:      b.Excuse(),
		ID:       b.ID(),
		Currency: b.Supplier().Currency(),
		SeatBid:  bids,
	})

	assert.Nil(err)
	w.Write(r)
	header := http.Header{}
	header.Set("content-type", "application/json")
	return header
}
