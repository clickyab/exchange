package static

import (
	"github.com/clickyab/services/framework/router"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

type initRouter struct {
}

func (initRouter) Routes(mux *xmux.Mux, _ string) {
	// Exam paths

	mux.POST("/exam/get/ad", xhandler.HandlerFuncC(demandHandler))
	mux.GET("/exam/ad/:impTrackID/:slotTrackId", xhandler.HandlerFuncC(adHandler))
	mux.GET("/exam/pixel/:impTrackID/:slotTrackId", xhandler.HandlerFuncC(pixelHandler))
	mux.GET("/exam/show/:impTrackID/:slotTrackId", xhandler.HandlerFuncC(showHandler))
	mux.GET("/exam/click/:impTrackID/:slotTrackId", xhandler.HandlerFuncC(clickHandler))
}

func init() {
	router.Register(&initRouter{})
}
