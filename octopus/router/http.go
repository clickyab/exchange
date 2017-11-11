// Package router is a glu package to mix all parts together
package router

import (
	"clickyab.com/exchange/octopus/router/internal/demands"
	"clickyab.com/exchange/octopus/router/internal/restful"

	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
)

type initRouter struct {
}

func (initRouter) Routes(mux framework.Mux) {
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
