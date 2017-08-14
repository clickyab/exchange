package static

import (
	"context"
	"encoding/base64"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"
	"github.com/rs/xmux"
)

// pixelHandler return an one by one pixel
func pixelHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(data)
	imp := xmux.Param(ctx, "impTrackID")
	slot := xmux.Param(ctx, "slotTrackId")
	if imp == "" || slot == "" {
		logrus.Debug("both track id and demand are empty")
		return
	}
	k := kv.NewEavStore(slotKeyGen(imp, slot))
	if len(k.AllKeys()) == 0 {
		return
	}

}

const message = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEUAAACnej3aAAAAAXRSTlMAQObYZgAAAApJREFUCNdjYAAAAAIAAeIhvDMAAAAASUVORK5CYII="

var data []byte

func init() {
	var err error
	data, err = base64.StdEncoding.DecodeString(message)
	assert.Nil(err)
}
