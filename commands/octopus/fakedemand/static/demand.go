package static

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/random"
)

const webTempl = `<!DOCTYPE html>
<html lang="en"><head><meta charset="UTF-8"><title>clickyab.com</title>
<style>#adhere iframe {max-width:100%;margin: 0 auto;}
    .show {position: absolute; top : -1000px; left : -1000px}
    </style>
</head>
<body style="margin: 0; padding: 0; display:flex; width:100%;">
    <img class="show" src="{{.Pixel}}" alt="">
    <div id="adhere">
<iframe id="thirdad_frame" width="{{.Width}}" height="{{.Height}}" frameborder=0 marginwidth="0" marginheight="0" vspace="0" hspace="0" allowtransparency="true" scrolling="no"
 src="{{.Show}}" class="thirdad thrdadok">
 </iframe>
 </div></body>
 </html>`

func render(s codeModel) string {
	t, e := template.New("webTempl").Parse(webTempl)
	assert.Nil(e)
	buf := &bytes.Buffer{}

	e = t.Execute(buf, s)
	assert.Nil(e)
	return buf.String()
}

var (
	expire     = config.RegisterDuration("exam.expire", 1*time.Hour, "")
	mountPoint = config.RegisterString("services.framework.controller.mount_point", "", "")
)

const (
	prefixImpression = "EXAM_IMP"
	prefixSlot       = "EXAM_SLOT"

	slots    = "s"
	clickURL = "u"
	raw      = "r"
	org      = "o"
	ad       = "a"
)

type Result struct {
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
	res, err := demandRequest(*rq)
	assert.Nil(err)
	jr, e := json.Marshal(res)
	assert.Nil(e)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jr)
}

func demandRequest(rq request) ([]Result, error) {
	tid := <-random.ID
	res := make([]Result, 0)
	//code := make(map[string]string)
	min := rq.Publisher.PubSoftFloorCPM
	if min == 0 {
		min = rq.Publisher.PubFloorCPM
	}
	for _, v := range rq.Slots {
		tk := slotKeyGen(tid, v.TID)
		sk := kv.NewEavStore(tk)
		res = append(res, addSlot(v, rq, min, tid))
		prepareShow(sk, v, rq, min, tid)
		prepareClickURL(sk, v, rq, min, tid)
		sk.Save(expire.Duration())
	}
	return res, nil

}

func addSlot(v Slot, rq request, cpm int64, tid string) Result {

	return Result{
		SlotTrackID: v.TID,
		Height:      v.H,
		Width:       v.W,
		ID:          <-random.ID,
		Landing:     fmt.Sprintf(`%s://%s`, rq.Scheme, rq.Host),
		MaxCPM:      cpm,
		URL:         fmt.Sprintf(`%s://%s%s/exam/ad/%s/%s`, rq.Scheme, rq.Host, mountPoint.String(), tid, v.TID),
	}
}

func prepareClickURL(k kv.Kiwi, v Slot, rq request, cpm int64, tid string) error {
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
				return fmt.Errorf("not valid url")
			}
			tq := tu.Query()
			tq.Set(cp, curl)
			tu.RawQuery = tq.Encode()
			curl = tu.String()
		}
	}
	k.SetSubKey(clickURL, curl)
	return nil
}

func prepareShow(k kv.Kiwi, v Slot, rq request, cpm int64, tid string) {
	tm := response{
		IsFilled: true,
		Height:   v.H,
		Width:    v.W,
		TrackID:  v.TID,
		Winner:   cpm,
	}

	tm.AdTrackID = <-random.ID

	tm.Code = render(codeModel{

		Width:  v.W,
		Height: v.H,
		Show:   fmt.Sprintf(`%s://%s%s/exam/show/%s/%s`, rq.Scheme, rq.Host, mountPoint.String(), tid, v.TID),
		Pixel:  fmt.Sprintf(`%s://%s%s/exam/pixel/%s/%s`, rq.Scheme, rq.Host, mountPoint.String(), tid, v.TID),
	})
	tmj, e := json.Marshal(tm)
	assert.Nil(e)

	k.SetSubKey(ad, string(tmj))
}
