// Code generated by MockGen. DO NOT EDIT.
// Source: clickyab.com/exchange/octopus/exchange (interfaces: Impression,Demand,Advertise,Publisher,Location,Slot,Supplier)

package mock_exchange

import (
	exchange "clickyab.com/exchange/octopus/exchange"
	context "context"
	gomock "github.com/golang/mock/gomock"
	net "net"
	http "net/http"
	time "time"
)

// MockImpression is a mock of Impression interface
type MockImpression struct {
	ctrl     *gomock.Controller
	recorder *MockImpressionMockRecorder
}

// MockImpressionMockRecorder is the mock recorder for MockImpression
type MockImpressionMockRecorder struct {
	mock *MockImpression
}

// NewMockImpression creates a new mock instance
func NewMockImpression(ctrl *gomock.Controller) *MockImpression {
	mock := &MockImpression{ctrl: ctrl}
	mock.recorder = &MockImpressionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockImpression) EXPECT() *MockImpressionMockRecorder {
	return _m.recorder
}

// Attributes mocks base method
func (_m *MockImpression) Attributes() map[string]interface{} {
	ret := _m.ctrl.Call(_m, "Attributes")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Attributes indicates an expected call of Attributes
func (_mr *MockImpressionMockRecorder) Attributes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Attributes")
}

// Category mocks base method
func (_m *MockImpression) Category() []exchange.Category {
	ret := _m.ctrl.Call(_m, "Category")
	ret0, _ := ret[0].([]exchange.Category)
	return ret0
}

// Category indicates an expected call of Category
func (_mr *MockImpressionMockRecorder) Category() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Category")
}

// IP mocks base method
func (_m *MockImpression) IP() net.IP {
	ret := _m.ctrl.Call(_m, "IP")
	ret0, _ := ret[0].(net.IP)
	return ret0
}

// IP indicates an expected call of IP
func (_mr *MockImpressionMockRecorder) IP() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IP")
}

// Location mocks base method
func (_m *MockImpression) Location() exchange.Location {
	ret := _m.ctrl.Call(_m, "Location")
	ret0, _ := ret[0].(exchange.Location)
	return ret0
}

// Location indicates an expected call of Location
func (_mr *MockImpressionMockRecorder) Location() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Location")
}

// PageTrackID mocks base method
func (_m *MockImpression) PageTrackID() string {
	ret := _m.ctrl.Call(_m, "PageTrackID")
	ret0, _ := ret[0].(string)
	return ret0
}

// PageTrackID indicates an expected call of PageTrackID
func (_mr *MockImpressionMockRecorder) PageTrackID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PageTrackID")
}

// Platform mocks base method
func (_m *MockImpression) Platform() exchange.ImpressionPlatform {
	ret := _m.ctrl.Call(_m, "Platform")
	ret0, _ := ret[0].(exchange.ImpressionPlatform)
	return ret0
}

// Platform indicates an expected call of Platform
func (_mr *MockImpressionMockRecorder) Platform() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Platform")
}

// Scheme mocks base method
func (_m *MockImpression) Scheme() string {
	ret := _m.ctrl.Call(_m, "Scheme")
	ret0, _ := ret[0].(string)
	return ret0
}

// Scheme indicates an expected call of Scheme
func (_mr *MockImpressionMockRecorder) Scheme() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Scheme")
}

// Slots mocks base method
func (_m *MockImpression) Slots() []exchange.Slot {
	ret := _m.ctrl.Call(_m, "Slots")
	ret0, _ := ret[0].([]exchange.Slot)
	return ret0
}

// Slots indicates an expected call of Slots
func (_mr *MockImpressionMockRecorder) Slots() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Slots")
}

// Source mocks base method
func (_m *MockImpression) Source() exchange.Publisher {
	ret := _m.ctrl.Call(_m, "Source")
	ret0, _ := ret[0].(exchange.Publisher)
	return ret0
}

// Source indicates an expected call of Source
func (_mr *MockImpressionMockRecorder) Source() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Source")
}

// Time mocks base method
func (_m *MockImpression) Time() time.Time {
	ret := _m.ctrl.Call(_m, "Time")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Time indicates an expected call of Time
func (_mr *MockImpressionMockRecorder) Time() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Time")
}

