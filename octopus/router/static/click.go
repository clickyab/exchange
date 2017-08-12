package static

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/clickyab/services/eav"
	"github.com/rs/xmux"
)

// ClickHandler for exam
func ClickHandler(ctx context.Context, w http.ResponseWriter, _ *http.Request) {
	imp := xmux.Param(ctx, "impTrackID")
	slot := xmux.Param(ctx, "slotTrackId")
	if imp == "" || slot == "" {
		logrus.Debug("both track id and demand are empty")
		return
	}
	k := eav.NewEavStore(slotKeyGen(imp, slot))
	if len(k.AllKeys()) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	o := eav.NewEavStore(fmt.Sprintf(`%s_%s`, prefixImpression, imp))
	r := o.SubKey(raw)
	if r != "" {
		r = "Original request: " + r
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Request: %s\nSlot: %s\nStatus: Successful\nTime: %s\n%s", imp, slot,
		time.Now().Format("2006-02-03 04:05:06 -07:00"), r)))
}
