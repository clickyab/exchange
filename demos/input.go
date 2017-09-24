package demos

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bsm/openrtb"
	"github.com/clickyab/services/assert"
)

func bidRequestParser(r *http.Request) (openrtb.BidRequest, error) {
	data, err := ioutil.ReadAll(r.Body)
	defer func() {
		err := r.Body.Close()
		assert.Nil(err)
	}()
	assert.Nil(err)

	var br openrtb.BidRequest
	err = json.Unmarshal(data, &br)
	if err != nil {
		return openrtb.BidRequest{}, err
	}

	validationError := br.Validate()
	if validationError != nil {
		return openrtb.BidRequest{}, validationError
	}
	return br, nil
}
