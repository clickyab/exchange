package restful

import (
	"encoding/json"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"

	"errors"

	"github.com/sirupsen/logrus"
)

// GetBidRequest try to create an impression object from a request (rest)
func GetBidRequest(sup exchange.Supplier, r *http.Request) (exchange.BidRequest, error) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	req := bidRequest{}
	err := dec.Decode(&req)
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}

	//validation
	if req.ISite != nil && req.IApp != nil {
		return nil, errors.New("both site and app can't be filled a the same time")
	}

	if req.ISite != nil {
		req.ISite.setSupplierFloors(sup)
	} else if req.IApp != nil {
		req.IApp.setSupplierFloors(sup)
	} else {
		return nil, errors.New("unsupported publisher type")
	}

	return req, nil
}
