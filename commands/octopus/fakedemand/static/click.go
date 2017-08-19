package static

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/clickyab/services/kv"
	"github.com/rs/xmux"
)

// clickHandler for exam
func clickHandler(ctx context.Context, w http.ResponseWriter, _ *http.Request) {
	imp := xmux.Param(ctx, "impTrackID")
	slot := xmux.Param(ctx, "slotTrackId")
	if imp == "" || slot == "" {
		logrus.Debug("both track id and demand are empty")
		return
	}
	k := kv.NewEavStore(slotKeyGen(imp, slot))
	if len(k.AllKeys()) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	o := kv.NewEavStore(fmt.Sprintf(`%s_%s`, prefixImpression, imp))
	r := o.SubKey(raw)
	q := o.SubKey(org)
	if q != "" {
		q = "Original request: " + q
	}
	if r != "" {
		r = "Response: " + r
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Request: %s\nSlot: %s\nStatus: Successful\nTime: %s\n%s\n\n%s", imp, slot,
		time.Now().Format("2006-02-03 04:05:06 -07:00"), r, q)))
}