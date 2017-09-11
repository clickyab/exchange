package exchange

// SupplierClickMode means the click type that the supplier support
type SupplierClickMode string

const (
	// SupplierClickModeNone not supported click mode, just pass it as is
	SupplierClickModeNone SupplierClickMode = "none"
	// SupplierClickModeQueryParam url modified by adding query parameter to it
	SupplierClickModeQueryParam SupplierClickMode = "query"
	// SupplierClickModeReplace url is modified with replace command
	SupplierClickModeReplace SupplierClickMode = "replace"
	// SupplierClickModeReplaceB64 url is modified with replace command and with raw base64 // fucking adro
	SupplierClickModeReplaceB64 SupplierClickMode = "replaceb"
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
	// Renderer return the renderer of this supplier
	Renderer() Renderer
	// TestMode means this is in test mode, just pass them to test providers
	TestMode() bool
	// ClickMode return the supported click mode
	ClickMode() SupplierClickMode
	// Type return the supplier type currently only rest is supported
	Type() string
}