// TrackID mocks base method
func (_m *MockImpression) TrackID() string {
	ret := _m.ctrl.Call(_m, "TrackID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TrackID indicates an expected call of TrackID
func (_mr *MockImpressionMockRecorder) TrackID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TrackID")
}

// UnderFloor mocks base method
func (_m *MockImpression) UnderFloor() bool {
	ret := _m.ctrl.Call(_m, "UnderFloor")
	ret0, _ := ret[0].(bool)
	return ret0
}

// UnderFloor indicates an expected call of UnderFloor
func (_mr *MockImpressionMockRecorder) UnderFloor() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnderFloor")
}

// UserAgent mocks base method
func (_m *MockImpression) UserAgent() string {
	ret := _m.ctrl.Call(_m, "UserAgent")
	ret0, _ := ret[0].(string)
	return ret0
}

// UserAgent indicates an expected call of UserAgent
func (_mr *MockImpressionMockRecorder) UserAgent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UserAgent")
}

// UserTrackID mocks base method
func (_m *MockImpression) UserTrackID() string {
	ret := _m.ctrl.Call(_m, "UserTrackID")
	ret0, _ := ret[0].(string)
	return ret0
}

// UserTrackID indicates an expected call of UserTrackID
func (_mr *MockImpressionMockRecorder) UserTrackID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UserTrackID")
}

// MockDemand is a mock of Demand interface
type MockDemand struct {
	ctrl     *gomock.Controller
	recorder *MockDemandMockRecorder
}

// MockDemandMockRecorder is the mock recorder for MockDemand
type MockDemandMockRecorder struct {
	mock *MockDemand
}

// NewMockDemand creates a new mock instance
func NewMockDemand(ctrl *gomock.Controller) *MockDemand {
	mock := &MockDemand{ctrl: ctrl}
	mock.recorder = &MockDemandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDemand) EXPECT() *MockDemandMockRecorder {
	return _m.recorder
}

// CallRate mocks base method
func (_m *MockDemand) CallRate() int {
	ret := _m.ctrl.Call(_m, "CallRate")
	ret0, _ := ret[0].(int)
	return ret0
}

// CallRate indicates an expected call of CallRate
func (_mr *MockDemandMockRecorder) CallRate() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CallRate")
}

// ExcludedSuppliers mocks base method
func (_m *MockDemand) ExcludedSuppliers() []string {
	ret := _m.ctrl.Call(_m, "ExcludedSuppliers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ExcludedSuppliers indicates an expected call of ExcludedSuppliers
func (_mr *MockDemandMockRecorder) ExcludedSuppliers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExcludedSuppliers")
}

// Handicap mocks base method
func (_m *MockDemand) Handicap() int64 {
	ret := _m.ctrl.Call(_m, "Handicap")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Handicap indicates an expected call of Handicap
func (_mr *MockDemandMockRecorder) Handicap() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Handicap")
}

// Name mocks base method
func (_m *MockDemand) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (_mr *MockDemandMockRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

// Provide mocks base method
func (_m *MockDemand) Provide(_param0 context.Context, _param1 exchange.Impression, _param2 chan exchange.Advertise) {
	_m.ctrl.Call(_m, "Provide", _param0, _param1, _param2)
}

// Provide indicates an expected call of Provide
func (_mr *MockDemandMockRecorder) Provide(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Provide", arg0, arg1, arg2)
}

// Status mocks base method
func (_m *MockDemand) Status(_param0 context.Context, _param1 http.ResponseWriter, _param2 *http.Request) {
	_m.ctrl.Call(_m, "Status", _param0, _param1, _param2)
}

// Status indicates an expected call of Status
func (_mr *MockDemandMockRecorder) Status(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Status", arg0, arg1, arg2)
}

// TestMode mocks base method
func (_m *MockDemand) TestMode() bool {
	ret := _m.ctrl.Call(_m, "TestMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TestMode indicates an expected call of TestMode
func (_mr *MockDemandMockRecorder) TestMode() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TestMode")
}

// WhiteListCountries mocks base method
func (_m *MockDemand) WhiteListCountries() []string {
	ret := _m.ctrl.Call(_m, "WhiteListCountries")
	ret0, _ := ret[0].([]string)
	return ret0
}

// WhiteListCountries indicates an expected call of WhiteListCountries
func (_mr *MockDemandMockRecorder) WhiteListCountries() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WhiteListCountries")
}

// Win mocks base method
func (_m *MockDemand) Win(_param0 context.Context, _param1 string, _param2 int64) {
	_m.ctrl.Call(_m, "Win", _param0, _param1, _param2)
}

// Win indicates an expected call of Win
func (_mr *MockDemandMockRecorder) Win(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Win", arg0, arg1, arg2)
}

// MockAdvertise is a mock of Advertise interface
type MockAdvertise struct {
	ctrl     *gomock.Controller
	recorder *MockAdvertiseMockRecorder
}

// MockAdvertiseMockRecorder is the mock recorder for MockAdvertise
type MockAdvertiseMockRecorder struct {
	mock *MockAdvertise
}

// NewMockAdvertise creates a new mock instance
func NewMockAdvertise(ctrl *gomock.Controller) *MockAdvertise {
	mock := &MockAdvertise{ctrl: ctrl}
	mock.recorder = &MockAdvertiseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockAdvertise) EXPECT() *MockAdvertiseMockRecorder {
	return _m.recorder
}

// Demand mocks base method
func (_m *MockAdvertise) Demand() exchange.Demand {
	ret := _m.ctrl.Call(_m, "Demand")
	ret0, _ := ret[0].(exchange.Demand)
	return ret0
}

// Demand indicates an expected call of Demand
func (_mr *MockAdvertiseMockRecorder) Demand() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Demand")
}

// Height mocks base method
func (_m *MockAdvertise) Height() int {
	ret := _m.ctrl.Call(_m, "Height")
	ret0, _ := ret[0].(int)
	return ret0
}

// Height indicates an expected call of Height
func (_mr *MockAdvertiseMockRecorder) Height() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Height")
}

