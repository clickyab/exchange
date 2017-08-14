package static

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/Sirupsen/logrus"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"
	"github.com/rs/xmux"
)

var showAd = `
<style>
.cyb-cnt {
display: block;
display: flex;
width: 100px;
background:#d00;
color:#eee;
}
.landing {
text-align: center;
text-decoration: none;
font-weight: 500;
}
</style>
<div class="cyb-cnt">
<a class="landing" href="{{.URL}}">
{{.Message}}
</a>
</div>

`

// showHandler handle show url for exam
func showHandler(ctx context.Context, w http.ResponseWriter, _ *http.Request) {
	imp := xmux.Param(ctx, "impTrackID")
	slot := xmux.Param(ctx, "slotTrackId")
	if imp == "" || slot == "" {
		logrus.Debug("both track id and demand are empty")
		return
	}
	k := kv.NewEavStore(slotKeyGen(imp, slot))
	if len(k.AllKeys()) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write(filler("#", "NOT FOUND"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(filler(k.SubKey(clickURL), fmt.Sprintf("Slot ID: %s", slot)))
}

func filler(u, m string) []byte {
	at := template.Template{}
	t, e := at.Parse(showAd)
	assert.Nil(e)
	b := &bytes.Buffer{}
	t.Execute(b, struct {
		Message,
		URL string
	}{
		m,
		u,
	})
	return b.Bytes()
}
