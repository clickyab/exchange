package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/xlog"
	"github.com/clickyab/simple-rtb"
)

type srtbHandler struct {
}

func (srtbHandler) ServeHTTPC(c context.Context, w http.ResponseWriter, r *http.Request) {
	var payload = &srtbBidRequest{}
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := dec.Decode(payload)

	if err != nil {
		xlog.GetWithError(c, err).Debug("unmarshal error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// send request to exchange
	res, err := json.Marshal(payload.Request)
	if err != nil {
		xlog.GetWithError(c, err).Debug("marshal error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	g := &bytes.Buffer{}
	_, err = g.Write(res)
	assert.Nil(err)
	req, err := http.NewRequest("POST", exchangeURL.String()+"/"+payload.Meta.Key, g)
	assert.Nil(err)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		xlog.GetWithError(c, err).Debug("request failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data2, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte("Exchange error"))
		return
	}
	var result = &srtb.BidResponse{}
	err = json.Unmarshal(data2, &result)

	assert.Nil(err)
	enc1 := json.NewEncoder(w)
	enc1.Encode(result)
	return
}
