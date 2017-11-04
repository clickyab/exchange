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
	mux.POST("/rest/get/:key", restful.GetAd)
	mux.GET("/click/:id/:hash", restful.Click)
	mux.GET("/pixel/:id", restful.Pixel)

	// The demand status routes
	mux.GET("/demands/status/:name", demands.Status)
	mux.POST("/demands/status/:name", demands.Status)
	mux.DELETE("/demands/status/:name", demands.Status)
	mux.PUT("/demands/status/:name", demands.Status)
	mux.OPTIONS("/demands/status/:name", demands.Status)
}

func init() {
	router.Register(&initRouter{})
}
