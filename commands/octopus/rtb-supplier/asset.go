package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/bsm/openrtb"
	"github.com/clickyab/simple-rtb"
	"github.com/rs/xmux"
)

type assetHandler struct {
}

func (assetHandler) ServeHTTPC(c context.Context, w http.ResponseWriter, r *http.Request) {
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
