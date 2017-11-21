package biding

import (
	"clickyab.com/exchange/octopus/exchange"
)

type sortedBid []exchange.Bid

func (sa sortedBid) Len() int {
	return len(sa)
}

func (sa sortedBid) Less(i, j int) bool {
	cpi := sa[i].Price() * float64(sa[i].Demand().Handicap())
	cpj := sa[j].Price() * float64(sa[j].Demand().Handicap())
	return cpi > cpj
}

func (sa sortedBid) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}
