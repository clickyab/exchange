package rtb

import (
	"clickyab.com/exchange/octopus/exchange"
)

type sortedAd []exchange.Bid

func (sa sortedAd) Len() int {
	return len(sa)
}

func (sa sortedAd) Less(i, j int) bool {
	cpi := sa[i].Price() * sa[i].Demand().Handicap()
	cpj := sa[j].Price() * sa[j].Demand().Handicap()
	return cpi > cpj
}

func (sa sortedAd) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}
