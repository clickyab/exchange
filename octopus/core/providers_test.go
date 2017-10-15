package core

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"clickyab.com/exchange/octopus/exchange"
	mock_entity "clickyab.com/exchange/octopus/exchange/mock_exchange"
	"github.com/clickyab/services/random"

	"github.com/clickyab/services/config"

	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/broker/mock"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/fzerorubigd/onion.v2"
)

func newInventory(c *gomock.Controller) exchange.Inventory {
	s := mock_entity.NewMockSupplier(c)
	s.EXPECT().TestMode().Return(false).AnyTimes()
	s.EXPECT().ExcludedDemands().Return([]string{}).AnyTimes()
	tmp := mock_entity.NewMockInventory(c)
	tmp.EXPECT().Name().Return("publisher").AnyTimes()
	tmp.EXPECT().Supplier().Return(s).AnyTimes()
	tmp.EXPECT().FloorCPM().Return(int64(100)).AnyTimes()
	return tmp
}
func newBidRequest(c *gomock.Controller, count int) exchange.BidRequest {
	tmp := make([]exchange.Impression, count)
	for i := range tmp {
		s := mock_entity.NewMockImpression(c)
		s.EXPECT().ID().Return(<-random.ID).AnyTimes()
		tmp[i] = s
	}
	m := mock_entity.NewMockBidRequest(c)
	m.EXPECT().Imp().Return(tmp).AnyTimes()
	m.EXPECT().ID().Return(<-random.ID).AnyTimes()
	return m
}

