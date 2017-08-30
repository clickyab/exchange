package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/clickyab/services/assert"
	"github.com/sirupsen/logrus"
)

const typeWeb = "web"

type categories string

func getSupplierDemo(w http.ResponseWriter, r *http.Request) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(dir)
	t, err := template.New("index").Parse(index)
	assert.Nil(err)
	t.Execute(w, nil)
}

type dumbAd struct {
	TrackID  string `json:"track_id"`
	Winner   int64  `json:"winner"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Code     string `json:"code"`
	Landing  string `json:"landing"`
	IsFilled bool   `json:"is_filled"`
}

func postSupplierDemo(w http.ResponseWriter, r *http.Request) {
	var respon []dumbAd
	var cats []categories
	var resSlots = make([]*slotRest, 0)
	if r.Method != "POST" {
		return
	}
	r.ParseForm()
	data := r.Form
	rType := data["type"][0]
	if rType != typeWeb {
		rType = typeWeb
	}
	refferer := data["refferer"][0]
	parent := data["parent"][0]
	userAgent := data["user_agent"][0]
	if userAgent == "" {
		userAgent = r.UserAgent()
	}
	res := requestBody{}
	for i := range data["width[]"] {
		if data["width[]"][i] != "" {
			wi, _ := strconv.ParseInt(data["width[]"][i], 10, 0)
			he, _ := strconv.ParseInt(data["height[]"][i], 10, 0)
			resSlots = append(resSlots, &slotRest{
				W:    int(wi),
				H:    int(he),
				TID:  data["track[]"][i],
				Fall: "example.com",
			})
		}

	}
	if data["categories"][0] != "" {
		category := strings.Split(data["categories"][0], ",")
		for i := range category {
			cats = append(cats, categories(category[i]))
		}

	}
	floorCPM, _ := strconv.ParseInt(data["floor_cpm"][0], 10, 0)
	if floorCPM == 0 {
		floorCPM = 0
	}
	softFloor, _ := strconv.ParseInt(data["soft_floor"][0], 10, 0)
	if softFloor == 0 {
		softFloor = floorCPM
	}
	res.UnderFloor = true
	res.IP = data["ip"][0]
	if res.IP == "" {
		res.IP = "46.209.239.50"
	}
	res.Categories = cats
	res.Source = &restPublisher{
		PubFloorCPM:     floorCPM,
		PubName:         data["publisher_name"][0],
		PubSoftFloorCPM: softFloor,
	}
	res.Type = rType
	res.PageTrackID = data["page_track_id"][0]
	res.UserTrackID = data["user_track_id"][0]
	res.Scheme = "http"
	switch rType {
	case typeWeb:
		res.Web.Referrer = refferer
		res.Web.Parent = parent
		res.Web.UserAgent = userAgent
	}
	res.Slots = resSlots
	resData, err := json.MarshalIndent(res, "\t", "\t")
	assert.Nil(err)
	logrus.Debug(string(resData))
	request, err := http.NewRequest("POST", "http://exchange.clickyab.com/api/rest/get/7f423ad4ccaf4ddf84d0227d584b7f9c", bytes.NewBuffer(resData))
	if err != nil {
		return
	}
	request.Header.Add("X-REAL-IP", "46.209.239.50")
	fmt.Println(err)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		logrus.Debug(err)
		return
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return
	}
	logrus.Debug(string(responseData))
	err = json.Unmarshal(responseData, &respon)
	assert.Nil(err)
	t, err := template.New("show").Parse(show)
	assert.Nil(err)
	w.Header().Set("Content-Type", "text/html")
	logrus.Warn(t.Execute(w, respon))
}

type slotRest struct {
	W    int    `json:"width"`
	H    int    `json:"height"`
	TID  string `json:"track_id"`
	Fall string `json:"fallback_url"`
}

type restPublisher struct {
	PubName         string `json:"name"`
	PubFloorCPM     int64  `json:"floor_cpm"`
	PubSoftFloorCPM int64  `json:"soft_floor_cpm"`
}

type requestBody struct {
	IP          string         `json:"ip"`
	PageTrackID string         `json:"page_track_id"`
	UserTrackID string         `json:"user_track_id"`
	Scheme      string         `json:"scheme"`
	Source      *restPublisher `json:"publisher"`
	Categories  []categories   `json:"categories"`
	Type        string         `json:"type"`
	UnderFloor  bool
	Web         struct {
		Referrer  string `json:"referrer,omitempty"`
		Parent    string `json:"parent,omitempty"`
		UserAgent string `json:"user_agent,omitempty"`
	} `json:"web,omitempty"`
	Slots []*slotRest `json:"slots"`
}
