package main

import (
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
)

type initRouter struct {
}

func (initRouter) Routes(mux framework.Mux) {

	// Exam paths
	f := mux.RootMux()
	f.GET("/fake/*static", assetHandler{})
	f.POST("/fake/ortb", ortbHandler{})
	f.POST("/fake/srtb", srtbHandler{})
}

func init() {
	router.Register(&initRouter{})
}
