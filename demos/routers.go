package demos

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

type initRouter struct {
}

func (initRouter) Routes(mux *xmux.Mux, mountPoint string) {
	mux.POST("/", xhandler.HandlerFuncC(getFakeResponse))
}

func getFakeResponse(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	request, err := bidRequestParser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := bidFakeResponse(request)
	framework.JSON(w, http.StatusOK, response)
}

func init() {
	router.Register(&initRouter{})
}
