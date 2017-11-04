package srtb

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"

	"clickyab.com/exchange/octopus/demands/internal/base"
	"clickyab.com/exchange/octopus/exchange/srtb"
	"github.com/clickyab/services/xlog"
	simple "github.com/clickyab/simple-rtb"
)

// Demand srtb demand structure
type Demand struct {
	exchange.DemandBase
}

// Provide method for demand
func (d *Demand) Provide(ctx context.Context, bq exchange.BidRequest, ch chan exchange.BidResponse) {
	base.Provide(ctx, d, bq, ch)
}

// GetBidResponse try to get bidresponse from response
func (d *Demand) GetBidResponse(ctx context.Context, resp *http.Response, sup exchange.Supplier) (exchange.BidResponse, error) {
	t := &bytes.Buffer{}
	_, err := io.Copy(t, resp.Body)
	assert.Nil(err)

	defer resp.Body.Close()

	p := t.Bytes()
	xlog.Get(ctx).WithField("key", d.Name()).WithField("result", string(p)).Debug("Call done")

	res := &simple.BidResponse{}
	err = json.Unmarshal(p, res)
	if err != nil {
		return nil, err
	}
	return srtb.NewBidResponse(d, sup, res), nil
}

// RenderBidRequest cast bid request to ortb
func (d *Demand) RenderBidRequest(ctx context.Context, w io.Writer, bq exchange.BidRequest) http.Header {
	j := json.NewEncoder(w)
	// render in rtb style
	r, err := srtb.NewSimpleRTBFromBidRequest(bq)
	assert.Nil(err)
	err = j.Encode(r)
	assert.Nil(err)
	// TODO : simple rtb headers
	return http.Header{}
}
