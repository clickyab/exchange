package mocks

import "clickyab.com/exchange/octopus/exchange"

type Video struct {
}

func (*Video) Width() int {
	panic("implement me")
}

func (*Video) Height() int {
	panic("implement me")
}

func (*Video) Mimes() []string {
	panic("implement me")
}

func (*Video) Linearity() bool {
	panic("implement me")
}

func (*Video) BlockedAttributes() []exchange.CreativeAttribute {
	panic("implement me")
}

func (*Video) Attributes() map[string]interface{} {
	panic("implement me")
}
