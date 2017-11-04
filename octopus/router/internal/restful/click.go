package restful

import (
	"context"
	"encoding/base64"
	"net/http"

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
	hashParam := xmux.Param(ctx, "hash")
	var hash []byte
	_, err := base64.URLEncoding.WithPadding('.').Decode(hash, []byte(hashParam))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("invalid targetURL"))
		return
	}

	id := xmux.Param(ctx, "id")
	store := kv.NewEavStore(id).AllKeys()

	if len(store) < 3 {
		http.Redirect(w, r, string(hash), http.StatusFound)
		xlog.GetWithField(ctx, "click url route", "expired click url")
		return
	}

	publisher := store["publisher"]
	source := store["source"]
	demand := store["demand"]

	clickJob := materialize.ClickJob(source, publisher, demand, framework.RealIP(r))
	safe.GoRoutine(func() { broker.Publish(clickJob) })

	http.Redirect(w, r, string(hash), http.StatusFound)
}
