package renderer

import (
	cfg "services/config"

	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"octopus/exchange"
)

type dumbAd struct {
	ID      string `json:"id"`
	Winner  int64  `json:"winner"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Code    string `json:"code"`
	Landing string `json:"land"`
}

type restful struct {
	pixelPattern *url.URL
	sup          exchange.Supplier
}

func (rf restful) Render(in map[string]exchange.Advertise, w io.Writer) error {
	res := make(map[string]*dumbAd, len(in))
	for i := range in {
		if in[i] == nil {
			res[i] = nil
			continue
		}

		d := &dumbAd{
			ID:      in[i].ID(),
			Winner:  in[i].WinnerCPM() * int64(100-rf.sup.Share()) / 100,
			Width:   in[i].Width(),
			Height:  in[i].Height(),
			Landing: in[i].Landing(),
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

		host := cfg.GetStringDefault("exchange.host.name", "localhost:3412")
		trackURL := fmt.Sprintf(`%s/pixel/%s/%s`, host, in[i].Demand().Name(), in[i].TrackID())

		d.Code = fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<iframe>
    <img src="%s" alt="">
    <iframe src="%s"></iframe>
</iframe>
</body>
</html>`, trackURL, in[i].URL())
		res[i] = d
	}

	enc := json.NewEncoder(w)
	return enc.Encode(res)
}

// NewRestfulRenderer return a restful renderer
func NewRestfulRenderer(sup exchange.Supplier, pixel *url.URL) exchange.Renderer {
	return &restful{
		pixelPattern: pixel,
		sup:          sup,
	}
}
