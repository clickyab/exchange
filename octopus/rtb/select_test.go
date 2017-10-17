package rtb

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

func TestSelect(t *testing.T) {

	kv.Register(nil, nil, mock2.NewMockDistributedLocker, mock2.NewMockDsetStore, nil, nil, nil)

	ctrl := gomock.NewController(t)

	b := mock.GetChannelBroker()
	broker.SetActiveBroker(b)
	for _, u := range cases() {
		Convey("SelectCPM function test with case number "+strconv.Itoa(u.Case), t, func() {
			rq := mock_exchange.NewMockBidRequest(ctrl)
			rq.EXPECT().Attributes().Return(map[string]interface{}{}).AnyTimes()
			id := <-random.ID
			rq.EXPECT().ID().Return(id).AnyTimes()
			inv := mock_exchange.NewMockInventory(ctrl)
			sup := mock_exchange.NewMockSupplier(ctrl)

			sup.EXPECT().SoftFloorCPM().Return(u.SoftFloorCPM).AnyTimes()
			sup.EXPECT().FloorCPM().Return(int64(u.FloorCPM)).AnyTimes()
			sup.EXPECT().Name().Return("Hello").AnyTimes()

			inv.EXPECT().Supplier().Return(sup).AnyTimes()
			rq.EXPECT().Inventory().Return(inv).AnyTimes()

			imp := mock_exchange.NewMockImpression(ctrl)
			impID := <-random.ID
			imp.EXPECT().ID().Return(impID).AnyTimes()

			imp.EXPECT().BidFloor().Return(u.FloorCPM).AnyTimes()
			rq.EXPECT().Imp().Return([]exchange.Impression{imp})

			s := mock_exchange.NewMockSupplier(ctrl)
			s.EXPECT().Name().Return("test").AnyTimes()

			bids := make([]exchange.BidResponse, 0)

			for _, a := range u.Demands {
				d := mock_exchange.NewMockDemand(ctrl)
				d.EXPECT().Handicap().Return(a.HandyCap).AnyTimes()

				bi := mock_exchange.NewMockBid(ctrl)
				bi.EXPECT().ImpID().Return(impID).AnyTimes()
				bi.EXPECT().Price().Return(a.Price).AnyTimes()
				bi.EXPECT().AdID().Return(<-random.ID).AnyTimes()

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
