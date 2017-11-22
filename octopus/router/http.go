// Package router is a glu package to mix all parts together
package router

import (
	"context"
	"net/http"

	"clickyab.com/exchange/octopus/router/internal/demands"
	"clickyab.com/exchange/octopus/router/internal/restful"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"

	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
	"github.com/rs/xmux"
)

type initRouter struct {
}

func (initRouter) Routes(mux framework.Mux) {
	mux.POST("panic", "/panic/:token", panic)

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

var token = config.RegisterString("exchange.panic_token", "", "")

func panic(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if token.String() == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("panic route is disabled"))
		return
	}
	t := xmux.Param(ctx, "token")
	if token.String() == t {
		assert.True(false)
	}

}
