package restful

import (
	"encoding/json"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"

	"errors"

	"time"

	"github.com/clickyab/services/ip2location"
	"github.com/sirupsen/logrus"
)

var (

	// ErrInvalidReqNoID request has no id
	ErrInvalidReqNoID = errors.New("rest: request ID missing")
	// ErrInvalidReqNoImps request has no impressions
	ErrInvalidReqNoImps = errors.New("rest: request has no impressions")
	// ErrInvalidReqMultiInv request has multiple impressions
	ErrInvalidReqMultiInv = errors.New("rest: request has multiple inventory sources")
	// ErrInvalidImpNoID Imp id missing
	ErrInvalidImpNoID = errors.New("rest: impression ID missing")
	// ErrInvalidImpAssets impression assets error
	ErrInvalidImpAssets = errors.New("rest: impression has multiple assets or no assets")
)

// GetBidRequest try to create an impression object from a request (rest)
func GetBidRequest(sup exchange.Supplier, r *http.Request) (exchange.BidRequest, error) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	req := BidRequest{}
	err := dec.Decode(&req)
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}
	//validation
	err = req.Validate()
	if err != nil {
		return nil, err
	}
	req.time = time.Now()
	req.IDevice.IGeo = getGeo(req.IDevice.IIP)
	if req.ISite != nil {
		req.ISite.setSupplier(sup)
	} else if req.IApp != nil {
		req.IApp.setSupplier(sup)
	} else {
		return nil, errors.New("unsupported publisher type")
	}

	return req, nil
}

func getGeo(ip string) Geo {
	t := ip2location.GetAll(ip)
	return Geo{
		IIsp: exchange.ISP{
			Valid: t.Isp != "",
			Name:  t.Isp,
		},
		IRegion: exchange.Region{
			Valid: t.Region != "",
			Name:  t.Region,
			ISO:   t.Region,
		},
		ICountry: exchange.Country{
			Valid: t.CountryLong != "",
			Name:  t.CountryLong,
			ISO:   t.CountryShort,
		},
		ILatLon: exchange.LatLon{
			Valid: true,
			Lat:   float64(t.Latitude),
			Lon:   float64(t.Longitude),
		},
	}
}

// validate decoded bid request object
func (bq BidRequest) Validate() error {
	if bq.IID == "" {
		return ErrInvalidReqNoID
	} else if len(bq.Imp()) == 0 {
		return ErrInvalidReqNoImps
	} else if bq.ISite != nil && bq.IApp != nil {
		return ErrInvalidReqMultiInv
	}
	for i := range bq.Imp() {
		if bq.Imp()[i].ID() == "" {
			return ErrInvalidImpNoID
		}
		assetCount := func(x exchange.Impression) int {
			var n = 0
			if x.Banner() != nil {
				n++
			}
			if x.Video() != nil {
				n++
			}
			if x.Native() != nil {
				n++
			}
			return n

		}(bq.Imp()[i])
		if assetCount != 1 {
			return ErrInvalidImpAssets
		}
	}
	return nil
}
