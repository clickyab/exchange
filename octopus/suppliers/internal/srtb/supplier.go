package srtb

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"
	simple "github.com/clickyab/simple-rtb"
)

// RenderBidResponse for rendering simple rtb
func RenderBidResponse(ctx context.Context, s exchange.Supplier, w io.Writer, b exchange.BidResponse) http.Header {
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
	header := http.Header{}
	header.Set("content-type", "application/json")
	return header
}
