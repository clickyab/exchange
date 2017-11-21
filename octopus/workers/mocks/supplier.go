package mocks

import (
	"context"
	"net/http"

	"io"

	"clickyab.com/exchange/octopus/exchange"
)

type Supplier struct {
	SName         string
	SFloorCPM     int64
	SSoftFloorCPM int64
	SShare        int
	SCurrency     string
}

func (s Supplier) Currency() string {
	return s.SCurrency
}

func (s Supplier) RenderBidResponse(context.Context, io.Writer, exchange.BidResponse) http.Header {
	return http.Header{}
}

func (s Supplier) GetBidRequest(context.Context, *http.Request) (exchange.BidRequest, error) {
	panic("implement me")
}

func (s Supplier) Name() string {
	return s.SName
}

func (s Supplier) FloorCPM() int64 {
	return s.SFloorCPM
}

func (s Supplier) SoftFloorCPM() int64 {
	return s.SSoftFloorCPM
}

func (s Supplier) ExcludedDemands() []string {
	return []string{}
}

func (s Supplier) Share() int {
	return s.SShare
}

func (s Supplier) TestMode() bool {
	return false
}

func (s Supplier) Type() string {
	panic("implement me")
}
