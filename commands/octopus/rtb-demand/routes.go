package main

import (
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
)

type initRouter struct {
}

func (initRouter) Routes(mux framework.Mux) {
	// Exam paths
	mux.POST("ortb-demand-get", "/get/:name/:mode/ortb", ortbHandler)
	mux.POST("srtb-demand-get", "/get/:name/:mode/srtb", srtbHandler)
	mux.GET("rtb-demand-click", "/click/:id", clickHandler)
	mux.GET("rtb-demand-ad", "/landing/:id", adHandler)
	//mux.POST("/show/:id", showHandler)
	mux.GET("rtb-demand-show", "/ad/:id", fragmentHandler)
	mux.GET("rtb-demand-burl", "/burl/:id", billHandler)
	mux.GET("rtb-demand-nurl", "/nurl/:id", billHandler)
	mux.GET("rtb-demand-lurl", "/lurl/:id", billHandler)

}

func init() {
	router.Register(&initRouter{})
}
