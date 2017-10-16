package restful

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/clickyab/services/kv/mock"

	. "github.com/smartystreets/goconvey/convey"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/mock_exchange"
	"github.com/clickyab/services/kv"
	"github.com/golang/mock/gomock"
	"gopkg.in/jarcoal/httpmock.v1"
)

func TestDemandProvide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	kv.Register(nil, nil, nil, nil, mock.NewAtomicMockStore, nil, nil)

	Convey("restful demand", t, func() {
		Convey("provide should panic", func() {
			host := "http://127.0.0.1:9898"
			httpmock.RegisterResponder("POST", host, func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(200, bidResponse{
					FID: "a",
					FBids: []bid{
						bid{
							FID: "b",
						}, bid{
							FID: "c",
						},
					},
				})
			})
			ctx, cl := context.WithTimeout(context.Background(), time.Second*2)
			defer cl()
			bq := mock_exchange.NewMockBidRequest(ctrl)
			bq.EXPECT().ID().Return("HAHAHA").AnyTimes()
			res := make(chan exchange.BidResponse)
			d := demand{
				client:     &http.Client{},
				key:        "test demand key",
				dayLimit:   1000,
				monthLimit: 1000,
				hourLimit:  1000,
				weekLimit:  1000,
				endPoint:   host,
				encoder: func(imp exchange.BidRequest) interface{} {
					return imp
				},
			}
			go d.Provide(ctx, bq, res)
			f := <-res
			So(f.ID(), ShouldEqual, "a")
			So(len(f.Bids()), ShouldEqual, 2)

		})

	})
}
