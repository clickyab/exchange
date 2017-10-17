package winner

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"time"

	"context"

	"clickyab.com/exchange/octopus/workers/internal/datamodels"
	"clickyab.com/exchange/octopus/workers/mocks"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/config"
)

//func newImpression(t time.Time, slotCount int, source, sup string) exchange.BidRequest {
//	return mocks.BidRequest{
//		ITime: t,
//		IInventory: mocks.Publisher{
//			PName: source,
//			PSupplier: mocks.Supplier{
//				SName: sup,
//			},
//		},
//		IImps: make([]mocks.Imp, slotCount),
//	}
//}
//func newAdvertiser(cpm int64, dname string) exchange.Advertise {
//	return mocks.Advertiser{
//		MMaxCPM: cpm,
//		MDemand: mocks.Demands{
//			Mkey: dname,
//		},
//	}
//}

//func winnerToDelivery(imp exchange.BidRequest, ad exchange.Advertise, slot string) broker.Delivery {
//	job := materialize.WinnerJob(imp, ad, slot)
//	d, err := job.Encode()
//	assert.Nil(err)
//	return mocks.JsonDelivery{Data: d}
//}

var raw = `{"bid":{"height":1,"width":1,"ad_domains":["a.com"],"winner_cpm":12312312,"id":"id1","imp_id":"impID1","demand":{"name":"clickyab-demo.com"}},"request":{"id":"req_id_1","time":"2017-10-17T14:55:52.135644895+03:30","inventory":{"floor_cpm":111,"soft_floor_cpm":123,"name":"inventory1","supplier":{"floor_cpm":1111,"soft_floor_cpm":110,"name":"sup1","share":100},"domain":"inventory_domain"}}}`

type agg struct {
	c chan datamodels.TableModel
}

func (a *agg) Channel() chan<- datamodels.TableModel {
	return a.c
}

func winToDelivery() broker.Delivery {
	return mocks.JsonDelivery{Data: []byte(raw)}
}

func TestImpression(t *testing.T) {
	config.Initialize("test", "test", "test")
	a := &agg{c: make(chan datamodels.TableModel, 2)}
	datamodels.RegisterAggregator(a)
	base := context.Background()
	Convey("the demand test with the winner job", t, func() {
		//imp := newImpression(t1, 10, "test_winner", "test_demand")
		//adv := newAdvertiser(200, "adad")
		ctx, cl := context.WithCancel(base)
		defer cl()
		dem := consumer{ctx: ctx}
		delivery := dem.Consume()
		//data := winnerToDelivery(imp, adv, "aaa")
		data := winToDelivery()
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

		So(t.Supplier, ShouldEqual, "sup1")
		So(t.Source, ShouldEqual, "inventory_domain")
		So(t.Demand, ShouldEqual, "clickyab-demo.com")
		So(t.AdOutCount, ShouldEqual, 1)
		So(t.AdOutBid, ShouldEqual, 12312312)

		//So(t.Time, ShouldEqual, 1)
		//So(t.Request, ShouldEqual, 1)
		//So(t.Impression, ShouldEqual, 10)
		//
		//So(t.WinnerBid, ShouldBeZeroValue)
		//So(t.ShowBid, ShouldBeZeroValue)
		//So(t.Show, ShouldBeZeroValue)
		//So(t.Demands, ShouldBeZeroValue)
	})
}
