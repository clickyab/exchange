package restful

import (
	"context"
	"net/http"

	"strconv"
	"time"

	"errors"

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
		if len(store) < 5 {
			xlog.GetWithError(ctx, errors.New("session expired for this id"))
			return
		}

		demand := store["demand"]
		supplier := store["supplier"]
		publisher := store["publisher"]

		profit := store["profit"]
		profitInt, err := strconv.ParseInt(profit, 10, 0)
		if err != nil {
			xlog.GetWithError(ctx, errors.New("non int value for profit was set")).Panicln()
			return
		}

		winner := store["winner"]
		winnerInt, err := strconv.ParseInt(winner, 10, 0)
		if err != nil {
			xlog.GetWithError(ctx, errors.New("non int value for winner was set")).Panicln()
			return
		}

		showJob := materialize.ShowJob(demand, framework.RealIP(r), winnerInt, time.Now().String(), supplier, publisher, profitInt)
		broker.Publish(showJob)
	})
}
