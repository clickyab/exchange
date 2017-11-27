package models

import (
	"fmt"
	"time"

	"clickyab.com/exchange/octopus/console/user/aaa"
	"github.com/clickyab/services/assert"
)

// updateSupplierReport will update supplier report (inclusive)
func (m *Manager) updateSupplierReport(t time.Time) {
	td := t.Format("2006-01-02")
	from, to := factTableRange(t)
	var q = fmt.Sprintf(`INSERT INTO %s (
								supplier,
								target_date,
								ad_in,
								ad_out,
								ad_deliver,
								earn,
								click
								)
							SELECT supplier,
							"%s",
							sum(ad_in),
							sum(ad_out),
							sum(ad_deliver),
							sum(bid_deliver),
							sum(click)
								FROM %s WHERE time_id BETWEEN %d AND %d
							GROUP BY supplier
							 ON DUPLICATE KEY UPDATE
							  supplier=VALUES(supplier),
							  target_date=VALUES(target_date),
							  ad_in=VALUES(ad_in),
							  ad_out=VALUES(ad_out),
							  ad_deliver=VALUES(ad_deliver),
							  earn=VALUES(earn),
							  click=VALUES(click)`, SupplierReportTableName, td, SupplierTableName, from, to)

	_, err := NewManager().GetWDbMap().Exec(q)
	assert.Nil(err)
}

// UpdateSupplierRange will update supplier report in range of two date (inclusive)
func (m *Manager) UpdateSupplierRange(from time.Time, to time.Time) {
	if from.Unix() > to.Unix() {
		from, to = to, from
	}
	to = to.Add(24 * time.Hour)
	for from.Unix() < to.Unix() {
		m.updateSupplierReport(from)
		from = from.Add(time.Hour * 24)
	}
}

// FillSupplierReport supplier report
func (m *Manager) FillSupplierReport(p, c int, sort, order string, from, to string, user *aaa.User) ([]SupplierReporter, int64) {
	var res []SupplierReporter
	var params []interface{}
	limit := c
	offset := (p - 1) * c
	params = append(params, from, to)
	countQuery := fmt.Sprintf("SELECT COUNT(sr.id) FROM %s AS sr "+
		"INNER JOIN %s AS s ON s.name=sr.supplier WHERE sr.target_date BETWEEN ? AND ? ", SupplierReportTableName, "suppliers")
	query := fmt.Sprintf("SELECT sr.*,"+
		"CASE WHEN ad_out=0 THEN 0 ELSE ROUND(ad_deliver/ad_out,2) END AS deliver_rate,"+
		"CASE WHEN ad_in=0 THEN 0 ELSE ROUND(ad_out/ad_in,2) END AS success_rate FROM %s AS sr "+
		"INNER JOIN %s AS s ON s.name=sr.supplier WHERE sr.target_date BETWEEN ? AND ? ", SupplierReportTableName, "suppliers")
	//check user perm
	if user.UserType != aaa.AdminUserType {
		countQuery += "AND s.user_id = ? "
		query += "AND s.user_id = ? "
		params = append(params, user.ID)
	}
	if sort != "" {
		query += fmt.Sprintf("ORDER BY %s %s ", sort, order)
	}
	query += fmt.Sprintf("LIMIT %d OFFSET %d ", limit, offset)
	count, err := m.GetRDbMap().SelectInt(countQuery, params...)
	assert.Nil(err)

	_, err = m.GetRDbMap().Select(&res, query, params...)
	assert.Nil(err)
	return res, count
}
