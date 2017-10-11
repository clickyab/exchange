package jsonbackend

import "clickyab.com/exchange/octopus/exchange"

// winner is result of selectCPM
type Winner interface {
	Bid() exchange.Bid
	Price() int64
}

type winners struct {
	bid   exchange.Bid
	price int64
}

func (w winners) Bid() exchange.Bid {
	return w.bid
}

func (w winners) Price() int64 {
	return w.price
}
