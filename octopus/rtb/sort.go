package rtb

import (
	"clickyab.com/exchange/octopus/exchange"
)

type sortedBid []exchange.Bid

func (sa sortedBid) Len() int {
	return len(sa)
}

func (sa sortedBid) Less(i, j int) bool {
	cpi := sa[i].Price() * sa[i].Demand().Handicap()
	cpj := sa[j].Price() * sa[j].Demand().Handicap()
	return cpi > cpj
}

func (sa sortedBid) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}
