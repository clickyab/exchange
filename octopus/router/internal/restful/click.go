package restful

import (
	"context"
	"net/http"

	"github.com/clickyab/services/kv"
	"github.com/rs/xmux"
)

// Click is the rouute for click worker
func Click(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	trackIDs := xmux.Param(ctx, "key")
	megaImpStore := kv.NewEavStore("MEGA_IMP_" + trackIDs)

	winURL := megaImpStore.SubKey("URL")

	// TODO: calling click worker

	http.Redirect(w, r, winURL, http.StatusTemporaryRedirect)
}
