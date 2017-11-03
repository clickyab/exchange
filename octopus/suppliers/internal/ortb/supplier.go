package ortb

import (
	"context"
	"io"
	"net/http"

	"encoding/json"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/ortb"
	"github.com/bsm/openrtb"
	"github.com/clickyab/services/assert"
)

// Supplier is ortb version of exchange-supplier
type Supplier struct {
	exchange.SupplierBase
}

// RenderBidResponse for rendering open-rtb
func (s *Supplier) RenderBidResponse(ctx context.Context, w io.Writer, b exchange.BidResponse) http.Header {
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
		Currency: "IRR",
		SeatBid:  bids,
	})

	assert.Nil(err)
	w.Write(r)
	return http.Header{}
}

// GetBidRequest transform request object to internal model
func (s *Supplier) GetBidRequest(ctx context.Context, r *http.Request) (exchange.BidRequest, error) {

	return ortb.NewOpenRTBFromRequest(s, r)
}
