package rest

import (
	"context"
	"net/http"

	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

type initRouter struct {
}

// Routes is the main route handler
func (initRouter) Routes(r *xmux.Mux, _ string) {
	r.GET("/", xhandler.HandlerFuncC(html))
	r.POST("/ad", xhandler.HandlerFuncC(handler))

}

func html(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tmp))
}
