package ortb

import (
	"context"
	"net/http"

	"io"

	"encoding/json"

	"bytes"

	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/ortb"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/xlog"
)

// GetBidResponse try to get bidresponse from response
func GetBidResponse(ctx context.Context, d exchange.Demand, r *http.Response, s exchange.Supplier) (exchange.BidResponse, error) {
	t := &bytes.Buffer{}
	_, err := io.Copy(t, r.Body)
	assert.Nil(err)

	defer r.Body.Close()

	p := t.Bytes()
	xlog.GetWithField(ctx, "key", d.Name()).WithField("result", string(p)).Debug("Call done")

	res := ortb.NewBidResponse(d, s)
	err = json.Unmarshal(p, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RenderBidRequest cast bid request to ortb
func RenderBidRequest(ctx context.Context, d exchange.Demand, w io.Writer, bq exchange.BidRequest) http.Header {
	o, e := ortb.NewOpenRTBFromBidRequest(bq)
	assert.Nil(e)
	err := json.NewEncoder(w).Encode(o)
	assert.Nil(err)
	// TODO : Add open rtb headers
	return http.Header{}

}
