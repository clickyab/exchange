package biding

import (
	"sort"
	"time"

	"context"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/kv"
)

var (
	pageLock = config.RegisterDuration(
		"exchange.rtb.lock_duration",
		500*time.Millisecond,
		"the time to lock one page for other requests",
	)
	pageLifeTime = config.RegisterDuration(
		"exchange.rtb.page_lifetime",
		5*time.Minute,
		"the lifetime of the page key in redis to prevent duplicate ad in one page",
	)
)

// SelectCPM is the simplest way to bid. sort the value, return the
func SelectCPM(ctx context.Context, bq exchange.BidRequest, all []exchange.BidResponse) exchange.BidResponse {
	bids := make([]exchange.Bid, 0)
	lock := kv.NewDistributedLock("LOCK"+bq.Inventory().Supplier().Name()+bq.ID(), pageLock.Duration())
	lock.Lock()
	defer lock.Unlock()
	set := kv.NewDistributedSet("EXC" + bq.Inventory().Supplier().Name() + bq.ID())
	for _, m := range bq.Imp() {
		reds := reduce(m, all, bq.Inventory().Supplier())
		sorted := sortedBid(rmDuplicate(set, reds))
		if len(sorted) == 0 {
			continue
		}
		sort.Sort(sorted)

		tb := sorted[0]
		var tp int64
		lower := bq.Inventory().Supplier().SoftFloorCPM()
		if lower > tb.Price() {
			lower = int64(m.BidFloor())
		}
		if len(sorted) > 1 && sorted[1].Price() > lower {
			lower = sorted[1].Price()
		}
		if lower < tb.Price() {
			tp = lower + 1.0
		} else {
			tp = tb.Price()
		}
		set.Add(tb.AdID())

		rep := replacer(ctx, bq, tb)
		rb := bid{
			markup:  rep.Replace(tb.AdMarkup()),
			winurl:  rep.Replace(tb.WinURL()),
			billurl: rep.Replace(tb.BillURL()),
			bid:     tb,
			price:   int64(exchange.DecShare(float64(tp), bq.Inventory().Supplier().Share())),
		}
		if !hasTracker(tb) {
			rb.Demand().Bill(ctx, rb)
		}
		bids = append(bids, rb)
	}
	var res = rsp{
		supplier:   bq.Inventory().Supplier(),
		id:         bq.ID(),
		attributes: bq.Attributes(),
	}
	if len(bids) == 0 {
		res.excuse = exchange.ExcuseNoDemandWantThisBid
	} else {
		res.bids = bids
	}
	set.Save(pageLifeTime.Duration())
	return res
}

func reduce(bq exchange.Impression, b []exchange.BidResponse, s exchange.Supplier) []exchange.Bid {
	imp := bq.ID()

	res := make([]exchange.Bid, 0)
	for _, br := range b {
		for _, bid := range br.Bids() {
			if bid.ImpID() == imp && bid.Price() >= int64(bq.BidFloor()) {
				res = append(res, bid)
				break
			}
		}
	}
	return res
}

func rmDuplicate(set kv.DistributedSet, ads []exchange.Bid) []exchange.Bid {
	all := set.Members()
	var res []exchange.Bid
bigLoop:
	for id := range ads {
		for _, adID := range all {
			if adID == ads[id].ID() {
				continue bigLoop
			}
		}
		res = append(res, ads[id])
	}
	return res
}
