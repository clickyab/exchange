package restful

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"strconv"

	"clickyab.com/exchange/octopus/biding"
	"clickyab.com/exchange/octopus/dispatcher"
	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/materialize"
	"clickyab.com/exchange/octopus/suppliers"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/safe"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

func log(ctx context.Context, imp exchange.BidRequest) *logrus.Entry {
	return xlog.GetWithFields(ctx, logrus.Fields{
		"track_id": imp.ID(),
		"type":     "provider",
	})
}

func storeKeys(ctx context.Context, bq exchange.BidRequest, res exchange.BidResponse) {
	// Publish them into message broker
	for _, val := range res.Bids() {
		profit := exchange.ProfitShare(val.Price(), bq.Inventory().Supplier().Share())
		store := kv.NewEavStore(exchange.PixelPrefix + "_" + biding.GenRedisKey(ctx, bq, val))
		store.SetSubKey("supplier", bq.Inventory().Supplier().Name()).
			SetSubKey("publisher", bq.Inventory().Name()).
			SetSubKey("demand", val.Demand().Name()).
			SetSubKey("winner", strconv.FormatInt(val.Price(), 10)).
			SetSubKey("profit", strconv.FormatInt(profit, 10)).
			SetSubKey("request_time", bq.Time().String())

		assert.Nil(store.Save(time.Hour * 72))
		job := materialize.WinnerJob(bq, val)
		safe.GoRoutine(func() { broker.Publish(job) })
	}
}

func doError(w http.ResponseWriter, err error) {
	dec := json.NewEncoder(w)
	w.WriteHeader(http.StatusBadRequest)
	dec.Encode(struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

// GetAd is route to get the ad from a restful endpoint
func GetAd(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	key := xmux.Param(ctx, "key")
	sup, err := suppliers.GetSupplierByKey(key)
	if err != nil {
		doError(w, err)
		xlog.Get(ctx).WithError(err).Error("wrong supplier key")
		return
	}

	bq, err := sup.GetBidRequest(ctx, r)
	if err != nil {
		doError(w, err)
		xlog.Get(ctx).WithError(err).Error("bid request rendering issue")
		return
	}

	// OK push it to broker
	jImp := materialize.ImpressionJob(bq)
	broker.Publish(jImp)
	nCtx, cnl := context.WithCancel(ctx)
	defer cnl()

	bidResponses := dispatcher.Call(nCtx, bq)
	log(nCtx, bq).WithField("count", len(bidResponses)).Debug("bidResponses is passed the system from exchange calls")
	res := biding.SelectCPM(nCtx, bq, bidResponses)
	log(nCtx, bq).WithField("count", len(res.Bids())).Debug("bidResponses is passed the system select")
	storeKeys(ctx, bq, res)

	bq.Inventory().Supplier().RenderBidResponse(nCtx, w, res)
}
