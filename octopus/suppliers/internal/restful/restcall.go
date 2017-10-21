package restful

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"clickyab.com/exchange/octopus/exchange"

	"github.com/sirupsen/logrus"
)


// GetBidRequest try to create an impression object from a request (rest)
func GetBidRequest(sup exchange.Supplier, r *http.Request) (exchange.BidRequest, error) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	bq := bidRequest{}
	err := dec.Decode(&bq)
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}

	var res *bidRequestRest
	switch strings.ToLower(string(bq.FType)) {
	case "web":
		res, err = newImpressionFromWebRequest(sup, &bq)
	default:
		err = fmt.Errorf("type is not supported: %s", bq.FType)
	}

	if err != nil {
		return nil, err
	}

	// Hidden profit is here. the floor and soft floor are rising here
	share := int64(100 + res.Pub.sup.Share())
	res.Pub.PubFloorCPM = (res.Pub.PubFloorCPM * share) / 100
	if res.Pub.PubFloorCPM == 0 {
		res.Pub.PubFloorCPM = (sup.FloorCPM() * share) / 100
	}
	res.Pub.PubSoftFloorCPM = (res.Pub.PubSoftFloorCPM * share) / 100
	if res.Pub.PubSoftFloorCPM == 0 {
		res.Pub.PubSoftFloorCPM = (sup.SoftFloorCPM() * share) / 100
	}
	return res, nil
}
