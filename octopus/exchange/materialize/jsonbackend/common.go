package jsonbackend

import (
	"clickyab.com/exchange/octopus/exchange"
)

func requestToMap(req exchange.BidRequest) map[string]interface{} {
	return map[string]interface{}{
		"track_id":   req.ID(),
		"ip":         req.Device().IP(),
		"user_agent": req.Device().UserAgent(),
		"supplier":   supplierToMap(req.Inventory().Supplier()),
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

func winnerBidToMap(bid exchange.Bid) map[string]interface{} {
	return map[string]interface{}{
		"height":        bid.AdHeight(),
		"id":            bid.ID(),
		"landing":       bid.AdDomains(),
		"max_cpm":       bid.Price(),
		"track_id":      bid.ID(),
		"url":           bid.WinURL(),
		"width":         bid.AdWidth(),
		"winner_cpm":    bid.Price(),
		"slot_track_id": bid.ImpID(),
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

func winnerToMap(bq exchange.BidRequest, bid exchange.Bid) map[string]interface{} {
	return map[string]interface{}{
		"demand":    demandToMap(bid.Demand()),
		"request":   requestToMap(bq),
		"advertise": winnerBidToMap(bid),
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
		"type":              exchange.AdTypeBanner,
		"blocked_type":      banner.BlockedTypes(),
		"blocked_attribute": banner.BlockedAttributes(),
		"mimes":             banner.Mimes(),
		"attributes":        banner.Attributes(),
	}
}

func videoToMap(video exchange.Video) map[string]interface{} {
	return map[string]interface{}{
		"width":             video.Width(),
		"height":            video.Height(),
		"linearity":         video.Linearity(),
		"type":              exchange.AdTypeVideo,
		"blocked_attribute": video.BlockedAttributes(),
		"mimes":             video.Mimes(),
		"attributes":        video.Attributes(),
	}
}

func nativeToMap(native exchange.Native) map[string]interface{} {
	return map[string]interface{}{
		"extension":  native.Request(),
		"is_valid":   native.IsExtValid(),
		"type":       exchange.AdTypeNative,
		"ad_length":  native.AdLength(),
		"attributes": native.Attributes(),
	}
}