// ID mocks base method
func (_m *MockAdvertise) ID() string {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (_mr *MockAdvertiseMockRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

// Landing mocks base method
func (_m *MockAdvertise) Landing() string {
	ret := _m.ctrl.Call(_m, "Landing")
	ret0, _ := ret[0].(string)
	return ret0
}

// Landing indicates an expected call of Landing
func (_mr *MockAdvertiseMockRecorder) Landing() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Landing")
}

// MaxCPM mocks base method
func (_m *MockAdvertise) MaxCPM() int64 {
	ret := _m.ctrl.Call(_m, "MaxCPM")
	ret0, _ := ret[0].(int64)
	return ret0
}

// MaxCPM indicates an expected call of MaxCPM
func (_mr *MockAdvertiseMockRecorder) MaxCPM() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MaxCPM")
}

// Rates mocks base method
func (_m *MockAdvertise) Rates() []exchange.Rate {
	ret := _m.ctrl.Call(_m, "Rates")
	ret0, _ := ret[0].([]exchange.Rate)
	return ret0
}

// Rates indicates an expected call of Rates
func (_mr *MockAdvertiseMockRecorder) Rates() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rates")
}

// SetWinnerCPM mocks base method
func (_m *MockAdvertise) SetWinnerCPM(_param0 int64) {
	_m.ctrl.Call(_m, "SetWinnerCPM", _param0)
}

// SetWinnerCPM indicates an expected call of SetWinnerCPM
func (_mr *MockAdvertiseMockRecorder) SetWinnerCPM(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetWinnerCPM", arg0)
}

// SlotTrackID mocks base method
func (_m *MockAdvertise) SlotTrackID() string {
	ret := _m.ctrl.Call(_m, "SlotTrackID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SlotTrackID indicates an expected call of SlotTrackID
func (_mr *MockAdvertiseMockRecorder) SlotTrackID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SlotTrackID")
}

// TrackID mocks base method
func (_m *MockAdvertise) TrackID() string {
	ret := _m.ctrl.Call(_m, "TrackID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TrackID indicates an expected call of TrackID
func (_mr *MockAdvertiseMockRecorder) TrackID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TrackID")
}

// URL mocks base method
func (_m *MockAdvertise) URL() string {
	ret := _m.ctrl.Call(_m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL
func (_mr *MockAdvertiseMockRecorder) URL() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "URL")
}

// Width mocks base method
func (_m *MockAdvertise) Width() int {
	ret := _m.ctrl.Call(_m, "Width")
	ret0, _ := ret[0].(int)
	return ret0
}

// Width indicates an expected call of Width
func (_mr *MockAdvertiseMockRecorder) Width() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Width")
}

// WinnerCPM mocks base method
func (_m *MockAdvertise) WinnerCPM() int64 {
	ret := _m.ctrl.Call(_m, "WinnerCPM")
	ret0, _ := ret[0].(int64)
	return ret0
}

// WinnerCPM indicates an expected call of WinnerCPM
func (_mr *MockAdvertiseMockRecorder) WinnerCPM() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WinnerCPM")
}

// MockPublisher is a mock of Publisher interface
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return _m.recorder
}

// Attributes mocks base method
func (_m *MockPublisher) Attributes() map[string]interface{} {
	ret := _m.ctrl.Call(_m, "Attributes")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Attributes indicates an expected call of Attributes
func (_mr *MockPublisherMockRecorder) Attributes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Attributes")
}

