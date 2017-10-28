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
	f.GET("/fake/*ggg", s{})
	f.POST("/fake/rtb", d{})
	f.POST("/fake/srtb", e{})
}

func init() {
	router.Register(&initRouter{})
}
