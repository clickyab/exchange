package static

import (
	"github.com/clickyab/services/framework/router"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

// TODO (NOT IMPORTANT! ITS JUST AN IDEA): Add the annotation so the framework generate all this shit
type initRouter struct {
}

func (initRouter) Routes(mux *xmux.Mux, mountPoint string) {
	// Exam paths
	mux.GET(mountPoint+"/exam/pixel/:impTrackID/:slotTrackId", xhandler.HandlerFuncC(pixelHandler))
	mux.GET(mountPoint+"/exam/show/:impTrackID/:slotTrackId", xhandler.HandlerFuncC(showHandler))
	mux.GET(mountPoint+"/exam/click/:impTrackID/:slotTrackId", xhandler.HandlerFuncC(clickHandler))
	// The demand status routes
}

func init() {
	router.Register(&initRouter{})
}
