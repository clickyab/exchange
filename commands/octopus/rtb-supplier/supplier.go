package main

import (
	"context"
	"net/http"

	"clickyab.com/exchange/octopus/srtb"

	"strings"

	"encoding/json"

	"bytes"

	"io/ioutil"

	"fmt"

	"github.com/bsm/openrtb"
	"github.com/clickyab/services/assert"
	"github.com/rs/xmux"
)

type s struct {
}

type d struct {
}

type e struct {
}

var exchangeURL = "http://127.0.0.1:8090/api/rest/get"
var prefix = "commands/octopus/rtb-supplier/static/template"

func (s) ServeHTTPC(c context.Context, w http.ResponseWriter, r *http.Request) {
	g := xmux.Param(c, "static")
	for key := range _bindata {
		fmt.Println(key)
	}
	fmt.Println(g)
	if g == "/" {
		_, x := Asset("commands/octopus/rtb-supplier/static/template/index.html")
		fmt.Println(x)
		b, e := Asset(prefix + "/index.html")
		println(e)
		if e != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(b)
		return
	}
	g = g[5:]

	b, e := Asset(prefix + g)
	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if strings.HasSuffix(g, ".js") {
		w.Header().Set("Content-Type", "text/javascript")
	} else if strings.HasSuffix(g, ".css") {
		w.Header().Set("Content-Type", "text/css")
	}
	w.Write(b)
	return
}

type rtbBidRequest struct {
	Request openrtb.BidRequest `json:"request"`
	Meta    struct {
		Key string `json:"key"`
	} `json:"meta"`
}

type srtbBidRequest struct {
	Request srtb.BidRequest `json:"request"`
	Meta    struct {
		Key string `json:"key"`
	} `json:"meta"`
}

func (d) ServeHTTPC(c context.Context, w http.ResponseWriter, r *http.Request) {
	i := rtbBidRequest{}
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := dec.Decode(&i)
	if err != nil {
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
	req, err := http.NewRequest("POST", exchangeURL+"/"+i.Meta.Key, bytes.NewBuffer(g.Bytes()))
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

func (e) ServeHTTPC(c context.Context, w http.ResponseWriter, r *http.Request) {
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
		w.Write(data2)
		return
	}
	var result = &srtb.BidResponse{}
	err = json.Unmarshal(data2, &result)
	assert.Nil(err)
	enc1 := json.NewEncoder(w)
	enc1.Encode(result)
}
