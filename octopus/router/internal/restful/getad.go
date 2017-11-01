package restful

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"clickyab.com/exchange/octopus/core"
	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/materialize"
	"clickyab.com/exchange/octopus/rtb"
	"clickyab.com/exchange/octopus/suppliers"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/kv"
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

func storeKeys(bq exchange.BidRequest, res exchange.BidResponse) {

	// Publish them into message broker
	for _, val := range res.Bids() {
		broker.Publish(materialize.WinnerJob(
			bq,
			val,
		))

		store := kv.NewEavStore("PIXEL_" + val.ImpID())
		store.SetSubKey("IP",
			bq.Device().IP(),
		).SetSubKey("DEMAND",
			val.Demand().Name(),
		).SetSubKey("PRICE",
			fmt.Sprintf("%d", val.Price()),
		).SetSubKey("ADID",
			val.AdID(),
		).SetSubKey("TIME",
			fmt.Sprint(bq.Time().Unix()),
		).SetSubKey("PUBLISHER",
			bq.Inventory().Publisher().Name(),
		).SetSubKey("SUPPLIER",
			bq.Inventory().Supplier().Name(),
		).SetSubKey("PROFIT",
			fmt.Sprintf("%d", int64(bq.Inventory().Supplier().Share())*val.Price()/100),
		)
		assert.Nil(store.Save(1 * time.Hour)) // TODO : Config

	}
}

// GetAd is route to get the ad from a restful endpoint
func GetAd(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	dec := json.NewEncoder(w)
	key := xmux.Param(ctx, "key")
	sup, err := suppliers.GetSupplierByKey(key)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		dec.Encode(struct {
			Error error
		}{
			Error: err,
		})
		return
	}
	bq := sup.GetBidRequest(ctx, r)
	// OK push it to broker
	jImp := materialize.ImpressionJob(bq)
	broker.Publish(jImp)
	nCtx, cnl := context.WithCancel(ctx)
	defer cnl()
	bidResponses := core.Call(nCtx, bq)
	log(nCtx, bq).WithField("count", len(bidResponses)).Debug("bidResponses is passed the system from exchange calls")
	res := rtb.SelectCPM(nCtx, bq, bidResponses)
	log(nCtx, bq).WithField("count", len(res.Bids())).Debug("bidResponses is passed the system select")
	storeKeys(bq, res)

	bq.Inventory().Supplier().RenderBidResponse(nCtx, w, res)

}
