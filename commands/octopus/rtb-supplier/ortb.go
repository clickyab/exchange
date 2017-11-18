package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/clickyab/services/assert"
)

type ortbHandler struct {
}

func (ortbHandler) ServeHTTPC(c context.Context, w http.ResponseWriter, r *http.Request) {
	i := rtbBidRequest{}
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := dec.Decode(&i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// send request to exchange
	res, err := json.Marshal(i.Request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	g := &bytes.Buffer{}
	_, err = g.Write(res)
	assert.Nil(err)
	req, err := http.NewRequest("POST", exchangeURL+"/"+i.Meta.Key, g)
	assert.Nil(err)
	client := &http.Client{}
	resp, err := client.Do(req.WithContext(c))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)
	defer resp.Body.Close()
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}
