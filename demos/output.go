package demos

import (
	"bytes"
	"html/template"
	"strconv"

	"fmt"

	"crypto/sha1"
	"time"

	"github.com/bsm/openrtb"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"
)

type ad struct {
	Link   string
	Width  string
	Height string
	Src    string
	Tiny   string
}

func bidFakeResponse(request openrtb.BidRequest) (openrtb.BidResponse) {
	var noticeURL string = "asd.com"
	buffer := &bytes.Buffer{}
	temp := template.Must(template.New("single_ad").Parse(singleAd))

	var sb = make([]openrtb.SeatBid, len(request.Imp))
	for i, imp := range request.Imp {
		clickKey := sha1.New()
		clickKey.Write([]byte(fmt.Sprintf("%v", imp)))
		clickURL := fmt.Sprintf("/click/%s", clickKey.Sum(nil))

		err := temp.Execute(buffer, ad{
			Width:  strconv.Itoa(imp.Banner.W),
			Height: strconv.Itoa(imp.Banner.H),
			Link:   clickURL,
			Src:    fmt.Sprintf("http://a.clickyab.com/ads/?a=4471405272967&width=%d&height=%d&slot=71634138754&domainname=p30download.com&eventpage=416310534&loc=http://p30download.com/agahi/plan/a1i.php&ref=http://p30download.com/&adcount=1", imp.Banner.W, imp.Banner.H),
		})
		assert.Nil(err)

		sb[i] = openrtb.SeatBid{
			Bid: []openrtb.Bid{{
				ID:       strconv.Itoa(i),
				ImpID:    imp.ID,
				Price:    imp.BidFloor,
				NURL:     noticeURL,
				AdMarkup: buffer.String(),
				W:        imp.Banner.W,
				H:        imp.Banner.H,
			}},
		}

		store := kv.NewEavStore(clickURL)
		assert.Nil(store.Save(time.Hour))
	}

	return openrtb.BidResponse{
		ID:       request.ID,
		SeatBid:  sb,
		Currency: "rial",
	}
}

var singleAd = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta name="robots" content="nofollow">
    <meta content="text/html; charset=utf-8" http-equiv="Content-Type">
    <style>
html, body, div, span, applet, object, iframe,
h1, h2, h3, h4, h5, h6, p, blockquote, pre,
a, abbr, acronym, address, big, cite, code,
del, dfn, em, font, img, ins, kbd, q, s, samp,
small, strike, strong, sub, sup, tt, var,
dl, dt, dd, ol, ul, li,
fieldset, form, label, legend,
table, caption, tbody, tfoot, thead, tr, th, td {
  margin: 0;
  padding: 0;
  border: 0;
  outline: 0;
  font-weight: inherit;
  font-style: inherit;
  font-size: 100%;
  font-family: inherit;
  vertical-align: baseline;
}
/* remember to define focus styles! */
:focus {
  outline: 0;
}
body {
  line-height: 1;
  color: black;
  background: white;
}
ol, ul {
  list-style: none;
}
/* tables still need 'cellspacing="0"' in the markup */
table {
  border-collapse: separate;
  border-spacing: 0;
}
caption, th, td {
  text-align: left;
  font-weight: normal;
}
blockquote:before, blockquote:after,
q:before, q:after {
  content: "";
}
blockquote, q {
  quotes: "" "";
}
        body{ margin: 0; padding: 0; text-align: center; }
        .o{ position:absolute; top:0; left:0; border:0; height:250px; width:300px; z-index: 99; }
        #showb{ position:absolute; top:0; left:0; border:0; line-height: 250px; height:250px; width:300px; z-index: 100; background: rgba(0, 0, 0, 0.60); text-align: center; }
        {{ if .Tiny }}
        .tiny2{ height: 18px; width: 19px; position: absolute; bottom: 0px; right: 0; z-index: 100; background: url("//static.clickyab.com/img/clickyab-tiny.png") right top no-repeat; border-top-left-radius:4px; -moz-border-radius-topleft:4px  }
        .tiny2:hover{ width: 66px;  }
        .tiny{ height: 18px; width: 19px; position: absolute; top: 0px; right: 0; z-index: 100; background: url("//static.clickyab.com/img/clickyab-tiny.png") right top no-repeat; border-bottom-left-radius:4px; -moz-border-radius-bottomleft:4px  }
        .tiny:hover{ width: 66px;  }
        .tiny3{ position: absolute; top: 0px; right: 0; z-index: 100; }
        {{ end }}
        .butl {background: #4474CB;color: #FFF;padding: 10px;text-decoration: none;border: 2px solid #FFFFFF;font-family: tahoma;font-size: 13px;}
        img.adhere {max-width:100%;height:auto;}
        video {background: #232323 none repeat scroll 0 0;}
		</style>
</head>
<body>

    {{ if .Tiny }}<a class="tiny" href="http://clickyab.com/?ref=icon" target="_blank"></a>{{ end }}
	<a href="{{ .Link }}" target="_blank"><img  src="{{ .Src }}" border="0" height="{{ .Height }}" width="{{ .Width }}"/></a>
<br style="clear: both;"/>
</body></html>`
