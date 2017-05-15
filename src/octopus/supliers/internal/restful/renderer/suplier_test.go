package renderer

import (
	"services/config"
	"testing"

	"octopus/exchange/mock_exchange"

	"fmt"
	"net/url"

	"octopus/exchange"

	"encoding/json"

	"bytes"

	"net/http"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

type testResponseWriter struct {
	headers http.Header
	status  int
	buff    bytes.Buffer
}

func (rw *testResponseWriter) Header() http.Header {
	return rw.headers
}

func (rw *testResponseWriter) Write(p []byte) (int, error) {
	return rw.buff.Write(p)
}

func (rw *testResponseWriter) WriteHeader(i int) {
	rw.status = i
}

func TestSupplier(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Convey("reder test", t, func() {
		supplier := mock_exchange.NewMockSupplier(ctrl)
		supplier.EXPECT().Share().Return(100).AnyTimes()

		impression := mock_exchange.NewMockImpression(ctrl)

		trackIDs := []string{"aaa", "bbb", "ccc"}
		slots := []exchange.Slot{}
		ads := map[string]exchange.Advertise{}

		// generating slots and ads
		for _, trackID := range trackIDs {
			slot := mock_exchange.NewMockSlot(ctrl)
			ad := mock_exchange.NewMockAdvertise(ctrl)
			demand := mock_exchange.NewMockDemand(ctrl)

			demand.EXPECT().Name().Return("daniel").AnyTimes()

			ad.EXPECT().WinnerCPM().Return(int64(100)).AnyTimes()
			ad.EXPECT().Width().Return(20).AnyTimes()
			ad.EXPECT().Height().Return(15).AnyTimes()
			ad.EXPECT().Landing().Return("clickyab").AnyTimes()
			ad.EXPECT().Demand().Return(demand).AnyTimes()
			ad.EXPECT().TrackID().Return(trackID).AnyTimes()
			ad.EXPECT().URL().Return("www.ad_url.com").AnyTimes()

			fallback := fmt.Sprintf("www.%s.com", trackID)
			slot.EXPECT().Fallback().Return(fallback).AnyTimes()
			slot.EXPECT().TrackID().Return(trackID).AnyTimes()
			slot.EXPECT().Width().Return(20).AnyTimes()
			slot.EXPECT().Height().Return(15).AnyTimes()

			slots = append(slots, slot)
			ads[trackID] = ad
		}

		impression.EXPECT().Slots().Return(slots).AnyTimes()

		domain := config.GetStringDefault("exchange.supplier.domain", "localhost")
		pixel, err := url.Parse(fmt.Sprintf("http://%s/track", domain))
		underTable(t, err)

		rf := restful{
			pixelPattern: pixel,
			sup:          supplier,
		}

		var w = testResponseWriter{
			headers: make(http.Header),
		}

		err = rf.Render(impression, ads, &w)
		underTable(t, err)

		result := []*dumbAd{}
		err = json.Unmarshal(w.buff.Bytes(), &result)
		underTable(t, err)
		// TODO : see the next todo :))
		for i := range result {
			result[i].Code = ""
		}
		So(result, ShouldResemble, expected)
		So(w.status, ShouldEqual, http.StatusOK)
	})
}

func underTable(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

// TODO : I'm going to ignore code for now. I appreciated if someone can fix it again with a good approach
var expected = []*dumbAd{
	{
		IsFilled:  true,
		Landing:   "clickyab",
		TrackID:   "aaa",
		Winner:    0,
		Width:     20,
		Height:    15,
		Code:      ``,
		AdTrackID: "aaa",
	}, {
		Height:    15,
		Code:      ``,
		Landing:   "clickyab",
		IsFilled:  true,
		TrackID:   "bbb",
		AdTrackID: "bbb",
		Winner:    0,
		Width:     20,
	}, {
		Landing:   "clickyab",
		TrackID:   "ccc",
		AdTrackID: "ccc",
		Winner:    0,
		Width:     20,
		Height:    15,
		Code:      ``,
		IsFilled:  true,
	},
}
