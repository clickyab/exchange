package demands

import (
	"time"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/sirupsen/logrus"
)

type rawPub struct {
	// Name of publisher
	Name string `json:"name"`
	// FloorCPM is the floor cpm for publisher
	FloorCPM int64 `json:"floor_cpm"`
	// SoftFloorCPM is the soft version of floor cpm. if the publisher ahs it, then the system
	// try to use this as floor, but if this is not available, the FloorCPM is used
	SoftFloorCPM int64 `json:"soft_floor_cpm"`
	// FAttributes is the generic attribute system
	Attributes map[string]interface{} `json:"attributes"`
	// Supplier the supplier
	Supplier string `json:"supplier"`
}

func getRawPub(in exchange.Inventory) rawPub {
	return rawPub{
		Name:         in.Name(),
		Attributes:   in.Attributes(),
		SoftFloorCPM: in.SoftFloorCPM(),
		FloorCPM:     in.FloorCPM(),
		Supplier:     in.Supplier().Name(),
	}
}

type rawLocation struct {
	// Country get the country if available
	Country exchange.Country `json:"country"`
	// Region get the province of request if available
	Region exchange.Region `json:"province"`
	// LatLon return the latitude longitude if any
	LatLon exchange.LatLon `json:"lat_lon"`
	// ISP return the isp
	ISP exchange.ISP `json:"isp"`
}

func getRawLocation(in exchange.Location) rawLocation {
	return rawLocation{
		LatLon:  in.LatLon(),
		Region:  in.Region(),
		Country: in.Country(),
		ISP:     in.ISP(),
	}
}

type rawImps struct {
	// Size return the primary size of this slot
	Width  int `json:"width"`
	Height int `json:"height"`
	// ID is an string for this slot, its a random at first but the value is not changed at all other calls
	ImpID string `json:"track_id"`
	// data needed for supplier and it's optional
	FAttributes map[string]interface{} `json:"attributes"`
}

func (*rawImps) ID() string {
	panic("implement me")
}

func (*rawImps) BidFloor() float64 {
	panic("implement me")
}

func (*rawImps) Banner() exchange.Banner {
	panic("implement me")
}

func (*rawImps) Video() exchange.Video {
	panic("implement me")
}

func (*rawImps) Native() exchange.Native {
	panic("implement me")
}

func (*rawImps) Attributes() map[string]interface{} {
	panic("implement me")
}

func (*rawImps) Type() exchange.ImpressionType {
	panic("implement me")
}

func (*rawImps) Secure() bool {
	panic("implement me")
}

func getRawImps(in exchange.BidRequest) []rawImps {
	res := make([]rawImps, len(in.Imp()))
	for i := range in.Imp() {
		if in.Imp()[i].Type() == exchange.AdTypeBanner {
			res[i] = rawImps{
				Width:       in.Imp()[i].Banner().Width(),
				Height:      in.Imp()[i].Banner().Height(),
				ImpID:       in.Imp()[i].ID(),
				FAttributes: in.Imp()[i].Attributes(),
			}
		} else if in.Imp()[i].Type() == exchange.AdTypeNative {
			res[i] = rawImps{
				ImpID:       in.Imp()[i].ID(),
				FAttributes: in.Imp()[i].Attributes(),
			}
		} else if in.Imp()[i].Type() == exchange.AdTypeVideo {
			res[i] = rawImps{
				Width:       in.Imp()[i].Video().Width(),
				Height:      in.Imp()[i].Video().Height(),
				ImpID:       in.Imp()[i].ID(),
				FAttributes: in.Imp()[i].Attributes(),
			}
		} else {
			logrus.Panic("[BUG] imp type not specified")
		}

	}

	return res
}

type rawBidRequest struct {
	TrackID     string `json:"track_id"`
	IP          string `json:"ip"`
	UserAgent   string `json:"user_agent"`
	PageTrackID string `json:"page_track_id"`
	UserTrackID string `json:"user_track_id"`
	// Source return the publisher that this client is going into system from that
	Source rawPub `json:"source"`
	// Location of the request
	Location rawLocation `json:"location"`
	// FAttributes is the generic attribute system
	Attributes map[string]interface{} `json:"attributes"`
	// Imp is the slot for this request
	Imps []rawImps `json:"slots"`
	// Category returns category obviously
	Category []exchange.Category `json:"category"`
	// Platform return the publisher Platform
	Platform exchange.DeviceType `json:"platform"`
}

func (*rawBidRequest) ID() string {
	panic("implement me")
}

func (*rawBidRequest) Imp() []exchange.Impression {
	panic("implement me")
}

func (*rawBidRequest) Inventory() exchange.Inventory {
	panic("implement me")
}

func (*rawBidRequest) Device() exchange.Device {
	panic("implement me")
}

func (*rawBidRequest) User() exchange.User {
	panic("implement me")
}

func (*rawBidRequest) Test() bool {
	panic("implement me")
}

func (*rawBidRequest) AuctionType() exchange.AuctionType {
	panic("implement me")
}

func (*rawBidRequest) TMax() time.Duration {
	panic("implement me")
}

func (*rawBidRequest) WhiteList() []string {
	panic("implement me")
}

func (*rawBidRequest) BlackList() []string {
	panic("implement me")
}

func (*rawBidRequest) AllowedLanguage() []string {
	panic("implement me")
}

func (*rawBidRequest) BlockedCategories() []string {
	panic("implement me")
}

func (*rawBidRequest) BlockedAdvertiserDomain() []string {
	panic("implement me")
}

func (*rawBidRequest) Time() time.Time {
	panic("implement me")
}

func getRawBidRequest(bq exchange.BidRequest) interface{} {
	return rawBidRequest{
		TrackID:     bq.ID(),
		PageTrackID: bq.ID(),
		UserTrackID: bq.User().ID(),
		IP:          bq.Device().IP(),
		UserAgent:   bq.Device().UserAgent(),
		Attributes:  bq.Attributes(),
		Category:    bq.Inventory().Cat(),
		Platform:    bq.Device().DeviceType(),
		Imps:        getRawImps(bq),
		Location:    getRawLocation(bq.Device().Geo()),
		Source:      getRawPub(bq.Inventory()),
	}

}
