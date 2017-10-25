package static

import (
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
)

type initRouter struct {
}

func (initRouter) Routes(mux framework.Mux) {
	// Exam paths
	mux.POST("/exam/get/ortb", ortbHandler)
}

func init() {
	router.Register(&initRouter{})
}
