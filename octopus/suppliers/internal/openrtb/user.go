package openrtb

import (
	"clickyab.com/exchange/octopus/exchange"
	"github.com/bsm/openrtb"
)

type user struct {
	id         string
	attributes map[string]interface{}
}

func (u user) Attributes() map[string]interface{} {
	return u.attributes
}

func (u user) ID() string {
	return u.id
}

func userAttributes(r *openrtb.User) map[string]interface{} {
	return map[string]interface{}{
		"BuyerID":    r.BuyerID,
		"BuyerUID":   r.BuyerUID,
		"YOB":        r.YOB,
		"Gender":     r.Gender,
		"Keywords":   r.Keywords,
		"CustomData": r.CustomData,
		"Geo":        r.Geo,
		"Data":       r.Data,
		"Ext":        r.Ext,
	}
}
func newUser(u *openrtb.User) exchange.User {
	return user{
		id:         u.ID,
		attributes: userAttributes(u),
	}
}
