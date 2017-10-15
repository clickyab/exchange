package renderer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/mock_exchange"
	"github.com/bsm/openrtb"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSupplier(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	Convey("render test", t, func() {

		bidResponse := mock_exchange.NewMockBidResponse(ctrl)

		var bids []exchange.Bid
		rands := []string{"123", "234", "345", "456"}
		for _, i := range rands {
			bid := mock_exchange.NewMockBid(ctrl)
			bid.EXPECT().ID().Return(i).AnyTimes()
			bid.EXPECT().ImpID().Return(i).AnyTimes()
			bid.EXPECT().Price().Return(int64(111)).AnyTimes()
			bid.EXPECT().AdID().Return(i).AnyTimes()
			bid.EXPECT().WinURL().Return("clickyab.com").AnyTimes()
			bid.EXPECT().AdMarkup().Return("<a href=\"http://adserver.com/click?adid=12345&tracker=clickurl.com\"><img src=\"http://image1.cdn.com/impid=102\"/></a>").AnyTimes()
			bid.EXPECT().AdDomains().Return([]string{"clickyab.com", "v.clickyab.com"}).AnyTimes()
			bid.EXPECT().Categories().Return([]string{"sport", "weather"}).AnyTimes()
			bid.EXPECT().AdWidth().Return(20).AnyTimes()
			bid.EXPECT().AdHeight().Return(200).AnyTimes()

			bids = append(bids, bid)
		}

		bidResponse.EXPECT().Bids().Return(bids).AnyTimes()
		writer := test{buff: &bytes.Buffer{}}
		err := NewRenderer().Render(bidResponse, http.ResponseWriter(writer))
		println(writer.buff.String())
		So(err, ShouldBeNil)

		var response openrtb.BidResponse
		err = json.Unmarshal(writer.buff.Bytes(), &response)
		So(err, ShouldBeNil)

		So(len(response.SeatBid), ShouldEqual, len(rands))

	})
}

type test struct {
	headers http.Header
	status  int
	buff    *bytes.Buffer
}

func (rw test) Header() http.Header {
	return rw.headers
}

func (rw test) Write(p []byte) (int, error) {
	return rw.buff.Write(p)
}

func (rw test) WriteHeader(i int) {
	rw.status = i
}
