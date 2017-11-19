package main

import (
	"context"
	"net/http"
	"net/url"

	"github.com/clickyab/services/framework/router"
	"github.com/rs/xmux"
)

func clickHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u := url.URL{
		Host: r.Host,
		Path: router.MustPath("rtb-demand-ad", map[string]string{"id": xmux.Param(ctx, "id")}),
		Scheme: func() string {
			if r.TLS != nil {
				return "https"
			}
			return "http"
		}(),
	}

	http.Redirect(w, r, u.String(), http.StatusTemporaryRedirect)
}
