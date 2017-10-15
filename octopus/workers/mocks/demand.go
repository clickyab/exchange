package mocks

import (
	"context"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
)

type Demands struct {
	IName     string
	ICallRate int
	IHandicap int64
}

func (d Demands) Name() string {
	return d.IName
}

func (Demands) Provide(context.Context, exchange.BidRequest, chan exchange.BidResponse) {
	panic("implement me")
}

func (Demands) Win(context.Context, string, int64) {
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

func (Demands) WhiteListCountries() []string {
	panic("implement me")
}

func (Demands) ExcludedSuppliers() []string {
	panic("implement me")
}

func (Demands) TestMode() bool {
	panic("implement me")
}
