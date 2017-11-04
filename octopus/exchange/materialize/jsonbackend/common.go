package jsonbackend

import (
	"clickyab.com/exchange/octopus/exchange"
)

func requestToMap(req exchange.BidRequest) map[string]interface{} {
	return map[string]interface{}{
		"id":         req.ID(),
		"inventory":  inventoryToMap(req.Inventory()),
		"impression": impressionToMap(req.Imp()),
		"time":       req.Time(),
	}
}

func demandToMap(dmn exchange.Demand) map[string]interface{} {
	return map[string]interface{}{
		"name": dmn.Name(),
	}

}

// no ad markup, dont think we need it
func responseToMap(res exchange.BidResponse) map[string]interface{} {
	return map[string]interface{}{
		"id":   res.ID(),
		"bids": bidsToMap(res.Bids()),
	}

}

func bidsToMap(bids []exchange.Bid) []map[string]interface{} {
	response := []map[string]interface{}{}
	for _, val := range bids {
		response = append(response, map[string]interface{}{
			"id":         val.ID(),
			"imp_id":     val.ImpID(),
			"price":      val.Price(),
			"win_url":    val.WinURL(),
			"categories": val.Categories(),
			"ad_id":      val.AdID(),
			"ad_height":  val.AdHeight(),
			"ad_width":   val.AdWidth(),
			"ad_domains": val.AdDomains(),
			"demand":     demandToMap(val.Demand()),
		})
	}
	return response
}

func winnerBidToMap(bid exchange.Bid) map[string]interface{} {
	return map[string]interface{}{
		"height":     bid.AdHeight(),
		"id":         bid.ID(),
		"ad_domains": bid.AdDomains(),
		"url":        bid.WinURL(),
		"width":      bid.AdWidth(),
		"winner_cpm": bid.Price(),
		"imp_id":     bid.ImpID(),
		"demand":     demandToMap(bid.Demand()),
	}
}

func inventoryToMap(inv exchange.Inventory) map[string]interface{} {
	return map[string]interface{}{
		"supplier":       supplierToMap(inv.Supplier()),
		"name":           inv.Name(),
		"soft_floor_cpm": inv.SoftFloorCPM(),
		"floor_cpm":      inv.FloorCPM(),
		"domain":         inv.Domain(),
	}
}

func supplierToMap(sup exchange.Supplier) map[string]interface{} {
	return map[string]interface{}{
		"floor_cpm":      sup.FloorCPM(),
		"soft_floor_cpm": sup.SoftFloorCPM(),
		"name":           sup.Name(),
		"share":          sup.Share(),
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
		"bid":     winnerBidToMap(bid),
		"request": requestToMap(bq),
	}
}

func showToMap(demand string, winner int64, supplier string, publisher string, profit int64) map[string]interface{} {
	return map[string]interface{}{
		"demand_name": demand,
		"price":       winner,
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
