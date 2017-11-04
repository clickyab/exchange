package srtb

import "github.com/clickyab/simple-rtb"

type user struct {
	inner *srtb.User
}

func (u *user) ID() string {
	return u.inner.ID
}

func (u *user) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}
