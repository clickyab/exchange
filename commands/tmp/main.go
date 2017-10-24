package main

import (
	"encoding/json"
	"fmt"

	"github.com/bsm/openrtb"
	"github.com/clickyab/services/assert"
)

func main() {
	x := openrtb.BidRequest{
		ID: "zxcvzxcvasdfasd3",
		Imp: []openrtb.Impression{
			openrtb.Impression{
				ID:     "czxcvzxcz",
				Secure: 1,
				Banner: &openrtb.Banner{
					ID: "zxcvasdad",
					W:  250,
					H:  350,
				},
				BidFloorCurrency: "IRR",
				BidFloor:         350,
			},
			openrtb.Impression{
				ID:     "xcvbxcvw",
				Secure: 1,
				Banner: &openrtb.Banner{
					ID: "vxcvbxcv",
					W:  250,
					H:  350,
				},
				BidFloorCurrency: "IRR",
				BidFloor:         450,
			},
		},
		AllImps: 2,
		Device: &openrtb.Device{
			IP:         "8.8.8.8",
			H:          1080,
			W:          1920,
			ConnType:   1,
			MCCMNC:     "234-543",
			Carrier:    "MCI",
			Language:   "US-en",
			OS:         "windows",
			Model:      "rt4",
			Make:       "xyz",
			DeviceType: 1,
			UA:         "user agent",
			DNT:        1,
			FlashVer:   "9.3",
		},
		Bcat:        []string{"iab-sport", "iab-food"},
		WLang:       []string{"IR-fa", "US-en"},
		BSeat:       []string{"p30download.ir", "isna.ir"},
		WSeat:       []string{"irna.ir"},
		TMax:        450,
		AuctionType: 1,
		Test:        1,
		User: &openrtb.User{
			ID:     "23452345243",
			Gender: "male",
		},
		App: &openrtb.App{
			Bundle: "string",
			Ver:    "1",
			Inventory: openrtb.Inventory{
				ID:   "abc12341234562345",
				Name: "string",
				Publisher: &openrtb.Publisher{
					Name:   "time.ir",
					ID:     "zxcvzxc-42345245245",
					Cat:    []string{},
					Domain: "time.ir",
				},
			},
		},
		Site: &openrtb.Site{
			Ref:  "string",
			Page: "time.ir/product/2342",
			Inventory: openrtb.Inventory{
				ID:   "abc12341234562345",
				Name: "string",
				Publisher: &openrtb.Publisher{
					Name:   "time.ir",
					ID:     "zxcvzxc-42345245245",
					Cat:    []string{},
					Domain: "time.ir",
				},
			},
		},
	}
	y := openrtb.BidResponse{
		ID:    "sdfgsdfgsdfg",
		BidID: "fgsdfgsdfg",
		NBR:   0,
		SeatBid: []openrtb.SeatBid{
			openrtb.SeatBid{
				Bid: []openrtb.Bid{openrtb.Bid{
					ID:       "sdfsdfsdf",
					ImpID:    "czxcvzxcz",
					W:        250,
					H:        350,
					Price:    360,
					AdID:     "wertwert",
					Protocol: 0,
					AdMarkup: `
					<!doctype html>
					<html lang="en">
					<head>
					<meta charset="UTF-8">
					             <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
					                         <meta http-equiv="X-UA-Compatible" content="ie=edge">
					             <title>Document</title>
					</head>
					<body>

					</body>
					</html>
					`,
					NURL:      "http://sdfsdf.com/sdfsdf",
					Cat:       []string{},
					AdvDomain: []string{"abca.com"},
				},
				},
			},
			openrtb.SeatBid{
				Bid: []openrtb.Bid{openrtb.Bid{
					ImpID:    "czxcvzxcz",
					ID:       "sdfsdfsdfdsf",
					W:        250,
					H:        350,
					Price:    360,
					AdID:     "345634656",
					Protocol: 1,
					AdMarkup: `
					<!doctype html>
					<html lang="en">
					<head>
					<meta charset="UTF-8">
					             <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
					                         <meta http-equiv="X-UA-Compatible" content="ie=edge">
					             <title>Document</title>
					</head>
					<body>

					</body>
					</html>
					`,
					NURL:      "http://sdfsdf.com/345345",
					Cat:       []string{},
					AdvDomain: []string{"abc.com"},
				},
				},
			},
		},
	}
	var _ = x
	m, e := json.MarshalIndent(y, "", " ")
	assert.Nil(e)
	fmt.Println(string(m))

}
