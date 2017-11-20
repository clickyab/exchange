package srtb

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"

	"clickyab.com/exchange/octopus/exchange/srtb"
	"github.com/clickyab/services/xlog"
	simple "github.com/clickyab/simple-rtb"
)

// GetBidResponse try to get bidresponse from response
func GetBidResponse(ctx context.Context, d exchange.Demand, r *http.Response, sup exchange.Supplier) (exchange.BidResponse, error) {
	t := &bytes.Buffer{}
	_, err := io.Copy(t, r.Body)
	assert.Nil(err)

	defer r.Body.Close()

	p := t.Bytes()
	xlog.GetWithField(ctx, "key", d.Name()).WithField("result", string(p)).Debug("Call done")

	res := &simple.BidResponse{}
	err = json.Unmarshal(p, res)
	if err != nil {
		return nil, err
	}
	return srtb.NewBidResponse(d, sup, res), nil
}

// RenderBidRequest cast bid request to ortb
func RenderBidRequest(ctx context.Context, d exchange.Demand, w io.Writer, bq exchange.BidRequest) http.Header {
	j := json.NewEncoder(w)
	// render in rtb style
	r, err := srtb.NewSimpleRTBFromBidRequest(bq)
	assert.Nil(err)
	err = j.Encode(r)
	assert.Nil(err)
	// TODO : simple rtb headers
	return http.Header{}
}
