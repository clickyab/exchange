package static

import (
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
)

type initRouter struct {
}

func (initRouter) Routes(mux framework.Mux) {
	// Exam paths
	mux.POST("/exam/get/ad", demandHandler)
	mux.GET("/exam/ad/:impTrackID/:slotTrackId/:width/:height", adHandler)
	mux.GET("/exam/click/:impTrackID/:slotTrackId", clickHandler)
}

func init() {
	router.Register(&initRouter{})
}