func TestProviders(t *testing.T) {
	def := onion.NewDefaultLayer()
	def.SetDefault("octupos.exchange.materialize.driver", "empty")
	config.Initialize("", "", "", def)
	ctrl := gomock.NewController(t)
	broker.SetActiveBroker(mock.GetChannelBroker())

	Convey("The provider's", t, func() {
		defer ctrl.Finish()
		maximumTimeout = 50 * time.Millisecond
		Reset(func() {
			allProviders = make(map[string]providerData)
		})

		Convey("Call func", func() {

			Convey("Should return two ads", func() {

				d1 := mock_entity.NewMockDemand(ctrl)
				d1.EXPECT().WhiteListCountries().Return([]string{}).AnyTimes()
				d1.EXPECT().TestMode().Return(false).AnyTimes()

				d1.EXPECT().Name().Return("d1").AnyTimes()

				d1.EXPECT().Handicap().Return(int64(100)).AnyTimes()
				d1.EXPECT().CallRate().Return(100).AnyTimes()
				d1.EXPECT().Provide(gomock.Any(), gomock.Any(), gomock.Any()).
					Do(func(ctx context.Context, imp exchange.BidRequest, ch chan exchange.BidResponse) {
						var bmp = []exchange.Bid{}
						tmp := mock_entity.NewMockBidResponse(ctrl)
						for _, s := range imp.Imp() {
							bmpp := mock_entity.NewMockBid(ctrl)
							bmpp.EXPECT().Price().Return(int64(200)).AnyTimes()
							bmpp.EXPECT().ImpID().Return(s.ID()).AnyTimes()
							bmp = append(bmp, bmpp)

						}
						tmp.EXPECT().Bids().Return(bmp).AnyTimes()
						ch <- tmp
						close(ch)
					}).AnyTimes()
				Register(d1, time.Millisecond*100)
				bq := newBidRequest(ctrl, 2)
				logrus.Warn(len(bq.Imp()))
				bk := context.Background()
				bidResponse := Call(bk, bq)
				So(len(bidResponse), ShouldEqual, 1)
				So(len(bidResponse[0].Bids()), ShouldEqual, 2)
			})

			Convey("Should return NO ads", func() {
				d1 := mock_entity.NewMockDemand(ctrl)
				d1.EXPECT().TestMode().Return(false).AnyTimes()
				d1.EXPECT().WhiteListCountries().Return([]string{}).AnyTimes()
				d1.EXPECT().Name().Return("d1").AnyTimes()
				d1.EXPECT().Handicap().Return(int64(100)).AnyTimes()
				d1.EXPECT().CallRate().Return(100).AnyTimes()
				d1.EXPECT().Provide(gomock.Any(), gomock.Any(), gomock.Any()).
					Do(func(ctx context.Context, imp exchange.BidRequest, ch chan exchange.BidResponse) {
						time.Sleep(150 * time.Millisecond)
						var bmp = []exchange.Bid{}
						tmp := mock_entity.NewMockBidResponse(ctrl)
						for _, s := range imp.Imp() {
							bmpp := mock_entity.NewMockBid(ctrl)
							bmpp.EXPECT().Price().Return(int64(200)).AnyTimes()
							bmpp.EXPECT().ImpID().Return(s.ID()).AnyTimes()
							bmp = append(bmp, bmpp)

						}
						tmp.EXPECT().Bids().Return(bmp).AnyTimes()
						ch <- tmp
						close(ch)
					}).AnyTimes()
				Register(d1, time.Millisecond*100)
				im := newBidRequest(ctrl, 2)
				bk := context.Background()
				bidResponses := Call(bk, im)
				So(len(bidResponses), ShouldEqual, 0)

			})

			Convey("Should return one provider with three ads (timeout test)", func() {
				d1 := mock_entity.NewMockDemand(ctrl)
				d1.EXPECT().TestMode().Return(false).AnyTimes()
				d1.EXPECT().WhiteListCountries().Return([]string{}).AnyTimes()
				d1.EXPECT().Name().Return("d1").AnyTimes()
				d1.EXPECT().Handicap().Return(int64(100)).AnyTimes()
				d1.EXPECT().CallRate().Return(100).AnyTimes()
				d1.EXPECT().Provide(gomock.Any(), gomock.Any(), gomock.Any()).
					Do(func(ctx context.Context, imp exchange.BidRequest, ch chan exchange.BidResponse) {
					time.Sleep(100*time.Millisecond)
					var bmp = []exchange.Bid{}
					tmp := mock_entity.NewMockBidResponse(ctrl)
					for _, s := range imp.Imp() {
						bmpp := mock_entity.NewMockBid(ctrl)
						bmpp.EXPECT().Price().Return(int64(200)).AnyTimes()
						bmpp.EXPECT().ImpID().Return(s.ID()).AnyTimes()
						bmp = append(bmp,bmpp)

					}
					tmp.EXPECT().Bids().Return(bmp).AnyTimes()
					ch <- tmp
					close(ch)
				}).AnyTimes()
				Register(d1, time.Millisecond*100)
				d2 := mock_entity.NewMockDemand(ctrl)
				d2.EXPECT().TestMode().Return(false).AnyTimes()
				d2.EXPECT().WhiteListCountries().Return([]string{}).AnyTimes()
				d2.EXPECT().Name().Return("d2").AnyTimes()
				d2.EXPECT().Handicap().Return(int64(100)).AnyTimes()
				d2.EXPECT().CallRate().Return(100).AnyTimes()
				d2.EXPECT().Provide(gomock.Any(), gomock.Any(), gomock.Any()).
					Do(func(ctx context.Context, imp exchange.BidRequest, ch chan exchange.BidResponse) {
					time.Sleep(10*time.Millisecond)
					var bmp = []exchange.Bid{}
					tmp := mock_entity.NewMockBidResponse(ctrl)
					for _, s := range imp.Imp() {
						bmpp := mock_entity.NewMockBid(ctrl)
						bmpp.EXPECT().Price().Return(int64(200)).AnyTimes()
						bmpp.EXPECT().ImpID().Return(s.ID()).AnyTimes()
						bmp = append(bmp,bmpp)

					}
					tmp.EXPECT().Bids().Return(bmp).AnyTimes()
					ch <- tmp
					close(ch)
				}).AnyTimes()
				Register(d2, time.Millisecond*100)
				im := newBidRequest(ctrl, 3)
				bk := context.Background()

				bidResponse := Call(bk, im)
				So(len(bidResponse), ShouldEqual, 1)
				So(len(bidResponse[0].Bids()), ShouldEqual, 3)
			})

		})

		Convey("Register func", func() {

			Convey("should panic if provider (name) is NOT unique", func() {
				demand := mock_entity.NewMockDemand(ctrl)
				demand.EXPECT().TestMode().Return(false).AnyTimes()
				demand.EXPECT().Name().Return("test1").AnyTimes()
				Register(demand, time.Second*2)
				So(len(allProviders), ShouldEqual, 1)
				demand2 := mock_entity.NewMockDemand(ctrl)
				demand2.EXPECT().Name().Return("test1").AnyTimes()

				So(func() {
					Register(demand2, time.Second*2)
				}, ShouldPanic)

			})

			Convey("should register multiple providers", func() {
				demand := mock_entity.NewMockDemand(ctrl)
				demand.EXPECT().TestMode().Return(false).AnyTimes()
				demand.EXPECT().Name().Return("test1")

				Register(demand, time.Second*2)
				So(len(allProviders), ShouldEqual, 1)
				demand2 := mock_entity.NewMockDemand(ctrl)
				demand.EXPECT().TestMode().Return(false).AnyTimes()
				demand2.EXPECT().Name().Return("test2")
				So(func() {
					Register(
						demand2, time.Second*2)
				}, ShouldNotPanic)
				So(len(allProviders), ShouldEqual, 2)

			})

		})
	})

	var counter [3000]int
	skips := [...]int{1, 10, 15, 27, 35, 48, 50, 68, 79, 87, 100}

	for _, s := range skips {
		Convey(fmt.Sprintf("Skip method should return true %d out of %d times hit for %d percent call rate.", int64(float64(len(counter))*(float64(s)/100.)), len(counter), s), t, func() {
			d := mock_entity.NewMockDemand(ctrl)
			d.EXPECT().CallRate().Return(s).AnyTimes()
			p := &providerData{name: <-random.ID, provider: d, timeout: time.Second}
			var tr int64
			wg := sync.WaitGroup{}
			for range counter {
				wg.Add(1)
				go func() {
					defer wg.Done()
					if p.Skip() {
						atomic.AddInt64(&tr, 1)
					}
				}()
			}
			wg.Wait()
			So(tr, ShouldEqual, 3000-int64(float64(len(counter))*(float64(s)/100.)))

		})
	}

	Convey("Reset function should empty allProviders", t, func() {
		allProviders = make(map[string]providerData)
		allProviders["prv1"] = providerData{}
		allProviders["prv2"] = providerData{}
		ResetProviders()
		So(len(allProviders), ShouldEqual, 0)

	})

	Convey("Filters:", t, func() {

		Convey("isSameProvider function should return", func() {

			Convey("true if impression provider and provider are the same", func() {
				p2 := mock_entity.NewMockInventory(ctrl)
				p2.EXPECT().Name().Return("prv1")
				m1 := mock_entity.NewMockBidRequest(ctrl)
				m1.EXPECT().Inventory().Return(p2)
				pd := providerData{name: "prv1"}
				So(isSameProvider(m1, pd), ShouldBeTrue)
			})

			//Convey("false if impression provider and provider are NOT the same", func() {
			//	p1 := mock_entity.NewMockInventory(ctrl)
			//	p1.EXPECT().Name().Return("prv1")
			//	m1 := mock_entity.NewMockBidRequest(ctrl)
			//	m1.EXPECT().Inventory().Return(p1)
			//	pd := providerData{name: "prv2"}
			//	So(isSameProvider(m1, pd), ShouldBeFalse)
			//})
		})

		//Convey("isNotwhitelistCountries function should return", func() {
		//
		//	Convey("false if impression country is not in provider white list ", func() {
		//
		//		pr := mock_entity.NewMockDemand(ctrl)
		//		pr.EXPECT().WhiteListCountries().Return([]string{"UAE", "IRAN"}).AnyTimes()
		//		pd := providerData{provider: pr}
		//		m := mock_entity.NewMockImpression(ctrl)
		//		l := mock_entity.NewMockLocation(ctrl)
		//		l.EXPECT().Country().Return(exchange.Country{ISO: "IRAN"})
		//		m.EXPECT().Location().Return(l)
		//		So(notWhitelistCountries(m, pd), ShouldBeFalse)
		//	})
		//
		//	Convey("true if impression country is in provider white list ", func() {
		//
		//		pr := mock_entity.NewMockDemand(ctrl)
		//		pr.EXPECT().WhiteListCountries().Return([]string{"UAE", "IRAN"}).AnyTimes()
		//		pd := providerData{provider: pr}
		//		m := mock_entity.NewMockImpression(ctrl)
		//		l := mock_entity.NewMockLocation(ctrl)
		//		l.EXPECT().Country().Return(exchange.Country{ISO: "USA"})
		//		m.EXPECT().Location().Return(l)
		//		So(notWhitelistCountries(m, pd), ShouldBeTrue)
		//	})
		//})
		//
		//Convey("isExcludedDemands function should return", func() {
		//
		//	Convey("true if impression exclude provider (by name)", func() {
		//		pub := mock_entity.NewMockPublisher(ctrl)
		//		sup := mock_entity.NewMockSupplier(ctrl)
		//		sup.EXPECT().ExcludedDemands().Return([]string{"PQ", "SAME", "PSD"})
		//		pub.EXPECT().Supplier().Return(sup)
		//		m := mock_entity.NewMockImpression(ctrl)
		//		m.EXPECT().Source().Return(pub)
		//
		//		pd := providerData{name: "SAME"}
		//		So(isExcludedDemands(m, pd), ShouldBeTrue)
		//
		//	})
		//
		//	Convey("false if impression exclude provider (by name)", func() {
		//		pub := mock_entity.NewMockPublisher(ctrl)
		//		sup := mock_entity.NewMockSupplier(ctrl)
		//		sup.EXPECT().ExcludedDemands().Return([]string{"PQ", "EFG", "PSD"})
		//		pub.EXPECT().Supplier().Return(sup)
		//		m := mock_entity.NewMockImpression(ctrl)
		//		m.EXPECT().Source().Return(pub)
		//
		//		pd := providerData{name: "UNIQUE"}
		//
		//		So(isExcludedDemands(m, pd), ShouldBeFalse)
		//	})
		//})
	})
}
