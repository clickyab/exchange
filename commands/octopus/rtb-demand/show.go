package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/random"
	"github.com/sirupsen/logrus"
)

var template = `<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    <style >
         body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }
        .ad{
            position: relative;

            background: #fff3cd;
        }
        .meta {
            line-height: 1.8em;
            -webkit-transition: all .3s;
            -moz-transition: all .3s;
            -ms-transition: all .3s;
            -o-transition: all .3s;
            transition: all .3s;
            padding: 5px;
            color: #fff3cd;
            font-weight: bold;
            position: absolute;
            left: 0;
            top: 0;
            bottom: 0;
            right: 0;
            background-color: darkred;
            opacity: .0;
        }
        .meta:hover {
            opacity: .7;
        }
    </style>
</head>
<body>
<div class="ad">

    <script src="%s"></script>
    <a target="_blank" href="%s">
        <div class="meta">
            <div>PRICE: %s</div>
            <div>IMP: %s</div>
            <div>CUR: %s</div>
            <div></div>
        </div>
        <img src="%s" alt="just for test">
    </a>
</div>
</body>
</html>
`

func fragmentHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	imp := r.URL.Query().Get("imp")
	prc := r.URL.Query().Get("prc")
	cur := r.URL.Query().Get("cur")
	clickURL, err := clickURL(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	adURL, err := adURL(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	showURL, err := showURL(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := fmt.Sprintf(template, showURL, clickURL, prc, imp, cur, adURL)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))

}

func adURL(r *http.Request) (string, error) {
	wi := r.URL.Query().Get("wi")
	he := r.URL.Query().Get("he")
	res := url.URL{
		Host:   "via.placeholder.com",
		Scheme: r.URL.Scheme,
		Path:   fmt.Sprintf("%sx%s", wi, he),
	}
	return res.String(), nil
}

func clickURL(r *http.Request) (string, error) {
	aid := r.URL.Query().Get("aid")
	if aid == "" {
		return "", errors.New("bad request")
	}
	crl, err := base64.URLEncoding.WithPadding('.').DecodeString(r.URL.Query().Get("crl"))
	if err != nil {
		return "", err
	}
	t := url.URL{
		Host:   r.Host,
		Scheme: r.URL.Scheme,
		Path:   fmt.Sprintf("%s?cache=%s", router.MustPath("rtb-demand-click", map[string]string{"id": aid}), <-random.ID),
	}
	exchangeClick, err := url.Parse(string(crl))
	exchangeClick.Scheme = r.URL.Scheme
	logrus.Warn(exchangeClick.String())
	if err != nil {
		return "", err
	}
	tq := exchangeClick.Query()
	tq.Add("ref", base64.URLEncoding.WithPadding('.').EncodeToString([]byte(t.String())))
	exchangeClick.RawQuery = tq.Encode()
	logrus.Warn(exchangeClick.String())

	return exchangeClick.String(), nil
}

func showURL(r *http.Request) (string, error) {
	sho, err := base64.URLEncoding.WithPadding('.').DecodeString(r.URL.Query().Get("sho"))
	if err != nil {
		return "", err
	}

	res, err := url.Parse(string(sho))
	if err != nil {
		return "", err
	}
	res.Host = exchange

	return res.String(), nil
}
