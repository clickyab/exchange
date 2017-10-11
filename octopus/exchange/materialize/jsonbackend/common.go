package jsonbackend

import (
	"clickyab.com/exchange/octopus/exchange"
)

func requestToMap(req exchange.BidRequest) map[string]interface{} {
	return map[string]interface{}{
		"track_id":   req.ID(),
		"ip":         req.Device().IP(),
		"user_agent": req.Device().UserAgent(),
		"supplier":   supplierToMap(req.Supplier()),
		"inventory":  inventoryToMap(req.Inventory()),
		"location":   locationToMap(req.Device().Geo()),
		//"attributes":  req.,
		"impression":       impressionToMap(req.Imp()),
		"blocked_category": req.BlockedCategories(),
		"platform":         req.Device().DeviceType(),
		"time":             req.Time(),
	}
}

func demandToMap(dmn exchange.Demand) map[string]interface{} {
	return map[string]interface{}{
		"name":                 dmn.Name(),
		"call_rate":            dmn.CallRate(),
		"handicap":             dmn.Handicap(),
		"white_list_countries": dmn.WhiteListCountries(),
		"excluded_suppliers":   dmn.WhiteListCountries(),
	}

}

// no ad markup, dont think we need it
func bidsToMap(bids []exchange.Bid) []map[string]interface{} {
	response := []map[string]interface{}{}
	for i := range bids {
		response = append(response, map[string]interface{}{
			"id":         bids[i].ID(),
			"imp_id":     bids[i].ImpID(),
			"price":      bids[i].Price(),
			"win_url":    bids[i].WinURL(),
			"categories": bids[i].Categories(),
			"ad_id":      bids[i].AdID(),
			"ad_height":  bids[i].AdHeight(),
			"ad_width":   bids[i].AdWidth(),
			"ad_domains": bids[i].AdDomains(),
		})
	}

	return response
}

func winnerBidToMap(winner winners) map[string]interface{} {
	return map[string]interface{}{
		"height":        winner.bid.AdHeight(),
		"id":            winner.bid.ID(),
		"landing":       winner.bid.AdDomains(),
		"max_cpm":       winner.bid.Price(),
		"track_id":      winner.bid.ID(),
		"url":           winner.bid.WinURL(),
		"width":         winner.bid.AdWidth(),
		"winner_cpm":    winner.price,
		"slot_track_id": winner.bid.ImpID(),
	}
}

func inventoryToMap(inv exchange.Inventory) map[string]interface{} {
	return map[string]interface{}{
		"name":           inv.Name(),
		"soft_floor_cpm": inv.SoftFloorCPM(),
		"floor_cpm":      inv.FloorCPM(),
		"attributes":     inv.Attributes(),
	}
}

func supplierToMap(sup exchange.Supplier) map[string]interface{} {
	return map[string]interface{}{
		"floor_cpm":        sup.FloorCPM(),
		"soft_floor_cpm":   sup.SoftFloorCPM(),
		"name":             sup.Name(),
		"share":            sup.Share(),
		"excluded_demands": sup.ExcludedDemands(),
	}
}

func locationToMap(loc exchange.Location) map[string]interface{} {
	return map[string]interface{}{
		"country":  loc.Country(),
		"province": loc.Region(),
		"lat_lon":  loc.LatLon(),
	}
}

func impressionToMap(imps []exchange.Impression) []map[string]interface{} {
	var resp []map[string]interface{}
	for i := range imps {
		temp := map[string]interface{}{}
		switch imps[i].Type() {
		case exchange.AdTypeBanner:
			temp["banner"] = bannerToMap(imps[i].Banner())
		case exchange.AdTypeNative:
			temp["native"] = nativeToMap(imps[i].Native())
		case exchange.AdTypeVideo:
			temp["video"] = videoToMap(imps[i].Video())
		}

		temp["id"] = imps[i].ID()
		temp["bid_floor"] = imps[i].BidFloor()
		temp["type"] = imps[i].Type()
		temp["secure"] = imps[i].Secure()

		resp = append(resp, temp)
	}

	return resp
}

func winnerToMap(req exchange.BidRequest, winner winners, slotID string) map[string]interface{} {
	return map[string]interface{}{
		"demand":    demandToMap(winner.bid.BidResponse().Demand()),
		"request":   requestToMap(req),
		"advertise": winnerBidToMap(winner),
		"slot_id":   slotID,
	}
}

func showToMap(trackID, demand, slotID, adID string, winner int64, supplier string, publisher string, profit int64) map[string]interface{} {
	return map[string]interface{}{
		"track_id":    trackID,
		"demand_name": demand,
		"price":       winner,
		"slot_id":     slotID,
		"ad_id":       adID,
		"supplier":    supplier,
		"publisher":   publisher,
		"profit":      profit,
	}
}

func bannerToMap(banner exchange.Banner) map[string]interface{} {
	return map[string]interface{}{
		"id":                banner.ID(),
		"width":             banner.Width(),
		"height":            banner.Height(),
		"type":              banner.Type(),
		"blocked_type":      banner.BlockedTypes(),
		"blocked_attribute": banner.BlockedAttributes(),
		"mimes":             banner.Mimes(),
		"attributes":        banner.Attributes(),
	}
}

func videoToMap(banner exchange.Video) map[string]interface{} {
	return map[string]interface{}{
		"width":             banner.Width(),
		"height":            banner.Height(),
		"linearity":         banner.Linearity(),
		"blocked_attribute": banner.BlockedAttributes(),
		"mimes":             banner.Mimes(),
		"attributes":        banner.Attributes(),
	}
}

func nativeToMap(banner exchange.Native) map[string]interface{} {
	return map[string]interface{}{
		"extension":  banner.Extension(),
		"is_valid":   banner.IsExtValid(),
		"ad_length":  banner.AdLength(),
		"attributes": banner.Attributes(),
	}
}
