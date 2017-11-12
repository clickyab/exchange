package biding

import (
	"encoding/base64"
	"fmt"
	"strings"

	"context"

	"crypto/sha1"
	"net/url"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/xlog"
)

const (
	clickURLHolderB = "${CLICK_URL:B64}"
	jsHolderB       = "${PIXEL_URL_JS:B64}"
	pixelHolderB    = "${PIXEL_URL_IMAGE:B64}"
	idHolderB       = "${AUCTION_ID:B64}"
	bidIDHolderB    = "${AUCTION_BID_ID:B64}"
	impIDHolderB    = "${AUCTION_IMP_ID:B64}"
	seatIDHolderB   = "${AUCTION_SEAT_ID:B64}"
	adIDHolderB     = "${AUCTION_AD_ID:B64}"
	priceHolderB    = "${AUCTION_PRICE:B64}"
	currencyHolderB = "${AUCTION_CURRENCY:B64}"
	mbrHolderB      = "${AUCTION_MBR:B64}"
	lossHolderB     = "${AUCTION_LOSS:B64}"
	clickURLHolder  = "${CLICK_URL}"
	jsHolder        = "${PIXEL_URL_JS}"
	pixelHolder     = "${PIXEL_URL_IMAGE}"
	idHolder        = "${AUCTION_ID}"
	bidIDHolder     = "${AUCTION_BID_ID}"
	impIDHolder     = "${AUCTION_IMP_ID}"
	seatIDHolder    = "${AUCTION_SEAT_ID}"
	adIDHolder      = "${AUCTION_AD_ID}"
	priceHolder     = "${AUCTION_PRICE}"
	currencyHolder  = "${AUCTION_CURRENCY}"
	mbrHolder       = "${AUCTION_MBR}"
	lossHolder      = "${AUCTION_LOSS}"
)

func hasTracker(b exchange.Bid) bool {
	if strings.Contains(b.AdMarkup(), jsHolder) ||
		strings.Contains(b.AdMarkup(), jsHolderB) ||
		strings.Contains(b.AdMarkup(), pixelHolder) ||
		strings.Contains(b.AdMarkup(), pixelHolderB) {
		return true
	}
	return false
}

func replacer(ctx context.Context, q exchange.BidRequest, b exchange.Bid) *strings.Replacer {
	scheme := func() string {
		for _, m := range q.Imp() {
			if b.ImpID() == m.ID() {
				if m.Secure() {
					return "https"
				}
				return "http"
			}
		}
		xlog.Get(ctx).Panic("[BUG]")
		return ""
	}()
	key := GenRedisKey(ctx, q, b)
	pat, err := router.Path("pixel", map[string]string{"id": key, "type": "show.js"})
	assert.Nil(err)
	show := url.URL{
		Scheme: scheme,
		Host:   q.Request().Host,
		Path:   pat,
	}
	js := show.String()
	show.Path, err = router.Path("pixel", map[string]string{"id": key, "type": "image.png"})
	assert.Nil(err)

	pixel := show.String()
	b64 := base64.URLEncoding.WithPadding('.')
	clk, err := router.Path("click", map[string]string{"id": key})
	assert.Nil(err)
	win := url.URL{
		Scheme: scheme,
		Host:   q.Request().Host,
		Path:   clk,
	}
	return strings.NewReplacer([]string{
		clickURLHolder, win.String(),
		clickURLHolderB, b64.EncodeToString([]byte(win.String())),
		jsHolder, js,
		jsHolderB, b64.EncodeToString([]byte(js)),
		pixelHolder, pixel,
		pixelHolderB, b64.EncodeToString([]byte(pixel)),
		idHolder, q.ID(),
		idHolderB, b64.EncodeToString([]byte(q.ID())),
		bidIDHolder, b.ID(),
		bidIDHolderB, b64.EncodeToString([]byte(b.ID())),
		impIDHolder, b.ImpID(),
		impIDHolderB, b64.EncodeToString([]byte(b.ImpID())),
		seatIDHolder, b.ImpID(),
		seatIDHolderB, b64.EncodeToString([]byte(b.ImpID())),
		adIDHolder, b.AdID(),
		adIDHolderB, b64.EncodeToString([]byte(b.AdID())),
		priceHolder, fmt.Sprintf("%d", b.Price()),
		priceHolderB, b64.EncodeToString([]byte(fmt.Sprintf("%d", b.Price()))),
		currencyHolder, "IRR",
		currencyHolderB, b64.EncodeToString([]byte("IRR")),
		mbrHolder, "1",
		mbrHolderB, b64.EncodeToString([]byte("1")),
		lossHolder, "“AUDIT",
		lossHolderB, b64.EncodeToString([]byte("AUDIT")),
	}...)

}

// GenRedisKey sets redis key for unique bid and bid request
func GenRedisKey(ctx context.Context, br exchange.BidRequest, bid exchange.Bid) string {
	wholeData := fmt.Sprintf("%v%s", br, bid.AdID())
	hash := sha1.New()
	_, err := hash.Write([]byte(wholeData))
	assert.Nil(err)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
