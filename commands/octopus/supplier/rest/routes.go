package rest

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

type initRouter struct {
}

// Routes is the main route handler
func (initRouter) Routes(r framework.Mux) {
	r.GET("/", html)
	r.POST("/ad", handler)

}

func html(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tmp))
}
