package report

import (
	"time"

	"fmt"

	"clickyab.com/exchange/octopus/models"
	"clickyab.com/exchange/services/assert"
)

func updateSupplierReport(t time.Time) {
	td := t.Format("2006-01-02")
	from, to := factTableRange(t)
	var q = fmt.Sprintf(`INSERT INTO %s (
								supplier,
								target_date,
								impression_in,
								impression_out,
								delivered_count,
								earn
								)

							SELECT supplier,
							"%s",
							sum(imp_in_count),
							sum(imp_out_count),
							sum(deliver_count),
							sum(deliver_bid)

								FROM %s WHERE time_id BETWEEN %d AND %d
							GROUP BY supplier

							 ON DUPLICATE KEY UPDATE
							  supplier=VALUES(supplier),
							  target_date=VALUES(target_date),
							  impression_in=VALUES(impression_in),
							  impression_out=VALUES(impression_out),
							  imp_out_count=VALUES(imp_out_count),
							  delivered_count=VALUES(delivered_count),
							  earn=VALUES(earn)`, models.SupplierReportTableName, models.SupplierTableName, td, from, to)

	_, err := models.NewManager().GetRDbMap().Exec(q)
	assert.Nil(err)
}

// UpdateSupplierRange UpdateSupplierRange
func UpdateSupplierRange(from, to time.Time) {
	if from.Unix() > to.Unix() {
		from, to = to, from
	}
	to = to.Add(24 * time.Hour)
	for from.Unix() < to.Unix() {
		updateSupplierReport(from)
		from = from.Add(time.Hour * 24)
	}

}
