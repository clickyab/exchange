package restful

import (
	"context"
	"net/http"

	"strconv"
	"time"

	"errors"

	"clickyab.com/exchange/octopus/exchange/materialize"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/safe"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

var pixel = `iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEUAAACnej3aAAAAAXRSTlMAQObYZgAAAApJREFUCNdjYAAAAAIAAeIhvDMAAAAASUVORK5CYII=`

// Pixel route handles the show job system
func Pixel(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "image/png")
	w.Write([]byte(pixel))

	safe.GoRoutine(func() {
		id := xmux.Param(ctx, "id")
		store := kv.NewEavStore(id).AllKeys()
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
