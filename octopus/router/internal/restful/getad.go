package restful

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"clickyab.com/exchange/octopus/core"
	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/materialize"
	"clickyab.com/exchange/octopus/rtb"
	"clickyab.com/exchange/octopus/suppliers"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/kv"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

var host = config.RegisterString("octopus.host.name", "exchange-dev.3rdad.com", "the exchange root")

func log(imp exchange.Impression) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"track_id": imp.TrackID(),
		"type":     "provider",
	})
}

// GetAd is route to get the ad from a restful endpoint
func GetAd(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	dec := json.NewEncoder(w)
	key := xmux.Param(ctx, "key")

	imp, err := suppliers.GetImpression(key, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		dec.Encode(struct {
			Error error
		}{
			Error: err,
		})
		return
	}

	// Change the click url
	for _, s := range imp.Slots() {
		att := s.Attributes()
		if att == nil {
			continue
		}
		exchangeClickURL := &url.URL{
			Scheme: imp.Scheme(),
			Host:   host.String(),
			Path:   fmt.Sprintf("/click/%s/%s", imp.Source().Supplier().Name(), s.TrackID()),
		}
		s.SetAttribute("_click_url", att["click_url"])
		s.SetAttribute("_click_parameter", att["_click_parameter"])
		s.SetAttribute("click_parameter", "ref")
		s.SetAttribute("click_url", exchangeClickURL.String())
		s.SetAttribute("type", "parameter")
	}

	// OK push it to broker
	jImp := materialize.ImpressionJob(imp)
	broker.Publish(jImp)
	nCtx, cnl := context.WithCancel(ctx)
	defer cnl()
	ads := core.Call(nCtx, imp)
	log(imp).WithField("count", len(ads)).Debug("ads is passed the system from exchange calls")
	res := rtb.SelectCPM(imp, ads)
	log(imp).WithField("count", len(res)).Debug("ads is passed the system select")
	for _, s := range imp.Slots() {
		i := s.TrackID()
		// Publish them into message broker
		if res[i] != nil {
			broker.Publish(materialize.WinnerJob(
				imp,
				res[i],
				i,
			))
			att := s.Attributes()
			assert.NotNil(att)

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
			assert.Nil(store.Save(1 * time.Hour)) // TODO : Config

			megaImpStore := kv.NewEavStore("SUP_CLICK_" + imp.Source().Supplier().Name() + res[i].TrackID())
			megaImpStore.SetSubKey("SUP_URL",
				att["_click_url"],
			).SetSubKey("SUP_PARAM",
				att["_click_parameter"],
			)
			assert.Nil(megaImpStore.Save(72 * time.Hour)) // TODO : Config
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
