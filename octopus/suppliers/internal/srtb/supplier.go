package srtb

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/srtb"
	srtb2 "clickyab.com/exchange/octopus/srtb"
	"github.com/clickyab/services/assert"
)

// Supplier is srtb version of exchange-supplier
type Supplier struct {
	exchange.SupplierBase
}

// RenderBidResponse for rendering simple rtb
func (s *Supplier) RenderBidResponse(ctx context.Context, w io.Writer, b exchange.BidResponse) http.Header {
	if b.LayerType() == exchange.SupplierSRTB {
		r, err := json.Marshal(b)
		assert.Nil(err)
		w.Write(r)
		return http.Header{}
	}
	bids := make([]srtb2.Bid, 0)
	for _, i := range b.Bids() {
		bids = append(bids, srtb2.Bid{
			ID:       i.ID(),
			Width:    i.AdWidth(),
			Height:   i.AdHeight(),
			Price:    i.Price(),
			ImpID:    i.ImpID(),
			AdMarkup: i.AdMarkup(),
		})
	}
	r, err := json.Marshal(srtb2.BidResponse{
		ID:   b.ID(),
		Bids: bids,
	})
	assert.Nil(err)
	w.Write(r)
	return http.Header{}
}

// GetBidRequest transform request object to internal model
func (s *Supplier) GetBidRequest(ctx context.Context, r *http.Request) (exchange.BidRequest, error) {
	var r1 = srtb2.BidRequest{}
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := d.Decode(&r1); err != nil {
		return nil, err
	}
	res := srtb.NewBidRequest(s, &r1)
	return res, nil
}
