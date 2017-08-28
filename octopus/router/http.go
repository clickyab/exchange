// Package router is a glu package to mix all parts together
package router

import (
	"clickyab.com/exchange/octopus/router/internal/demands"
	"clickyab.com/exchange/octopus/router/internal/restful"

	"github.com/clickyab/services/framework/router"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

// TODO (NOT IMPORTANT! ITS JUST AN IDEA): Add the annotation so the framework generate all this shit
type initRouter struct {
}

func (initRouter) Routes(mux *xmux.Mux, mountPoint string) {
	mux.POST(mountPoint+"/rest/get/:key", xhandler.HandlerFuncC(restful.GetAd))
	mux.GET(mountPoint+"/click/:key", xhandler.HandlerFuncC(restful.GetAd))
	mux.GET(mountPoint+"/pixel/:demand/:trackID", xhandler.HandlerFuncC(restful.TrackPixel))

	// The demand status routes
	mux.GET(mountPoint+"/demands/status/:name", xhandler.HandlerFuncC(demands.Status))
	mux.POST(mountPoint+"/demands/status/:name", xhandler.HandlerFuncC(demands.Status))
	mux.DELETE(mountPoint+"/demands/status/:name", xhandler.HandlerFuncC(demands.Status))
	mux.HEAD(mountPoint+"/demands/status/:name", xhandler.HandlerFuncC(demands.Status))
	mux.PUT(mountPoint+"/demands/status/:name", xhandler.HandlerFuncC(demands.Status))
	mux.OPTIONS(mountPoint+"/demands/status/:name", xhandler.HandlerFuncC(demands.Status))
}

func init() {
	router.Register(&initRouter{})
}
