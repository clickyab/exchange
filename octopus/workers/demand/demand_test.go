package demand

import (
	"testing"
	"time"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/materialize"
	"clickyab.com/exchange/octopus/workers/internal/datamodels"
	"clickyab.com/exchange/octopus/workers/mocks"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/broker"

	"context"

	. "github.com/smartystreets/goconvey/convey"

	"clickyab.com/exchange/octopus/exchange/mock_exchange"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/random"
	"github.com/golang/mock/gomock"
)

var (
	t1, _ = time.Parse("2006-01-02T15:04:05.000Z", "2017-03-21T00:01:00.000Z")
)

func newBidResponse(t time.Time, bidCount int, source, demand, sup string) exchange.BidResponse {
	a := make([]mocks.Bid, 0)
	for i := 1; i <= bidCount; i++ {
		a = append(a, mocks.Bid{
			IAdWidth:  300,
			IAdHeight: 250,
			IID:       <-random.ID,
			IDemand: mocks.Demands{
				IName: demand,
			},
		})
	}
	return mocks.BidResponse{
		ISupplier: mocks.Supplier{
			SName: sup,
		},
		IBids: a,
	}
}

func demToDelivery(rq exchange.BidRequest, resp exchange.BidResponse) broker.Delivery {
	job := materialize.DemandJob(rq, resp)
	d, err := job.Encode()
	assert.Nil(err)
	return mocks.JsonDelivery{Data: d}
}

type agg struct {
	c chan datamodels.TableModel
}

func (a *agg) Channel() chan<- datamodels.TableModel {
	return a.c
}

func TestDemand(t *testing.T) {
	ctrl := gomock.NewController(t)
	config.Initialize("test", "test", "test")
	a := &agg{c: make(chan datamodels.TableModel, 2)}
	datamodels.RegisterAggregator(a)
	base := context.Background()
	Convey("demand json job", t, func() {
		//d := newDemand("test_demand", 100, 50)
		resp := newBidResponse(t1, 2, "test_source", "test_demand", "test_supplier")

		//slots:=newSlots(2)
		ctx, cnl := context.WithCancel(base)
		defer cnl()
		dem := consumer{ctx: ctx}
		delivery := dem.Consume()
		rq := mock_exchange.NewMockBidRequest(ctrl)
		rq.EXPECT().ID().Return(<-random.ID).AnyTimes()
		rq.EXPECT().Time().Return(time.Now()).AnyTimes()
		imp := mock_exchange.NewMockImpression(ctrl)
		imp.EXPECT().ID().Return(<-random.ID).AnyTimes()
		imp.EXPECT().BidFloor().Return(250.).AnyTimes()
		imp.EXPECT().Secure().Return(false).AnyTimes()
		imp.EXPECT().Attributes().Return(map[string]interface{}{}).AnyTimes()
		imp.EXPECT().Type().Return(exchange.AdTypeBanner).AnyTimes()

		banner := mock_exchange.NewMockBanner(ctrl)
		banner.EXPECT().ID().Return(<-random.ID).AnyTimes()
		banner.EXPECT().Width().Return(300).AnyTimes()
		banner.EXPECT().Height().Return(250).AnyTimes()
		banner.EXPECT().Mimes().Return([]string{}).AnyTimes()
		banner.EXPECT().BlockedTypes().Return([]exchange.BannerType{}).AnyTimes()
		banner.EXPECT().BlockedAttributes().Return([]exchange.CreativeAttribute{}).AnyTimes()

		banner.EXPECT().Attributes().Return(map[string]interface{}{}).AnyTimes()
		imp.EXPECT().Banner().Return(banner).AnyTimes()
		imps := []exchange.Impression{
			imp,
		}
		rq.EXPECT().Imp().Return(imps).AnyTimes()
		inv := mock_exchange.NewMockInventory(ctrl)
		inv.EXPECT().Name().Return("inv").AnyTimes()
		sup := mock_exchange.NewMockSupplier(ctrl)
		sup.EXPECT().Name().Return("test_supplier").AnyTimes()
		sup.EXPECT().FloorCPM().Return(int64(100)).AnyTimes()
		sup.EXPECT().SoftFloorCPM().Return(int64(150)).AnyTimes()
		inv.EXPECT().SoftFloorCPM().Return(int64(150)).AnyTimes()
		inv.EXPECT().FloorCPM().Return(int64(150)).AnyTimes()
		inv.EXPECT().Attributes().Return(map[string]interface{}{}).AnyTimes()
		inv.EXPECT().Domain().Return("test_source").AnyTimes()
		sup.EXPECT().Share().Return(20).AnyTimes()
		sup.EXPECT().ExcludedDemands().Return([]string{}).AnyTimes()

		inv.EXPECT().Supplier().Return(sup).AnyTimes()
		rq.EXPECT().Inventory().Return(inv).AnyTimes()

		data := demToDelivery(rq, resp)
		// make sure this is not blocker, and if the test fails then may it hangs for ever
		select {
		case delivery <- data:
			So(true, ShouldBeTrue)
		case <-time.After(time.Second):
			So(true, ShouldBeFalse)
		}
		var t datamodels.TableModel
		select {
		case t = <-a.c:
			So(true, ShouldBeTrue)
		case <-time.After(time.Second):
			So(true, ShouldBeFalse)
		}
		select {
		case <-a.c:
			So(true, ShouldBeFalse)
		case <-time.After(time.Second):
			So(true, ShouldBeTrue)
		}
		So(t.Demand, ShouldEqual, "test_demand")
		So(t.RequestOutCount, ShouldEqual, 1)
		So(t.ImpressionInCount, ShouldEqual, 0)
		So(t.BidOutCount, ShouldEqual, 2)
		So(t.Source, ShouldEqual, "test_source")
		So(t.Supplier, ShouldEqual, "test_supplier")
	})

}
