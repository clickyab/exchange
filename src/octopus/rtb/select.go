package rtb

import (
	"octopus/exchange"
	"sort"
)

// SelectCPM is the simplest way to bid. sort the value, return the
func SelectCPM(imp exchange.Impression, all map[string][]exchange.Advertise) map[string]exchange.Advertise {
	res := make(map[string]exchange.Advertise, len(all))

	for id := range all {
		if len(all[id]) == 0 {
			res[id] = nil
			continue
		}
		sorted := sortedAd(all[id])
		sort.Sort(sorted)

		res[id] = sorted[0]
		lower := imp.Source().SoftFloorCPM()
		if lower > res[id].MaxCPM() {
			lower = imp.Source().FloorCPM()
		}
		if len(sorted) > 1 && sorted[1].MaxCPM() > lower {
			lower = sorted[1].MaxCPM()
		}

		if lower < res[id].MaxCPM() {
			res[id].SetWinnerCPM(lower + 1)
		} else {
			res[id].SetWinnerCPM(res[id].MaxCPM())
		}
	}

	return res
}