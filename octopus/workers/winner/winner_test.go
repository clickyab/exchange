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
		ctx, cl := context.WithCancel(base)
		defer cl()
		dem := consumer{ctx: ctx}
		delivery := dem.Consume()
		data := winToDelivery()
		select {
		case delivery <- data:
			So(true, ShouldBeTrue)
		case <-time.After(time.Second):
			So(true, ShouldBeFalse)
		}
		var x datamodels.TableModel
		select {
		case x = <-a.c:
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

		So(x.AdOutCount, ShouldEqual, 1)
	})
}
