package main

import (
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
)

type initRouter struct {
}

func (initRouter) Routes(mux framework.Mux) {
	// Exam paths
	mux.POST("/get/ortb", ortbHandler)
	mux.POST("/get/srtb", srtbHandler)
	mux.GET("/click/:id", clickHandler)
	//mux.POST("/show/:id", showHandler)
	mux.GET("/ad/:id", fragmentHandler)

}

func init() {
	router.Register(&initRouter{})
}
