package ad

import (
	"time"

	"clickyab.com/exchange/crane/entity"
)

type activeStatus string

const (
	activeStatusTrue  activeStatus = "yes"
	activeStatusFalse activeStatus = "no"
)

// Ad ad model in database
// @Model {
//		table = ads
//		primary = true, id
//		find_by = id
//		transaction = insert
//		list = yes
// }
type Ad struct {
	AdID       int          `json:"id" db:"id"`
	AdType     string       `json:"type" db:"type"`
	Width      int          `json:"width" db:"width"`
	Height     int          `json:"height" db:"height"`
	Active     activeStatus `json:"active" db:"active"`
	UserID     int          `json:"user_id" db:"user_id"`
	CampaignId int          `json:"campaign_id" db:"campaign_id"`
	URL        string       `json:"url" db:"url"`
	Attribute  string       `json:"attribute" db:"attribute"`
	CreatedAt  *time.Time   `json:"created_at" db:"created_ad"`
	UpdatedAt  *time.Time   `json:"updated_at" db:"updated_ad"`
}

// ID returns ID
func (*Ad) ID() int64 {
	panic("implement me")
}

// Type returns Type
func (*Ad) Type() entity.AdType {
	panic("implement me")
}

// Campaign returns Campaign
func (*Ad) Campaign() entity.Campaign {
	panic("implement me")
}

// SetCPM returns SetCPM
func (*Ad) SetCPM(int64) {
	panic("implement me")
}

// CPM returns CPM
func (*Ad) CPM() int64 {
	panic("implement me")
}

// SetWinnerBID returns SetWinnerBID
func (*Ad) SetWinnerBID(int64) {
	panic("implement me")
}

// WinnerBID returns WinnerBID
func (*Ad) WinnerBID() int64 {
	panic("implement me")
}

// AdCTR returns AdCTR
func (*Ad) AdCTR() float64 {
	panic("implement me")
}

// SetCTR returns SetCTR
func (*Ad) SetCTR(float64) {
	panic("implement me")
}

// CTR returns CTR
func (*Ad) CTR() float64 {
	panic("implement me")
}

// Size returns Size
func (*Ad) Size() int {
	panic("implement me")
}

// Category returns Category
func (*Ad) Category() []entity.Category {
	panic("implement me")
}

// Copy returns Copy
func (*Ad) Copy() entity.Advertise {
	panic("implement me")
}

// Capping returns Capping
func (*Ad) Capping() entity.Capping {
	panic("implement me")
}

// SetCapping returns SetCapping
func (*Ad) SetCapping(entity.Capping) {
	panic("implement me")
}

// BlackListPublisher returns BlackListPublisher
func (*Ad) BlackListPublisher() []int64 {
	panic("implement me")
}

// WhiteListPublisher returns WhiteListPublisher
func (*Ad) WhiteListPublisher() []int64 {
	panic("implement me")
}

// AllowedOS returns AllowedOS
func (*Ad) AllowedOS() []int64 {
	panic("implement me")
}

// Country returns Country
func (*Ad) Country() []int64 {
	panic("implement me")
}

// Province returns Province
func (*Ad) Province() []int64 {
	panic("implement me")
}

// LanLon returns LanLon
func (*Ad) LanLon() (float64, float64) {
	panic("implement me")
}
