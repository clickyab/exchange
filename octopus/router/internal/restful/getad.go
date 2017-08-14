package restful

import (
	"context"
	"encoding/json"
	"net/http"

	"clickyab.com/exchange/octopus/core"
	"clickyab.com/exchange/octopus/rtb"
	"clickyab.com/exchange/octopus/supliers"

	"fmt"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"

	"clickyab.com/exchange/octopus/exchange/materialize"
	"github.com/clickyab/services/broker"

	"strings"

	"clickyab.com/exchange/commands/octopus/fakedemand/static"
	"github.com/Sirupsen/logrus"
	"github.com/clickyab/services/config"
	"github.com/rs/xmux"
)

var testKeys = config.RegisterString("rest.test.keys", "", "")

// GetAd is route to get the ad from a restful endpoint
func GetAd(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	dec := json.NewEncoder(w)
	key := xmux.Param(ctx, "key")
	logrus.Warn(testKeys)
	for _, v := range strings.Split(testKeys.String(), ",") {
		if v == key {
			static.DemandHandler(ctx, w, r)
			return
		}
	}

	imp, err := supliers.GetImpression(key, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		dec.Encode(struct {
			Error error
		}{
			Error: err,
		})
		return
	}
	// OK push it to broker
	jImp := materialize.ImpressionJob(imp)
	broker.Publish(jImp)
	nCtx, cnl := context.WithCancel(ctx)
	defer cnl()
	ads := core.Call(nCtx, imp)
	logrus.Debugf("%d ads is passed the system from exchange calls", len(ads))
	res := rtb.SelectCPM(imp, ads)
	logrus.Debugf("%d ads is passed the system select", len(res))
	// Publish them into message broker
	for i := range res {
		// TODO : why this happen?
		if res[i] != nil {
			broker.Publish(materialize.WinnerJob(
				imp,
				res[i],
				i,
			))

			store := kv.NewEavStore("PIXEL_" + res[i].TrackID())
			store.SetSubKey("IP",
				imp.IP().String(),
			).SetSubKey("DEMAND",
				res[i].Demand().Name(),
			).SetSubKey("BID",
				fmt.Sprintf("%d", res[i].WinnerCPM()),
			).SetSubKey("ADID",
				res[i].ID(),
			).SetSubKey("TIME",
				fmt.Sprint(imp.Time().Unix()),
			).SetSubKey("PUBLISHER",
				imp.Source().Name(),
			).SetSubKey("SUPPLIER",
				imp.Source().Supplier().Name(),
			).SetSubKey("PROFIT",
				fmt.Sprintf("%d", int64(imp.Source().Supplier().Share())*res[i].WinnerCPM()/100),
			)
			assert.Nil(store.Save(1 * time.Hour))
		}
	}
	err = imp.Source().Supplier().Renderer().Render(imp, res, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		dec.Encode(struct {
			Error string
		}{
			Error: err.Error(),
		})
	}
}
