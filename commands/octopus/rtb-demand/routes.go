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
	//mux.POST("/show/:id", showHandler)
	mux.GET("rtb-demand-show", "/ad/:id", fragmentHandler)

}

func init() {
	router.Register(&initRouter{})
}
