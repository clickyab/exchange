package srtb

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/srtb"
	"github.com/clickyab/services/assert"
	simple "github.com/clickyab/simple-rtb"
)

// Supplier is srtb version of exchange-supplier
type Supplier struct {
	exchange.SupplierBase
}

// RenderBidResponse for rendering simple rtb
func (s *Supplier) RenderBidResponse(ctx context.Context, w io.Writer, b exchange.BidResponse) http.Header {
	bids := make([]simple.Bid, 0)
	for _, i := range b.Bids() {
		bids = append(bids, simple.Bid{
			ID:       i.ID(),
			Width:    i.AdWidth(),
			Height:   i.AdHeight(),
			Price:    i.Price(),
			ImpID:    i.ImpID(),
			AdMarkup: i.AdMarkup(),
		})
	}
	r, err := json.Marshal(simple.BidResponse{
		ID:   b.ID(),
		Bids: bids,
	})
	assert.Nil(err)
	w.Write(r)
	return http.Header{}
}

// GetBidRequest transform request object to internal model
func (s *Supplier) GetBidRequest(ctx context.Context, r *http.Request) (exchange.BidRequest, error) {
	return srtb.NewSimpleRTBFromRequest(s, r)
}
