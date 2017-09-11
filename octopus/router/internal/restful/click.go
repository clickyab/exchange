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
	targetURL := r.URL.Query().Get("return")

	translated := make([]byte, len([]byte(targetURL))+1)
	_, err := base64.URLEncoding.WithPadding('.').Decode(translated, []byte(targetURL))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("invalid targetURL"))
		return
	}
	targetURL = string(translated)
	sup, err := suppliers.GetSupplierByName(supplier)
	if err != nil {
		logrus.WithError(err).WithField("supplier", supplier).Debug("supplier not found")
		http.Redirect(w, r, targetURL, http.StatusFound)
		return
	}

	megaImpStore := kv.NewEavStore("SUP_CLICK_" + impTrackID + sup.Name() + trackID)

	supURL := strings.TrimSpace(megaImpStore.SubKey("SUP_URL"))
	supParam := strings.TrimSpace(megaImpStore.SubKey("SUP_PARAM"))

	if supURL == "" || supParam == "" {
		logrus.WithFields(logrus.Fields{
			"supplier":  supplier,
			"sup_url":   supURL,
			"sup_param": supParam,
		}).Debug("supplier url is invalid")
		http.Redirect(w, r, targetURL, http.StatusFound)
		return
	}

	var target = targetURL
	var b64 = make([]byte, 2*len(targetURL))
	switch sup.ClickMode() {
	case exchange.SupplierClickModeNone:
		// just redirect to user page. everything is fine
		target = targetURL
	case exchange.SupplierClickModeQueryParam:
		u, err := url.Parse(supURL)
		if err == nil {
			base64.URLEncoding.WithPadding('.').Encode(b64, []byte(targetURL))
			uq := u.Query()
			uq.Set(supParam, string(b64))
			u.RawQuery = uq.Encode()
			target = u.String()
		}
	case exchange.SupplierClickModeReplace:
		base64.URLEncoding.WithPadding('.').Encode(b64, []byte(targetURL))
		target = strings.Replace(supURL, supParam, string(b64), -1)
	case exchange.SupplierClickModeReplaceB64:
		base64.URLEncoding.Encode(b64, []byte(targetURL))
		target = strings.Replace(supURL, supParam, string(b64), -1)
	}

	if target == supURL {
		// No replace? this is bad :/
		logrus.WithFields(logrus.Fields{
			"supplier":  supplier,
			"sup_url":   supURL,
			"sup_param": supParam,
		}).Warn("supplier url is invalid, no replacement made")
		target = targetURL
	}

	logrus.WithField("supplier", supplier).WithField("action", "redirect").Debug(target)
	http.Redirect(w, r, target, http.StatusFound)
}
