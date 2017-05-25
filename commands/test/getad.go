package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"clickyab.com/exchange/services/random"
)

type restAd struct {
	RID     string `json:"id"`
	RMaxCPM int64  `json:"max_cpm"`
	RWidth  int    `json:"width"`
	RHeight int    `json:"height"`
	RURL    string `json:"url"`
}

func getAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	mp := make(map[string]interface{})
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := dec.Decode(&mp)
	if err != nil {
		return
	}
	t, _ := json.MarshalIndent(mp, "\t", "\t")
	fmt.Println(string(t))

	x := map[string]restAd{}

	x[mp["slots"].([]interface{})[0].(map[string]interface{})["track_id"].(string)] = restAd{
		RID:     <-random.ID,
		RMaxCPM: 150,
		RWidth:  320,
		RHeight: 250,
		RURL:    "http://google.com",
	}
	x[mp["slots"].([]interface{})[1].(map[string]interface{})["track_id"].(string)] = restAd{
		RID:     <-random.ID,
		RMaxCPM: 150,
		RWidth:  320,
		RHeight: 250,
		RURL:    "http://yahoo.com",
	}

	enc := json.NewEncoder(w)
	enc.Encode(x)
}