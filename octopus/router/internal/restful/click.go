package restful

import (
	"context"
	"net/http"

	"encoding/base64"

	"encoding/json"

	"clickyab.com/exchange/octopus/suppliers"
	"github.com/clickyab/services/kv"
	"github.com/rs/xmux"
)

// Click is the route for click worker
func Click(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	supplier := xmux.Param(ctx, "supplier")
	trackID := xmux.Param(ctx, "trackID")
	url := r.FormValue("url")
	translated := make([]byte, len([]byte(url))+1)
	_, err := base64.URLEncoding.WithPadding('.').Decode(translated, []byte(url))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("invalid url"))
		return
	}

	sup, err := suppliers.GetSupplier(supplier)
	if err != nil {
		http.Redirect(w, r, string(translated), http.StatusFound)
		return
	}

	megaImpStore := kv.NewEavStore("SUP_CLICK_" + sup.Name() + trackID)

	j := json.NewEncoder(w)
	j.SetIndent("", "\t")
	j.Encode(megaImpStore.AllKeys())
}
