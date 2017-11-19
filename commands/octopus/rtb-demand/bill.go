package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/clickyab/services/xlog"
)

func billHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	xlog.Get(ctx).Debug(fmt.Sprintf("route %s", r.URL.Path))
	// Add logic if needed
	w.Write([]byte("got it"))
}
