package biding

import (
	"sort"
	"time"

	"context"

	"net/http"
	"net/url"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/safe"
	"github.com/clickyab/services/xlog"
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
	ref := bq.User().ID()
	if ref == "" {
		ref = bq.ID()
	}

	bids := make([]exchange.Bid, 0)
	lock := kv.NewDistributedLock("LOCK"+ref, pageLock.Duration())
	lock.Lock()
	defer lock.Unlock()
	set := kv.NewDistributedSet("EXC" + bq.Inventory().Supplier().Name() + ref)
	for _, m := range bq.Imp() {
		reds := reduce(m, all, bq.Inventory().Supplier())
		sorted := sortedBid(rmDuplicate(set, reds))
		if len(sorted) == 0 {
			continue
		}
		sort.Sort(sorted)

		tb := sorted[0]
		var tp float64
		lower := bq.Inventory().Supplier().SoftFloorCPM()
		if lower < m.BidFloor() {
			lower = m.BidFloor()
		}

		if lower > tb.Price() {
			lower = m.BidFloor()
		}
		if len(sorted) > 1 && sorted[1].Price() > lower {
			lower = sorted[1].Price()
		}
		if lower < tb.Price() {
			tp = lower
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
			price:   exchange.DecShare(tp, bq.Inventory().Supplier().Share()),
		}
		if !hasTracker(tb) {
			rb.Demand().Bill(ctx, tp, rb.billurl)
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

func reduce(m exchange.Impression, b []exchange.BidResponse, s exchange.Supplier) []exchange.Bid {
	imp := m.ID()
	res := make([]exchange.Bid, 0)
	for x := range b {
		for i := range b[x].Bids() {
			bid := b[x].Bids()[i]
			if bid.ImpID() == imp && float64(exchange.DecShare(bid.Price(), s.Share())) >= m.BidFloor() {
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

// DoBillGetRequest make request to given url
func DoBillGetRequest(ctx context.Context, c *http.Client, hit string) {
	safe.GoRoutine(func() {
		u, err := url.Parse(hit)
		if err != nil {
			xlog.SetField(ctx, "bid bill url is not valid", err)
			return
		}
		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			xlog.SetField(ctx, "demand making bill request failure", err)
			return
		}
		_, err = c.Do(req.WithContext(ctx))
		if err != nil {
			xlog.SetField(ctx, "demand making bill request failure", err)
			return
		}
	})
}
