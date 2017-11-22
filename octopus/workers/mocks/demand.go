package mocks

import (
	"context"
	"io"
	"net/http"

	"time"

	"clickyab.com/exchange/octopus/exchange"
)

type Demands struct {
	IName            string
	ICallRate        int
	IHandicap        int64
	IWhiteList       []string
	IExcludeSupplier []string
}

func (d Demands) Currencies() []string {
	panic("implement me")
}

func (d Demands) Win(context.Context, float64, string) {
	panic("implement me")
}

func (d Demands) Type() exchange.DemandType {
	panic("implement me")
}

func (d Demands) HasLimits() bool {
	panic("implement me")
}

func (d Demands) EndPoint() string {
	panic("implement me")
}

func (d Demands) GetTimeout() time.Duration {
	panic("implement me")
}

func (d Demands) GetBidResponse(context.Context, *http.Response, exchange.Supplier) (exchange.BidResponse, error) {
	panic("implement me")
}

func (d Demands) RenderBidRequest(context.Context, io.Writer, exchange.BidRequest) http.Header {
	panic("implement me")
}

func (d Demands) Name() string {
	return d.IName
}

func (Demands) Provide(context.Context, exchange.BidRequest, chan exchange.BidResponse) {
	panic("implement me")
}

func (Demands) Status(context.Context, http.ResponseWriter, *http.Request) {
	panic("implement me")
}

func (d Demands) Handicap() int64 {
	return d.IHandicap
}

func (d Demands) CallRate() int {
	return d.ICallRate
}

func (d Demands) WhiteListCountries() []string {

	return d.IWhiteList
}

func (d Demands) ExcludedSuppliers() []string {
	return d.IExcludeSupplier
}

func (Demands) TestMode() bool {
	panic("implement me")
}
func (d Demands) Bill(context.Context, float64, string) {

}
