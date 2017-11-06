package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/simple-rtb"
)

type srtbHandler struct {
}

func (srtbHandler) ServeHTTPC(c context.Context, w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	assert.Nil(err)
	defer r.Body.Close()

	var payload = srtbBidRequest{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	assert.Nil(enc.Encode(payload.Request))

	req, err := http.NewRequest("POST", exchangeURL+"/"+payload.Meta.Key, bytes.NewBuffer(buf.Bytes()))
	assert.Nil(err)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	assert.Nil(err)

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
}
