package restful

import (
	"context"
	"net/http"

	"encoding/base64"

	"net/url"

	"strings"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/suppliers"
	"github.com/clickyab/services/kv"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

// Click is the route for click worker
func Click(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	supplier := xmux.Param(ctx, "supplier")
	trackID := xmux.Param(ctx, "trackID")
	impTrackID := xmux.Param(ctx, "impID")
	targetUrl := r.URL.Query().Get("ref")

	translated := make([]byte, len([]byte(targetUrl))+1)
	_, err := base64.URLEncoding.WithPadding('.').Decode(translated, []byte(targetUrl))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("invalid targetUrl"))
		return
	}
	targetUrl = string(translated)
	sup, err := suppliers.GetSupplierByName(supplier)
	if err != nil {
		w.Write([]byte(err.Error()))
		//http.Redirect(w, r, string(translated), http.StatusFound)
		return
	}

	megaImpStore := kv.NewEavStore("SUP_CLICK_" + impTrackID + sup.Name() + trackID)

	supURL := megaImpStore.SubKey("SUP_URL")
	supParam := megaImpStore.SubKey("SUP_PARAM")

	var target = targetUrl
	var b64 = make([]byte, 2*len(targetUrl))
	switch sup.ClickMode() {
	case exchange.SupplierClickModeNone:
		// just redirect to user page. everything is fine
		target = targetUrl
	case exchange.SupplierClickModeQueryParam:
		u, err := url.Parse(supURL)
		if err == nil {
			base64.URLEncoding.WithPadding('.').Encode(b64, []byte(targetUrl))
			uq := u.Query()
			uq.Set(supParam, string(b64))
			u.RawQuery = uq.Encode()
			target = u.String()
		}
	case exchange.SupplierClickModeReplace:
		base64.URLEncoding.WithPadding('.').Encode(b64, []byte(targetUrl))
		target = strings.Replace(supURL, supParam, string(b64), -1)
	case exchange.SupplierClickModeReplaceB64:
		base64.URLEncoding.Encode(b64, []byte(targetUrl))
		target = strings.Replace(supURL, supParam, string(b64), -1)
	}

	logrus.Debug(target)
	http.Redirect(w, r, target, http.StatusFound)
}
