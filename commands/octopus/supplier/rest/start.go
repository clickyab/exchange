package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"

	"io/ioutil"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/random"
)

type request struct {
	APIKey      string   `json:"api_key"`
	IP          string   `json:"ip"`
	Scheme      string   `json:"scheme"`
	PageTrackID string   `json:"page_track_id"`
	UserTrackID string   `json:"user_track_id"`
	Categories  []string `json:"categories"`
	Type        string   `json:"type"`
	UnderFloor  bool     `json:"under_floor"`
	Publisher   struct {
		Name         string `json:"name"`
		FloorCpm     int    `json:"floor_cpm"`
		SoftFloorCpm int    `json:"soft_floor_cpm"`
	} `json:"publisher"`
	Web struct {
		Referrer  string `json:"referrer"`
		Parent    string `json:"parent"`
		UserAgent string `json:"user_agent"`
	} `json:"web"`
	Slots []struct {
		Width       int               `json:"width"`
		Height      int               `json:"height"`
		TrackID     string            `json:"track_id"`
		Attributes  map[string]string `json:"attributes"`
		FallbackURL string            `json:"fallback_url"`
	} `json:"slots"`
}

var target = config.RegisterString("supplier.provider", "http://exchange.clickyab.com/", "exchange url")

var scheme = regexp.MustCompile("^https?$")

func handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	rq := &request{}
	e := d.Decode(rq)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(e.Error()))
		return
	}

	ff := target.String()
	fmt.Println(ff)
	if rq.APIKey == "" || rq.Publisher.Name == "" || rq.Publisher.FloorCpm == 0 || len(rq.Slots) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Data are not valid"))
		return
	}

	if net.ParseIP(rq.IP) == nil {
		rq.IP = framework.RealIP(r)
	}

	if !scheme.MatchString(strings.ToLower(rq.Scheme)) {
		rq.Scheme = func() string {
			if r.TLS != nil {
				return "https"
			}
			return "http"
		}()
	}
	if rq.Web.UserAgent == "" {
		rq.Web.UserAgent = r.UserAgent()
	}
	if rq.PageTrackID == "" {
		rq.PageTrackID = <-random.ID
	}
	if rq.UserTrackID == "" {
		rq.UserTrackID = <-random.ID
	}
	url := fmt.Sprintf("%sapi/rest/get/%s", target.String(), rq.APIKey)
	js, e := json.Marshal(rq)
	assert.Nil(e)

	res, e := http.Post(url, "application/json", bytes.NewBuffer(js))
	assert.Nil(e)
	w.WriteHeader(res.StatusCode)
	b, e := ioutil.ReadAll(res.Body)
	assert.Nil(e)
	defer res.Body.Close()

	f := response{Request: *rq}

	fmt.Println(string(b))
	var ex interface{}
	json.Unmarshal(b, &ex)
	f.Result = ex
	x, e := json.Marshal(f)
	assert.Nil(e)
	w.Write(x)

}

func init() {
	router.Register(&initRouter{})
}

type response struct {
	Request request     `json:"request"`
	Result  interface{} `json:"result"`
}
