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
)

const (
	warningMsg        = `this field is temporary and will not appear on real request\n`
	setSlotTrackIDMsg = `Slot TrackID has been set by system. possible reason
	 1. it was empty.
	 2. it is not unique on in this request.\n`
)

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

var chaosMonkey = config.RegisterBoolean("fake_demand.chaos", false, "")

func demandRequest(rq request) ([]response, error) {
	tid := <-random.ID

	res := make([]response, 0)
	code := make(map[string]string)
	min := rq.Publisher.PubSoftFloorCPM
	if min == 0 {
		min = rq.Publisher.PubFloorCPM
	}
	stq := make([]string, 0)
	for _, v := range rq.Slots {

		tm := response{}
		if _, ok := code[v.TID]; ok || v.TID == "" {
			v.TID = <-random.ID
			tm.Description += warningMsg + setSlotTrackIDMsg
		}
		tm.Height = v.H
		tm.Width = v.W

		if chaosMonkey.Bool() {
			// Just to send some slots empty
			if inRange(1, 10)%5 == 0 {
				tm.Code = v.FallbackURL
				tm.IsFilled = false
				res = append(res, tm)
				continue
			}
		}

		stq = append(stq, v.TID)
		tm.Winner = inRange(int(min)+1, int(min)+500)
		tm.AdTrackID = <-random.ID
		tm.Code = render(codeModel{
			Width:  v.W,
			Height: v.H,
			Show:   fmt.Sprintf(`%s://%s%s/exam/show/%s/%s`, rq.Scheme, rq.Host, mountPoint.String(), tid, v.TID),
			Pixel:  fmt.Sprintf(`%s://%s%s/exam/pixel/%s/%s`, rq.Scheme, rq.Host, mountPoint.String(), tid, v.TID),
		})

		tm.IsFilled = true
		tm.Landing = rq.Host

		ks := kv.NewEavStore(slotKeyGen(tid, v.TID))

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
					return nil, fmt.Errorf("not valid url")
				}
				tq := tu.Query()
				tq.Set(cp, curl)
				tu.RawQuery = tq.Encode()
				curl = tu.String()
			}

		}
		res = append(res, tm)
		ks.SetSubKey(clickURL, curl)
		ks.Save(expire.Duration())
	}
	r, e := json.MarshalIndent(res, "", "  ")
	assert.Nil(e)
	k := kv.NewEavStore(fmt.Sprintf(`%s_%s`, prefixImpression, tid))
	k.SetSubKey(slots, strings.Join(stq, ","))
	o, _ := json.MarshalIndent(rq, "", "  ")

	k.SetSubKey(org, string(o))
	k.SetSubKey(raw, string(r))
	k.Save(expire.Duration())

	return res, nil

}
