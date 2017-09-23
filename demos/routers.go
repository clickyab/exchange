package demos

import (
	"github.com/clickyab/services/framework/router"
	"github.com/rs/xmux"
)

type initRouter struct {
}

func (initRouter) Routes(mux *xmux.Mux, mountPoint string) {
}

func init() {
	router.Register(&initRouter{})
}
