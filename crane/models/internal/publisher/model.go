package publisher

import (
	"time"

	"clickyab.com/exchange/crane/entity"
)

type activeStatus string
type platforms string

const (
	activeStatusTrue  activeStatus = "yes"
	activeStatusFalse activeStatus = "no"

	appPlatform  platforms = "app"
	vastPlatform platforms = "vast"
	webPlatform  platforms = "web"
)

// Publisher user model in database
// @Model {
//		table = publishers
//		primary = true, id
//		find_by = id
//		transaction = insert
//		list = yes
// }
type Publisher struct {
	IDd           int          `json:"id" db:"id"`
	UserID        int          `json:"user_id" db:"user_id"`
	FloorCPMm     int          `json:"floor_cpm" db:"floor_cpm"`
	SoftFloorCPMm int          `json:"soft_floor_cpm" db:"soft_floor_cpm"`
	Namee         string       `json:"name" db:"name"`
	BidType       int          `json:"bid_type" db:"bid_type"`
	UnderFloorr   int          `json:"under_floor" db:"under_floor"`
	Platform      platforms    `json:"platform" db:"platform"`
	Activee       activeStatus `json:"active" db:"active"`
	CreatedAt     *time.Time   `json:"created_at"  db:"created_at"`
	UpdatedAt     *time.Time   `json:"updated_at" db:"updated_at"`
}

// ID returns ID
func (*Publisher) ID() int64 {
	panic("implement me")
}

// FloorCPM returns FloorCPM
func (*Publisher) FloorCPM() int64 {
	panic("implement me")
}

// SoftFloorCPM returns SoftFloorCPM
func (*Publisher) SoftFloorCPM() int64 {
	panic("implement me")
}

// Name returns Name
func (*Publisher) Name() string {
	panic("implement me")
}

// Active returns Active
func (*Publisher) Active() bool {
	panic("implement me")
}

// AcceptedTarget returns AcceptedTarget
func (*Publisher) AcceptedTarget() entity.Target {
	panic("implement me")
}

// Attributes returns Attributes
func (*Publisher) Attributes(entity.PublisherAttributes) interface{} {
	panic("implement me")
}

// BIDType returns BIDType
func (*Publisher) BIDType() entity.BIDType {
	panic("implement me")
}

// MinCPC returns MinCPC
func (*Publisher) MinCPC() int64 {
	panic("implement me")
}

// AcceptedTypes returns AcceptedTypes
func (*Publisher) AcceptedTypes() []entity.AdType {
	panic("implement me")
}

// UnderFloor returns UnderFloor
func (*Publisher) UnderFloor() bool {
	panic("implement me")
}

// Supplier returns Supplier
func (*Publisher) Supplier() entity.Supplier {
	panic("implement me")
}