// FloorCPM mocks base method
func (_m *MockPublisher) FloorCPM() int64 {
	ret := _m.ctrl.Call(_m, "FloorCPM")
	ret0, _ := ret[0].(int64)
	return ret0
}

// FloorCPM indicates an expected call of FloorCPM
func (_mr *MockPublisherMockRecorder) FloorCPM() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FloorCPM")
}

// Name mocks base method
func (_m *MockPublisher) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (_mr *MockPublisherMockRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

// Rates mocks base method
func (_m *MockPublisher) Rates() []exchange.Rate {
	ret := _m.ctrl.Call(_m, "Rates")
	ret0, _ := ret[0].([]exchange.Rate)
	return ret0
}

// Rates indicates an expected call of Rates
func (_mr *MockPublisherMockRecorder) Rates() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rates")
}

// SoftFloorCPM mocks base method
func (_m *MockPublisher) SoftFloorCPM() int64 {
	ret := _m.ctrl.Call(_m, "SoftFloorCPM")
	ret0, _ := ret[0].(int64)
	return ret0
}

// SoftFloorCPM indicates an expected call of SoftFloorCPM
func (_mr *MockPublisherMockRecorder) SoftFloorCPM() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SoftFloorCPM")
}

// Supplier mocks base method
func (_m *MockPublisher) Supplier() exchange.Supplier {
	ret := _m.ctrl.Call(_m, "Supplier")
	ret0, _ := ret[0].(exchange.Supplier)
	return ret0
}

// Supplier indicates an expected call of Supplier
func (_mr *MockPublisherMockRecorder) Supplier() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Supplier")
}

// MockLocation is a mock of Location interface
type MockLocation struct {
	ctrl     *gomock.Controller
	recorder *MockLocationMockRecorder
}

// MockLocationMockRecorder is the mock recorder for MockLocation
type MockLocationMockRecorder struct {
	mock *MockLocation
}

// NewMockLocation creates a new mock instance
func NewMockLocation(ctrl *gomock.Controller) *MockLocation {
	mock := &MockLocation{ctrl: ctrl}
	mock.recorder = &MockLocationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockLocation) EXPECT() *MockLocationMockRecorder {
	return _m.recorder
}

// Country mocks base method
func (_m *MockLocation) Country() exchange.Country {
	ret := _m.ctrl.Call(_m, "Country")
	ret0, _ := ret[0].(exchange.Country)
	return ret0
}

// Country indicates an expected call of Country
func (_mr *MockLocationMockRecorder) Country() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Country")
}

// LatLon mocks base method
func (_m *MockLocation) LatLon() exchange.LatLon {
	ret := _m.ctrl.Call(_m, "LatLon")
	ret0, _ := ret[0].(exchange.LatLon)
	return ret0
}

// LatLon indicates an expected call of LatLon
func (_mr *MockLocationMockRecorder) LatLon() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LatLon")
}

// Province mocks base method
func (_m *MockLocation) Province() exchange.Province {
	ret := _m.ctrl.Call(_m, "Province")
	ret0, _ := ret[0].(exchange.Province)
	return ret0
}

// Province indicates an expected call of Province
func (_mr *MockLocationMockRecorder) Province() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Province")
}

// MockSlot is a mock of Slot interface
type MockSlot struct {
	ctrl     *gomock.Controller
	recorder *MockSlotMockRecorder
}

// MockSlotMockRecorder is the mock recorder for MockSlot
type MockSlotMockRecorder struct {
	mock *MockSlot
}

// NewMockSlot creates a new mock instance
func NewMockSlot(ctrl *gomock.Controller) *MockSlot {
	mock := &MockSlot{ctrl: ctrl}
	mock.recorder = &MockSlotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockSlot) EXPECT() *MockSlotMockRecorder {
	return _m.recorder
}

// Attributes mocks base method
func (_m *MockSlot) Attributes() map[string]string {
	ret := _m.ctrl.Call(_m, "Attributes")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Attributes indicates an expected call of Attributes
func (_mr *MockSlotMockRecorder) Attributes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Attributes")
}

// Fallback mocks base method
func (_m *MockSlot) Fallback() string {
	ret := _m.ctrl.Call(_m, "Fallback")
	ret0, _ := ret[0].(string)
	return ret0
}

// Fallback indicates an expected call of Fallback
func (_mr *MockSlotMockRecorder) Fallback() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fallback")
}

