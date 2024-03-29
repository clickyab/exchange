package restful

import (
	"context"
	"net/http"

	"strconv"

	"errors"

	"clickyab.com/exchange/octopus/biding"
	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/materialize"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/safe"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

var pixel = `iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEUAAACnej3aAAAAAXRSTlMAQObYZgAAAApJREFUCNdjYAAAAAIAAeIhvDMAAAAASUVORK5CYII=`

const (
	showJS    = "show.js"
	showPixel = "image.png"
)

// Show route handles the show job system
func Show(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	t := xmux.Param(ctx, "type")
	if t == showJS {
		w.Header().Set("content-type", "text/javascript")
		w.Write([]byte(""))
	} else if t == showPixel {
		w.Header().Set("content-type", "image/png")
		w.Write([]byte(pixel))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	safe.GoRoutine(func() {
		id := xmux.Param(ctx, "id")
		store := kv.NewEavStore(exchange.PixelPrefix + "_" + id).AllKeys()
		if len(store) < 6 {
			xlog.GetWithError(ctx, errors.New("session expired for this id"))
			return
		}

		demand := store["demand"]
		supplier := store["supplier"]
		publisher := store["publisher"]
		rqTime := store["request_time"]
		billURL := store["bill_url"]

		profit := store["profit"]
		profitFloat, err := strconv.ParseFloat(profit, 64)
		if err != nil {
			xlog.GetWithError(ctx, errors.New("non int value for profit was set")).Panicln()
			return
		}

		winner := store["winner"]
		winnerFloat, err := strconv.ParseFloat(winner, 64)
		if err != nil {
			xlog.GetWithError(ctx, errors.New("non int value for winner was set")).Panicln()
			return
		}

		//call bill url
		d := &http.Client{}
		biding.DoBillGetRequest(ctx, d, billURL)
		showJob := materialize.ShowJob(demand, framework.RealIP(r), winnerFloat, rqTime, supplier, publisher, profitFloat)
		broker.Publish(showJob)
	})
}
