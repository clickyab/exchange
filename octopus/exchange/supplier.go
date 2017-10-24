package exchange

import (
	"context"
	"io"
	"net/http"
)

// Supplier is the ad-network interface
type Supplier interface {
	// Name of Supplier
	Name() string
	// CPMFloor is the floor for this network. the publisher must be greeter equal to this
	FloorCPM() int64
	// SoftFloorCPM is the soft version of floor cpm. if the publisher ahs it, then the system
	// try to use this as floor, but if this is not available, the FloorCPM is used
	SoftFloorCPM() int64
	// ExcludedDemands is the Excluded list network for this.
	ExcludedDemands() []string
	// Share return the share of this supplier
	Share() int
	// RenderBidResponse return the renderer of this supplier
	RenderBidResponse(context.Context, io.Writer, BidResponse) http.Header
	// TestMode means this is in test mode, just pass them to test providers
	TestMode() bool

	// Type return the supplier type currently only rest is supported
	Type() string
	// GetBidRequest generate bid-request from request
	GetBidRequest(context.Context, *http.Request) BidRequest
}
