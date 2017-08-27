package rest

import (
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

type initRouter struct {
}

// Routes is the main route handler
func (initRouter) Routes(r *xmux.Mux, _ string) {
	r.POST("/ad", xhandler.HandlerFuncC(handler))
}
