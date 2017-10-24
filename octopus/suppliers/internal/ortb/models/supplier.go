package models

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

// Supplier is ortb version of exchange supplier
type Supplier struct {
	name string
	floorCPM,
	softFloorCPM int64
	excludedDemands []string
	share           int
	test            bool
}

// Name of supplier
func (s *Supplier) Name() string {
	return s.name
}

// FloorCPM is default value for requests
func (s *Supplier) FloorCPM() int64 {
	return s.floorCPM
}

// SoftFloorCPM is default value for requests
func (s *Supplier) SoftFloorCPM() int64 {
	return s.softFloorCPM
}

// ExcludedDemands is the black list demands for this supplier
func (s *Supplier) ExcludedDemands() []string {
	return s.excludedDemands
}

// Share is profit margin of exchange
func (s *Supplier) Share() int {
	return s.share
}

// RenderBidResponse will write the response to the writer
func (s *Supplier) RenderBidResponse(ctx context.Context, w io.Writer, b exchange.BidResponse) http.Header {
	if b.LayerType() == ortb.ORTB {
		r, err := json.Marshal(b)
		assert.Nil(err)
		w.Write(r)
		return http.Header{}
	}

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

// TestMode will determine if the supplier is in test mode or not
func (s *Supplier) TestMode() bool {
	return s.test
}

// Type of supplier (ortb, srtb)
func (s *Supplier) Type() string {
	return ortb.ORTB
}

// GetBidRequest transform request object to internal model
func (s *Supplier) GetBidRequest(ctx context.Context, r *http.Request) exchange.BidRequest {
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	res := ortb.NewBidRequest(s, &openrtb.BidRequest{})
	assert.Nil(d.Decode(res))

	return res
}
