package restful

import (
	"context"
	"core"
	"encoding/base64"
	"net/http"
	"services/assert"
	"services/eav"
	"strconv"

	"github.com/fzerorubigd/xmux"
)

const message = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAAA1BMVEUAAACnej3aAAAAAXRSTlMAQObYZgAAAApJREFUCNdjYAAAAAIAAeIhvDMAAAAASUVORK5CYII="

var data []byte

func trackPixel(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(data)
	go func() {
		demand := xmux.Param(ctx, "demand")
		trackID := xmux.Param(ctx, "trackID")
		if trackID == "" || demand == "" {
			return
		}
		//get from store
		store := eav.NewEavStore(trackID)
		winnerDemand := store.SubKey("DEMAND")
		winnerID := store.SubKey("ID")
		winnerBID := store.SubKey("BID")
		winnerInt, err := strconv.ParseInt(winnerBID, 10, 0)
		if err != nil {
			return
		}
		//set winner
		d := core.GetDemand(winnerDemand)
		d.Win(ctx, winnerID, winnerInt)
	}()
}

func init() {
	var err error
	data, err = base64.StdEncoding.DecodeString(message)
	assert.Nil(err)
}
