package renderer

import (
	"encoding/json"
	"entity"
	"fmt"
	"io"
	"net/url"
)

type dumbAd struct {
	ID     string `json:"id"`
	Winner int64  `json:"winner"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Code   string `json:"code"`
}

type restful struct {
	pixelPattern *url.URL
	sup          entity.Supplier
}

func (rf restful) Render(in map[string]entity.Advertise, w io.Writer) error {
	res := make(map[string]*dumbAd, len(in))
	for i := range in {
		if in[i] == nil {
			res[i] = nil
			continue
		}

		d := &dumbAd{
			ID:     in[i].ID(),
			Winner: in[i].WinnerCPM() * int64(100-rf.sup.Share()) / 100,
			Width:  in[i].Width(),
			Height: in[i].Height(),
		}
		var x url.URL = *rf.pixelPattern
		q := x.Query()
		q.Set("id", i)
		x.RawQuery = q.Encode()

		winURL := in[i].URL()
		win, err := url.Parse(winURL)
		if err == nil {
			q := win.Query()
			q.Set("win", fmt.Sprint(in[i].WinnerCPM()))
			win.RawQuery = q.Encode()
			winURL = win.String()
		}

		d.Code = fmt.Sprintf("TODO , tracker code is %s, the actual route is %s", x.String(), winURL)
		res[i] = d
	}

	enc := json.NewEncoder(w)
	return enc.Encode(res)
}

// NewRestfulRenderer return a restful renderer
func NewRestfulRenderer(sup entity.Supplier, pixel *url.URL) entity.Renderer {
	return &restful{
		pixelPattern: pixel,
		sup:          sup,
	}
}
