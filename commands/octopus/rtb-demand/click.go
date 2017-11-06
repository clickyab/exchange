package main

import (
	"context"
	"net/http"

	"github.com/rs/xmux"
)

func clickHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://google.com/?aid="+xmux.Param(ctx, "id"), http.StatusTemporaryRedirect)
}
