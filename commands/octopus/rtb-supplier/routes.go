package main

import (
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
)

type initRouter struct {
}

func (initRouter) Routes(f framework.Mux) {

	// Exam paths
	f.GET("rtb-supplier-static", "/fake/*static", assetHandler{}.ServeHTTPC)
	f.POST("rtb-supplier-srtb", "/fake/srtb", srtbHandler{}.ServeHTTPC)

	f.POST("rtb-supplier-ortb", "/fake/ortb", ortbHandler{}.ServeHTTPC)
}

func init() {
	router.Register(&initRouter{})
}
