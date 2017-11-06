package restful

import (
	"context"
	"encoding/base64"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/materialize"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/safe"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// Click is the route for click worker
func Click(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	targetParam := r.URL.Query().Get("ref")
	if targetParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid targetURL"))
		return
	}
	target, err := base64.URLEncoding.WithPadding('.').DecodeString(targetParam)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("invalid targetURL"))
		return
	}

	id := xmux.Param(ctx, "id")
	store := kv.NewEavStore(exchange.PixelPrefix + "_" + id).AllKeys()

	if len(store) < 3 {
		http.Redirect(w, r, string(target), http.StatusFound)
		xlog.GetWithField(ctx, "click url route", "expired click url")
		return
	}

	publisher := store["publisher"]
	source := store["source"]
	demand := store["demand"]

	clickJob := materialize.ClickJob(source, publisher, demand, framework.RealIP(r))
	safe.GoRoutine(func() { broker.Publish(clickJob) })

	http.Redirect(w, r, string(target), http.StatusFound)
}
