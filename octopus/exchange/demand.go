package exchange

import (
	"context"
	"database/sql/driver"
	"errors"
	"io"
	"net/http"
	"time"
)

type (
	// DemandType list all supported demand type
	DemandType string
)

const (
	// DemandTypeSrtb is for rest demand type
	DemandTypeSrtb DemandType = "srtb"
	// DemandTypeOrtb is for rest demand type
	DemandTypeOrtb DemandType = "ortb"
)

// Scan convert the json array ino string slice
func (e *DemandType) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return errors.New("unsupported type")
	}
	if !DemandType(b).IsValid() {
		return errors.New("invalid value")
	}
	*e = DemandType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e DemandType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, errors.New("invalid demand type")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e DemandType) IsValid() bool {
	// TODO : Watch it, if you add a demand type add it here too
	if e == DemandTypeOrtb || e == DemandTypeSrtb {
		return true
	}
	return false
}

// Demand demanad interface
type Demand interface {
	// Name return the name of this demand
	Name() string
	// Win return the win response to the demand. it happen only if the request is the winner
	// the 2nd arg is the id of ad, the 3rd is the winner cpm bid
	Win(context.Context, float64, string)
	// Bill call the bill url
	Bill(context.Context, float64, string)
	// Status is called for getting the statistics of this Demand
	Status(context.Context, http.ResponseWriter, *http.Request)
	// Handicap return the handicap for this demand. higher handicap means higher chance to
	// win the bid. this is the factor to multiply with cpm, 100 means 1
	Handicap() int64
	// CallRate is the rate of calling the object, 0 < X <= 100
	CallRate() int
	//WhiteListCountries is the excluded list countries for this.
	WhiteListCountries() []string
	// ExcludedSuppliers is the white listed supplier for this.
	ExcludedSuppliers() []string
	// TestMode return true if this demand is a test demand. just test mode supplier are
	// sent to this demand
	TestMode() bool
	// Type of demand (ortb, srtb)
	Type() DemandType
	// HasLimits check demand limit
	HasLimits() bool
	// EndPoint demand end-point
	EndPoint() string
	// GetTimeout get demand timeout
	GetTimeout() time.Duration
	// GeBidResponse try to get bid response from demand and make it proper
	GetBidResponse(context.Context, *http.Response, Supplier) (BidResponse, error)
	// RenderBidRequest try to render bid request due to the proper demand (rest/rtb)
	RenderBidRequest(context.Context, io.Writer, BidRequest) http.Header
	// Provide is the function to handle the request, provider should response
	// to this function and return all eligible ads
	// a very important note about providers :
	// they must return as soon as possible, waiting for result must be done
	// in separate co-routine. just create a channel, run a co-routine, and return.
	Provide(context.Context, BidRequest, chan BidResponse)
	// Currencies return acceptable currency by this demand
	Currencies() []string
}
