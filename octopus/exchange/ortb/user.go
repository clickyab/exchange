package ortb

import "github.com/bsm/openrtb"

// User ortb user
type User struct {
	inner *openrtb.User
}

// ID return ortb ID
func (u *User) ID() string {
	return u.inner.ID
}

// Attributes return ortb Attributes
func (u *User) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"BuyerID":    u.inner.BuyerID,
		"BuyerUID":   u.inner.BuyerUID,
		"YOB":        u.inner.YOB,
		"Gender":     u.inner.Gender,
		"Keywords":   u.inner.Keywords,
		"CustomData": u.inner.CustomData,
		"Geo":        u.inner.Geo,
		"Data":       u.inner.Data,
		"Ext":        u.inner.Ext,
	}
}
