package biding

import (
	"strconv"
	"testing"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/random"

	x "github.com/smartystreets/assertions"

	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/broker/mock"

	"clickyab.com/exchange/octopus/exchange/mock_exchange"

	"context"

	"net/http"

	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/kv"
	mock2 "github.com/clickyab/services/kv/mock"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

type situation struct {
	Case         int
	SoftFloorCPM int64
	FloorCPM     float64
	//MarginPercent int
	//DSCPMFloor    int64
	//DCPMFloor     int64
	//UnderFloor    bool
	Demands      []dem
	WinnerDemand int64
	//Profit3rdAd   float64
	//ProfitSuplier float64
	//description string
	Expectation func(interface{}, ...interface{}) string
}

type dem struct {
	Price    int64
	HandyCap int64
}

func cases() []situation {
	return []situation{
		{1, 300, 200, []dem{{470, 100}, {440, 110}}, 440, x.ShouldEqual},
		{2, 300, 200, []dem{{470, 100}, {440, 100}}, 441, x.ShouldEqual},
		{3, 300, 200, []dem{{470, 120}, {230, 115}}, 301, x.ShouldEqual},
		{4, 300, 200, []dem{{340, 100}, {300, 125}}, 300, x.ShouldEqual},
		{5, 300, 200, []dem{{230, 100}, {250, 110}}, 231, x.ShouldEqual},
		{6, 300, 200, []dem{{200, 120}, {230, 90}}, 200, x.ShouldEqual},
		{7, 300, 200, []dem{{250, 100}, {250, 100}}, 250, x.ShouldEqual},
		{8, 300, 200, []dem{{350, 100}}, 301, x.ShouldEqual},
		{9, 300, 200, []dem{{310, 100}}, 301, x.ShouldEqual},
		{10, 300, 200, []dem{{240, 120}}, 201, x.ShouldEqual},
	}

}

type Advertise struct {
	cpm,
	win int64
	demand exchange.Demand
}

var mk = `<div>
    <a href="${CLICK_URL}aHR0cHM6Ly9zdGFja292ZXJmbG93LmNvbS9qb2JzLzE1MzE2Ny9jb250aW51b3VzLWRlbGl2ZXJ5LWVuZ2luZWVyLW5ldGVudD9zbz1pJnBnPTEmb2Zmc2V0PS0xJnE9Z29sYW5nJmw9c3dlZGVuJnU9S20mZD0yMA=="></a>
    <a href="${CLICK_URL:B64}"></a>
    <script href="${PIXEL_URL_JS}"></script>
    <img src="${PIXEL_URL_IMAGE}" alt="">
    <script>
        function check() {
            var AUCTION_ID = "${AUCTION_ID}";
            var AUCTION_BID_ID = "${AUCTION_BID_ID}";
            var AUCTION_IMP_ID = "${AUCTION_IMP_ID}";
            var AUCTION_SEAT_ID = "${AUCTION_SEAT_ID}";
            var AUCTION_AD_ID = "${AUCTION_AD_ID}";
            var AUCTION_PRICE = "${AUCTION_PRICE}";
            var AUCTION_CURRENCY = "${AUCTION_CURRENCY}";
            var AUCTION_MBR = "${AUCTION_MBR}";
            var AUCTION_LOSS = "${AUCTION_LOSS}";
// request to server


            var PIXEL_URL_JS_b64 = "${PIXEL_URL_JS:B64}";
            var PIXEL_URL_IMAGE_b64 = "${PIXEL_URL_IMAGE:B64}";
            var AUCTION_ID_b64 = "${AUCTION_ID:B64}";
            var AUCTION_BID_ID_b64 = "${AUCTION_BID_ID:B64}";
            var AUCTION_IMP_ID_b64 = "${AUCTION_IMP_ID:B64}";
            var AUCTION_SEAT_ID_b64 = "${AUCTION_SEAT_ID:B64}";
            var AUCTION_AD_ID_b64 = "${AUCTION_AD_ID:B64}";
            var AUCTION_PRICE_b64 = "${AUCTION_PRICE:B64}";
            var AUCTION_CURRENCY_b64 = "${AUCTION_CURRENCY:B64}";
            var AUCTION_MBR_b64 = "${AUCTION_MBR:B64}";
            var AUCTION_LOSS_b64 = "${AUCTION_LOSS:B64}"
        }
    </script>
</div>

`

func TestSelect(t *testing.T) {
	router.AddRoute("pixel", "/api/:id/:type")
	router.AddRoute("click", "/api/:id")

	kv.Register(nil, nil, mock2.NewMockDistributedLocker, mock2.NewMockDsetStore, nil, nil, nil)

	ctrl := gomock.NewController(t)

	b := mock.GetChannelBroker()
	broker.SetActiveBroker(b)
	for _, u := range cases() {

		Convey("SelectCPM function test with case number "+strconv.Itoa(u.Case), t, func() {
			hr := &http.Request{
				Host: "test",
			}
			rq := mock_exchange.NewMockBidRequest(ctrl)
			rq.EXPECT().Request().Return(hr).AnyTimes()
			rq.EXPECT().Attributes().Return(map[string]interface{}{}).AnyTimes()
			id := <-random.ID
			rq.EXPECT().ID().Return(id).AnyTimes()
			us := mock_exchange.NewMockUser(ctrl)
			us.EXPECT().ID().Return(<-random.ID).AnyTimes()
			rq.EXPECT().User().Return(us).AnyTimes()
			inv := mock_exchange.NewMockInventory(ctrl)
			sup := mock_exchange.NewMockSupplier(ctrl)

			sup.EXPECT().SoftFloorCPM().Return(u.SoftFloorCPM).AnyTimes()
			sup.EXPECT().FloorCPM().Return(int64(u.FloorCPM)).AnyTimes()
			sup.EXPECT().Name().Return("Hello").AnyTimes()
			sup.EXPECT().Share().Return(0).AnyTimes()

			inv.EXPECT().Supplier().Return(sup).AnyTimes()
			rq.EXPECT().Inventory().Return(inv).AnyTimes()

			imp := mock_exchange.NewMockImpression(ctrl)
			impID := <-random.ID
			imp.EXPECT().ID().Return(impID).AnyTimes()
			imp.EXPECT().Secure().Return(false).AnyTimes()

			imp.EXPECT().BidFloor().Return(u.FloorCPM).AnyTimes()
			rq.EXPECT().Imp().Return([]exchange.Impression{imp}).AnyTimes()

			s := mock_exchange.NewMockSupplier(ctrl)
			s.EXPECT().Name().Return("test").AnyTimes()

			bids := make([]exchange.BidResponse, 0)

			for _, a := range u.Demands {
				d := mock_exchange.NewMockDemand(ctrl)
				d.EXPECT().Handicap().Return(a.HandyCap).AnyTimes()
				d.EXPECT().Bill(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, bid exchange.Bid) {}).AnyTimes()
				bi := mock_exchange.NewMockBid(ctrl)
				bi.EXPECT().ImpID().Return(impID).AnyTimes()
				bi.EXPECT().Price().Return(a.Price).AnyTimes()
				bi.EXPECT().AdMarkup().Return(mk).AnyTimes()
				bi.EXPECT().ID().Return(<-random.ID).AnyTimes()
				bi.EXPECT().AdID().Return(<-random.ID).AnyTimes()
				bi.EXPECT().BillURL().Return("").AnyTimes()
				bi.EXPECT().WinURL().Return("").AnyTimes()

				br := mock_exchange.NewMockBidResponse(ctrl)
				br.EXPECT().Bids().Return([]exchange.Bid{bi}).AnyTimes()
				bi.EXPECT().Demand().Return(d).AnyTimes()

				bids = append(bids, br)
			}

			res := SelectCPM(context.Background(), rq, bids)
			So(res.Bids()[0].Price(), u.Expectation, u.WinnerDemand)
		})
	}

}
