// Package router is a glu package to mix all parts together
package router

import (
	"context"
	"net/http"

	"clickyab.com/exchange/octopus/router/internal/demands"
	"clickyab.com/exchange/octopus/router/internal/restful"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/hub"

	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
)

type initRouter struct {
}

func (initRouter) Routes(mux framework.Mux) {
	mux.POST("panic", "/panic", doPanic)
	mux.POST("reload", "/reload", reload)

	mux.POST("get-ad", "/rest/get/:key", restful.GetAd)
	mux.GET("click", "/click/:id", restful.Click)
	mux.GET("pixel", "/show/:id/:type", restful.Show)

	// The demand status routes
	mux.GET("demand-status-get", "/demands/status/:name", demands.Status)
	mux.POST("demand-status-post", "/demands/status/:name", demands.Status)
	mux.DELETE("demand-status-del", "/demands/status/:name", demands.Status)
	mux.PUT("demand-status-put", "/demands/status/:name", demands.Status)
	mux.OPTIONS("demand-status-opt", "/demands/status/:name", demands.Status)
}

func init() {
	router.Register(&initRouter{})
}

var panicToken = config.RegisterString("octopus.panic_token", "", "do panic")
var reloadToken = config.RegisterString("octopus.reload_token", "", "reload config")

func reload(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if reloadToken.String() == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("reload route is disabled"))
		return
	}
	t := r.Header.Get("token")
	if reloadToken.String() == t {
		hub.Publish("reload", nil)
	}
}

func doPanic(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if panicToken.String() == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("panic route is disabled"))
		return
	}
	t := r.Header.Get("token")
	if panicToken.String() == t {
		panic("DON'T PANIC, IT'S JUST A TEST")
	}

}
