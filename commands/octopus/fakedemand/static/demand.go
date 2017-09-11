package static

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/random"
	"github.com/sirupsen/logrus"
)

var (
	expire     = config.RegisterDuration("exam.expire", 1*time.Hour, "")
	mountPoint = config.RegisterString("services.framework.controller.mount_point", "", "")
)

const (
	prefixSlot = "EXAM_SLOT"

	ad = "a"
)

type exchangeResponse struct {
	ID          string `json:"id"`
	MaxCPM      int64  `json:"max_cpm"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	URL         string `json:"url"`
	Landing     string `json:"landing"`
	SlotTrackID string `json:"slot_track_id"`
}

// demandHandler for handling exam (test) account
func demandHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	d := json.NewDecoder(r.Body)
	defer r.Body.Close()

	rq := &request{}
	rq.Host = r.Host
	e := d.Decode(rq)

	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(e.Error()))
	}
	data, err := json.MarshalIndent(rq, "", "\t")
	logrus.WithField("err", err).Debug(string(data))
	res, err := demandRequest(*rq)
	assert.Nil(err)
	jr, e := json.Marshal(res)
	assert.Nil(e)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jr)
}

func demandRequest(rq request) ([]exchangeResponse, error) {
	tid := <-random.ID
	res := make([]exchangeResponse, 0)
	min := rq.Publisher.PubSoftFloorCPM
	if min < rq.Publisher.PubFloorCPM {
		min = rq.Publisher.PubFloorCPM
	}
	min = inRange(min+1, min+200)
	for _, v := range rq.Slots {
		tk := slotKeyGen(tid, v.TID)
		sk := kv.NewEavStore(tk)
		res = append(res, addSlot(v, rq, min, tid))
		clu, e := prepareClickURL(sk, v, rq, min, tid)
		if e != nil {
			return nil, e
		}
		prepareShow(sk, v, rq, min, tid, clu)

		sk.Save(expire.Duration())
	}
	return res, nil

}

func addSlot(v Slot, rq request, cpm int64, tid string) exchangeResponse {

	return exchangeResponse{
		SlotTrackID: v.TID,
		Height:      v.H,
		Width:       v.W,
		ID:          <-random.ID,
		Landing:     fmt.Sprintf(`%s://%s`, rq.Scheme, rq.Host),
		MaxCPM:      cpm,
		URL:         fmt.Sprintf(`%s://%s%s/exam/ad/%s/%s/%d/%d`, rq.Scheme, rq.Host, mountPoint.String(), tid, v.TID, v.W, v.H),
	}
}

func prepareClickURL(k kv.Kiwi, v Slot, rq request, cpm int64, tid string) (string, error) {
	cu, cuOk := v.FAttribute["click_url"]
	cp, cpOk := v.FAttribute["click_parameter"]
	ct := v.FAttribute["type"]
	curl := fmt.Sprintf(`%s://%s%s/exam/click/%s/%s`, rq.Scheme, rq.Host, mountPoint.String(), tid, v.TID)
	if cuOk && cpOk {
		curl = base64.URLEncoding.WithPadding('.').EncodeToString([]byte(curl))
		if ct == "replace" {
			curl = strings.Replace(cu, cp, curl, -1)
		} else {
			tu, e := url.Parse(cu)
			if e != nil {
				return "", fmt.Errorf("not valid url")
			}
			tq := tu.Query()
			tq.Set(cp, curl)
			tu.RawQuery = tq.Encode()
			curl = tu.String()
		}
	}

	return curl, nil
}

func prepareShow(k kv.Kiwi, v Slot, rq request, cpm int64, tid, clu string) {
	tm := response{
		IsFilled: true,
		Height:   v.H,
		Width:    v.W,
		TrackID:  v.TID,
		Winner:   cpm,
	}

	tm.AdTrackID = <-random.ID

	tm.Code = string(filler(clu, "SlotID:"+v.TID, fmt.Sprintf("%d", v.W), fmt.Sprintf("%d", v.H)))

	k.SetSubKey(ad, tm.Code)
}
