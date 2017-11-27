package manager

import (
	"fmt"
	"strings"

	"clickyab.com/exchange/octopus/models"
	"clickyab.com/exchange/octopus/workers/internal/datamodels"
)

func hasDataSupplierDemand(tm *datamodels.TableModel) (bool, string) {
	b := tm.SupDemRequestOut+tm.SupDemAdOut+tm.SupDemAdWin+tm.SupDemAdIn+tm.SupDemAdDeliver+tm.Click > 0 || tm.SupDemBidIn+tm.Profit+tm.SupDemBidWin+tm.SupDemBidDeliver > 0
	if b {
		//(supplier,demand,source,time_id,request_out,bid_in,ad_out,ad_win,bid_win,ad_in,ad_deliver,bid_deliver,profit,click)
		return true, fmt.Sprintf(`("%s", "%s", "%s", %d, %d,%f, %d, %d, %f,%d, %d, %f,%f,%d)`,
			tm.Supplier,
			tm.Demand,
			tm.Source,
			tm.Time,
			tm.SupDemRequestOut,
			tm.SupDemBidIn,
			tm.SupDemAdOut,
			tm.SupDemAdWin,
			tm.SupDemBidWin,
			tm.SupDemAdIn,
			tm.SupDemAdDeliver,
			tm.SupDemBidDeliver,
			tm.Profit,
			tm.Click,
		)
	}
	return false, ""
}

func hasDataSupplier(tm *datamodels.TableModel) (bool, string) {
	b := tm.SupRequestIn+tm.SupAdIn+tm.SupAdOut+tm.SupAdDeliver+tm.Click > 0 || tm.SupBidDeliver+tm.Profit+tm.SupBidOut > 0
	if b {
		// (supplier,source,time_id,request_in,ad_in,ad_out,bid_out,ad_deliver,bid_deliver,profit,click)
		return true, fmt.Sprintf(`("%s","%s",%d,%d,%d,%d,%f,%d,%f,%f,%d)`,
			tm.Supplier,
			tm.Source,
			tm.Time,
			tm.SupRequestIn,
			tm.SupAdIn,
			tm.SupAdOut,
			tm.SupBidOut,
			tm.SupAdDeliver,
			(tm.SupBidDeliver)-(tm.Profit),
			tm.Profit,
			tm.Click,
		)
	}

	return false, ""
}

const supDemSrcTable = `INSERT INTO sup_dem_src
(supplier,demand,source,time_id,request_out,bid_in,ad_out,ad_win,bid_win,ad_in,ad_deliver,bid_deliver,profit,click) VALUES
%s
ON DUPLICATE KEY UPDATE
 request_out=request_out+VALUES(request_out),
 bid_in=bid_in+VALUES(bid_in),
 ad_out=ad_out+VALUES(ad_out),
 ad_win=ad_win+VALUES(ad_win),
 bid_win=bid_win+VALUES(bid_win),
 ad_in=ad_in+VALUES(ad_in),
 ad_deliver=ad_deliver+VALUES(ad_deliver),
 bid_deliver=bid_deliver+VALUES(bid_deliver),
 profit=profit+VALUES(profit),
 click=click+VALUES(click)
`
const supSrcTable = `INSERT INTO sup_src
(supplier,source,time_id,request_in,ad_in,ad_out,bid_out,ad_deliver,bid_deliver,profit,click) VALUES
%s
ON DUPLICATE KEY UPDATE
 request_in=request_in+VALUES(request_in),
 ad_in=ad_in+VALUES(ad_in),
 ad_out=ad_out+VALUES(ad_out),
 bid_out=bid_out+VALUES(bid_out),
 ad_deliver=ad_deliver+VALUES(ad_deliver),
 bid_deliver=bid_deliver+VALUES(bid_deliver),
 profit=profit+VALUES(profit),
 click=click+VALUES(click)
`

func flush(supDemSrc map[string]*datamodels.TableModel, supSrc map[string]*datamodels.TableModel) error {

	if len(supDemSrc) == 0 && len(supSrc) == 0 {
		return nil
	}

	var (
		parts1, parts2 []string
	)

	for i := range supDemSrc {
		if has, part := hasDataSupplierDemand(supDemSrc[i]); has {
			parts1 = append(parts1, part)
		}
	}

	q1 := fmt.Sprintf(supDemSrcTable, strings.Join(parts1, ",\n"))

	for i := range supSrc {
		if has, part := hasDataSupplier(supSrc[i]); has {
			parts2 = append(parts2, part)
		}
	}

	q2 := fmt.Sprintf(supSrcTable, strings.Join(parts2, ",\n"))

	if len(parts1)+len(parts2) == 0 {
		return nil
	}

	if len(parts1) > 0 && len(parts2) == 0 {
		return models.NewManager().MultiQuery(models.Parts{
			Query: q1,
		})
	}
	if len(parts1) == 0 && len(parts2) > 0 {
		return models.NewManager().MultiQuery(models.Parts{
			Query: q2,
		})
	}
	return models.NewManager().MultiQuery(
		models.Parts{
			Query: q1,
		},
		models.Parts{
			Query: q2,
		},
	)

}
