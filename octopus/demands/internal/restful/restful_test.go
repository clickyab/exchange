package restful

import (
	"context"
	"net/http"
	"testing"
	"time"

	"clickyab.com/exchange/octopus/exchange"
	mock_entity "clickyab.com/exchange/octopus/exchange/mock_exchange"
	"clickyab.com/exchange/services/random"
	"clickyab.com/exchange/services/statistic"
	"clickyab.com/exchange/services/statistic/mock"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/jarcoal/httpmock.v1"
)

func handler(count int) (func(*http.Request) (*http.Response, error), []string) {
	ids := make([]string, 0)
	resps := make([]restAd, 0)
	for i := 0; i < count; i++ {
		ID := <-random.ID

		ids = append(ids, ID)
		resps = append(resps, restAd{
			RSlotTrackID: ID,
		})
	}
	return func(req *http.Request) (*http.Response, error) {
		resp, err := httpmock.NewJsonResponse(200, resps)
		if err != nil {
			return httpmock.NewStringResponse(500, ""), nil
		}
		return resp, nil
	}, ids
}

func TestDemandProvide(t *testing.T) {
	Convey("restful demand", t, func() {
		Convey("provide", func() {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			statistic.Register(mock.NewMockStatistic)
			ads, ids := handler(1)
			httpmock.RegisterResponder("POST", "http://127.0.0.1:9898", ads)
			ctx, cl := context.WithTimeout(context.Background(), time.Second*2)
			defer cl()
			imp := mock_entity.NewMockImpression(ctrl)
			res := make(chan exchange.Advertise)
			d := demand{
				client:     &http.Client{},
				key:        "test demand key",
				dayLimit:   1000,
				monthLimit: 1000,
				hourLimit:  1000,
				weekLimit:  1000,
				endPoint:   "http://127.0.0.1:9898",
				encoder: func(imp exchange.Impression) interface{} {
					return imp
				},
			}
			go d.Provide(ctx, imp, res)
			re := <-res
			So(re.SlotTrackID(), ShouldEqual, ids[0])
		})
	})
}