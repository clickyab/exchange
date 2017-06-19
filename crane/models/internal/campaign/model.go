package campaign

import (
	"time"

	"clickyab.com/exchange/crane/entity"
)

// campaign user model in database
// @Model {
//		table = campaigns
//		primary = true, id
//		find_by = id
//		transaction = insert
//		list = yes
// }
type campaign struct {
	CampainID  int        `json:"id" db:"id"`
	UserID     int        `json:"user_id" db:"user_id"`
	Namee      string     `json:"name" db:"name"`
	MaxBid     int        `json:"max_bid" db:"max_bid"`
	Frequenccy int        `json:"frequency" db:"frequency"`
	CreatedAt  *time.Time `json:"created_at"  db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at"`
}

// ID returns ID
func (*campaign) ID() int64 {
	panic("implement me")
}

// Name returns Name
func (*campaign) Name() string {
	panic("implement me")
}

// MaxBID returns MaxBID
func (*campaign) MaxBID() int64 {
	panic("implement me")
}

// Frequency returns Frequency
func (*campaign) Frequency() int {
	panic("implement me")
}

// Target returns Target
func (*campaign) Target() []entity.Target {
	panic("implement me")
}