// Height mocks base method
func (_m *MockSlot) Height() int {
	ret := _m.ctrl.Call(_m, "Height")
	ret0, _ := ret[0].(int)
	return ret0
}

// Height indicates an expected call of Height
func (_mr *MockSlotMockRecorder) Height() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Height")
}

// SetAttribute mocks base method
func (_m *MockSlot) SetAttribute(_param0 string, _param1 string) {
	_m.ctrl.Call(_m, "SetAttribute", _param0, _param1)
}

// SetAttribute indicates an expected call of SetAttribute
func (_mr *MockSlotMockRecorder) SetAttribute(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetAttribute", arg0, arg1)
}

// TrackID mocks base method
func (_m *MockSlot) TrackID() string {
	ret := _m.ctrl.Call(_m, "TrackID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TrackID indicates an expected call of TrackID
func (_mr *MockSlotMockRecorder) TrackID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TrackID")
}

// Width mocks base method
func (_m *MockSlot) Width() int {
	ret := _m.ctrl.Call(_m, "Width")
	ret0, _ := ret[0].(int)
	return ret0
}

// Width indicates an expected call of Width
func (_mr *MockSlotMockRecorder) Width() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Width")
}

// MockSupplier is a mock of Supplier interface
type MockSupplier struct {
	ctrl     *gomock.Controller
	recorder *MockSupplierMockRecorder
}

// MockSupplierMockRecorder is the mock recorder for MockSupplier
type MockSupplierMockRecorder struct {
	mock *MockSupplier
}

// NewMockSupplier creates a new mock instance
func NewMockSupplier(ctrl *gomock.Controller) *MockSupplier {
	mock := &MockSupplier{ctrl: ctrl}
	mock.recorder = &MockSupplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockSupplier) EXPECT() *MockSupplierMockRecorder {
	return _m.recorder
}

// ClickMode mocks base method
func (_m *MockSupplier) ClickMode() exchange.SupplierClickMode {
	ret := _m.ctrl.Call(_m, "ClickMode")
	ret0, _ := ret[0].(exchange.SupplierClickMode)
	return ret0
}

// ClickMode indicates an expected call of ClickMode
func (_mr *MockSupplierMockRecorder) ClickMode() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClickMode")
}

// ExcludedDemands mocks base method
func (_m *MockSupplier) ExcludedDemands() []string {
	ret := _m.ctrl.Call(_m, "ExcludedDemands")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ExcludedDemands indicates an expected call of ExcludedDemands
func (_mr *MockSupplierMockRecorder) ExcludedDemands() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ExcludedDemands")
}

// FloorCPM mocks base method
func (_m *MockSupplier) FloorCPM() int64 {
	ret := _m.ctrl.Call(_m, "FloorCPM")
	ret0, _ := ret[0].(int64)
	return ret0
}

// FloorCPM indicates an expected call of FloorCPM
func (_mr *MockSupplierMockRecorder) FloorCPM() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FloorCPM")
}

// Name mocks base method
func (_m *MockSupplier) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (_mr *MockSupplierMockRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

// Renderer mocks base method
func (_m *MockSupplier) Renderer() exchange.Renderer {
	ret := _m.ctrl.Call(_m, "Renderer")
	ret0, _ := ret[0].(exchange.Renderer)
	return ret0
}

// Renderer indicates an expected call of Renderer
func (_mr *MockSupplierMockRecorder) Renderer() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Renderer")
}

// Share mocks base method
func (_m *MockSupplier) Share() int {
	ret := _m.ctrl.Call(_m, "Share")
	ret0, _ := ret[0].(int)
	return ret0
}

// Share indicates an expected call of Share
func (_mr *MockSupplierMockRecorder) Share() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Share")
}

// SoftFloorCPM mocks base method
func (_m *MockSupplier) SoftFloorCPM() int64 {
	ret := _m.ctrl.Call(_m, "SoftFloorCPM")
	ret0, _ := ret[0].(int64)
	return ret0
}

// SoftFloorCPM indicates an expected call of SoftFloorCPM
func (_mr *MockSupplierMockRecorder) SoftFloorCPM() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SoftFloorCPM")
}

// TestMode mocks base method
func (_m *MockSupplier) TestMode() bool {
	ret := _m.ctrl.Call(_m, "TestMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TestMode indicates an expected call of TestMode
func (_mr *MockSupplierMockRecorder) TestMode() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TestMode")
}

// Type mocks base method
func (_m *MockSupplier) Type() string {
	ret := _m.ctrl.Call(_m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type
func (_mr *MockSupplierMockRecorder) Type() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Type")
}
