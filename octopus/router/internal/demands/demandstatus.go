package demands

import (
	"context"
	"net/http"

	"clickyab.com/exchange/octopus/dispatcher"
	"github.com/rs/xmux"
)

// Status is the demand status route
func Status(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := xmux.Param(ctx, "name")
	demand, err := dispatcher.GetDemand(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	demand.Status(ctx, w, r)
}
