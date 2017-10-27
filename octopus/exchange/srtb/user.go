package srtb

import "clickyab.com/exchange/octopus/srtb"

type user struct {
	inner *srtb.User
}

func (u *user) ID() string {
	return u.inner.ID
}

func (u *user) Attributes() map[string]interface{} {
	return make(map[string]interface{})
}
