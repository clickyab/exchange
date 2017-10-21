package restful

import (
	"time"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/random"
	"reflect"
)

//type bidRequestRest struct {
//	Categories    []exchange.Category `json:"categories"`
//	Imps          []*impRest          `json:"imps"`
//	Loc           exchange.Location   `json:"location"`
//	Mega          string              `json:"track_id"`
//	PTI           string              `json:"page_track_id"`
//	Pub           *restPublisher      `json:"source"`
//	Schm          string              `json:"scheme"`
//	SIP           string              `json:"ip"`
//	PlatDevice    exchange.DeviceType `json:"platform"`
//	STime         time.Time           `json:"time"`
//	UA            string              `json:"user_agent"`
//	UnderFloorCPM bool                `json:"under_floor"`
//	UTI           string              `json:"user_track_id"`
//
//	Attr map[string]interface{} `json:"attributes"`
//
//	dum    []exchange.Impression
//	latlon exchange.LatLon
//}
//
//func (ir *bidRequestRest) ID() string {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) Imp() []exchange.Impression {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) Inventory() exchange.Inventory {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) Device() exchange.Device {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) User() exchange.User {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) Test() bool {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) AuctionType() exchange.AuctionType {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) TMax() time.Duration {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) WhiteList() []string {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) BlackList() []string {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) AllowedLanguage() []string {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) BlockedCategories() []string {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) BlockedAdvertiserDomain() []string {
//	panic("implement me")
//}
//
//func (ir *bidRequestRest) UserTrackID() string {
//	return ir.UTI
//}
//
//func (ir *bidRequestRest) PageTrackID() string {
//	return ir.PTI
//}
//
//type location struct {
//	TheCountry exchange.Country `json:"country"`
//	TheRegion  exchange.Region  `json:"province"`
//	TheLatLon  exchange.LatLon  `json:"latlon"`
//	TheISP     exchange.ISP     `json:"isp"`
//}
//
//func (l location) Region() exchange.Region {
//	panic("implement me")
//}
//
//func (l location) ISP() exchange.ISP {
//	return l.TheISP
//}
//
//func (l location) Country() exchange.Country {
//	return l.TheCountry
//}
//
//func (l location) Province() exchange.Region {
//	return l.TheRegion
//}
//
//func (l location) LatLon() exchange.LatLon {
//	return l.TheLatLon
//}
//
//func (ir *bidRequestRest) TrackID() string {
//	if ir.Mega == "" {
//		ir.Mega = <-random.ID
//	}
//
//	return ir.Mega
//}
//
//func (ir bidRequestRest) IP() net.IP {
//	return net.ParseIP(ir.SIP)
//}
//
//func (ir bidRequestRest) Scheme() string {
//	if ir.Schm != "https" {
//		ir.Schm = "http"
//	}
//	return ir.Schm
//}
//
//func (ir bidRequestRest) UserAgent() string {
//	return ir.UA
//}
//
//func (ir bidRequestRest) Source() exchange.Inventory {
//	return ir.Pub
//}
//
//func (ir bidRequestRest) Location() exchange.Location {
//	return ir.Loc
//}
//
//func (ir bidRequestRest) Attributes() map[string]interface{} {
//	return ir.Attr
//}
//
//func (ir *bidRequestRest) Slots() []exchange.Impression {
//	if ir.dum == nil {
//		ir.dum = make([]exchange.Impression, len(ir.Imps))
//		for i := range ir.Imps {
//			ir.dum[i] = ir.Imps[i]
//		}
//	}
//	return ir.dum
//}
//
//func (ir bidRequestRest) Category() []exchange.Category {
//	return ir.Categories
//}
//
//func (ir bidRequestRest) Platform() exchange.DeviceType {
//	return ir.PlatDevice
//}
//
//func (ir bidRequestRest) UnderFloor() bool {
//	return ir.UnderFloorCPM
//}
//
//func (ir bidRequestRest) Raw() interface{} {
//	return ir
//}
//
//func (ir *bidRequestRest) extractData() {
//	d := ip2location.IP2Location(ir.SIP)
//	//logrus.Debug(d)
//	ir.Loc = location{
//		TheCountry: exchange.Country{
//			Name:  d.CountryLong,
//			ISO:   d.CountryShort,
//			Valid: d.CountryLong != "-",
//		},
//
//		TheRegion: exchange.Region{
//			Valid: d.Region != "-",
//			Name:  d.Region,
//		},
//
//		TheLatLon: ir.latlon,
//
//		TheISP: exchange.ISP{
//			Valid: d.ISP != "-",
//			Name:  d.ISP,
//		},
//	}
//
//}
//
//func (ir *bidRequestRest) Time() time.Time {
//	return ir.STime
//}
//
//func newImpressionFromAppRequest(sup exchange.Supplier, r *requestBody) (*bidRequestRest, error) {
//	resp := bidRequestRest{
//		Schm:          r.Scheme,
//		UTI:           r.UserTrackID,
//		PTI:           r.PageTrackID,
//		SIP:           r.IP,
//		UA:            r.App.UserAgent,
//		PlatDevice:    exchange.DeviceTypePhone,
//		Categories:    r.Categories,
//		Imps:          r.Slots,
//		Mega:          <-random.ID,
//		UnderFloorCPM: r.UnderFloor,
//		Pub:           r.Publisher,
//		Attr: map[string]interface{}{
//			"network":     r.App.Network,
//			"brand":       r.App.Brand,
//			"cid":         r.App.CID,
//			"lac":         r.App.LAC,
//			"mcc":         r.App.MCC,
//			"mnc":         r.App.MNC,
//			"language":    r.App.Language,
//			"model":       r.App.Model,
//			"operator":    r.App.Operator,
//			"os_identity": r.App.OSIdentity,
//		},
//	}
//	lat, lon, err := gmaps.LockUp(r.App.MCC, r.App.MNC, r.App.LAC, r.App.CID)
//	resp.latlon = exchange.LatLon{
//		Valid: err == nil,
//		Lat:   lat,
//		Lon:   lon,
//	}
//	resp.Pub.sup = sup
//	resp.STime = time.Now()
//	resp.extractData()
//	return &resp, nil
//}
//
//func newImpressionFromVastRequest(sup exchange.Supplier, r *requestBody) (*bidRequestRest, error) {
//	resp := bidRequestRest{
//		Schm:          r.Scheme,
//		UTI:           r.UserTrackID,
//		PTI:           r.PageTrackID,
//		SIP:           r.IP,
//		UA:            r.Vast.UserAgent,
//		PlatDevice:    exchange.DeviceTypePC,
//		Categories:    r.Categories,
//		Imps:          r.Slots,
//		Mega:          <-random.ID,
//		UnderFloorCPM: r.UnderFloor,
//
//		Attr: map[string]interface{}{
//			"referrer": r.Vast.Referrer,
//			"parent":   r.Vast.Parent,
//		},
//		Pub: r.Publisher,
//	}
//	resp.Pub.sup = sup
//	resp.STime = time.Now()
//	resp.extractData()
//	return &resp, nil
//}

func newImpressionFromWebRequest(sup exchange.Supplier, r *requestBody) (*bidRequestRest, error) {

	//r.Publisher.sup = sup
	//resp := bidRequestRest{
	//	Schm:          r.Scheme,
	//	UTI:           r.UserTrackID,
	//	PTI:           r.PageTrackID,
	//	SIP:           r.IP,
	//	UA:            r.Web.UserAgent,
	//	PlatDevice:    exchange.DeviceTypePC,
	//	Categories:    r.Categories,
	//	Imps:          r.Slots,
	//	Mega:          <-random.ID,
	//	UnderFloorCPM: r.UnderFloor,
	//
	//	Attr: map[string]interface{}{
	//		"referrer": r.Web.Referrer,
	//		"parent":   r.Web.Parent,
	//	},
	//	Pub: r.Publisher,
	//}
	//resp.STime = time.Now()
	//resp.extractData()
	return &resp, nil
}
