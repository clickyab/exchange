package exchange

import (
	"context"
	"io"
	"net/http"
)

const (
	// SupplierORTB is type of this layer
	SupplierORTB = "ortb"
	// SupplierSRTB is type of this layer
	SupplierSRTB = "srtb"
)

// Supplier base supplier interface
type Supplier interface {
	// Name of Supplier
	Name() string
	// CPMFloor is the floor for this network. the publisher must be greeter equal to this
	FloorCPM() float64
	// SoftFloorCPM is the soft version of floor cpm. if the publisher ahs it, then the system
	// try to use this as floor, but if this is not available, the FloorCPM is used
	SoftFloorCPM() float64
	// ExcludedDemands is the Excluded list network for this.
	ExcludedDemands() []string
	// Share return the share of this supplier
	Share() int
	// TestMode means this is in test mode, just pass them to test providers
	TestMode() bool
	// Type return the supplier type currently only rest is supported
	Type() string
	// Currency returns the supplier currency type
	Currency() string
	// GetBidRequest generate bid-request from request
	GetBidRequest(context.Context, *http.Request) (BidRequest, error)
	// RenderBidResponse return the renderer of this supplier
	RenderBidResponse(context.Context, io.Writer, BidResponse) http.Header
}
