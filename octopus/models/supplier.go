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
								impression_in_count,
								ad_out_count,
								delivered_count,
								earn
								)
							SELECT supplier,
							"%s",
							sum(imp_in_count),
							sum(ad_out_count),
							sum(deliver_count),
							sum(deliver_bid)
								FROM %s WHERE time_id BETWEEN %d AND %d
							GROUP BY supplier
							 ON DUPLICATE KEY UPDATE
							  supplier=VALUES(supplier),
							  target_date=VALUES(target_date),
							  impression_in_count=VALUES(impression_in_count),
							  ad_out_count=VALUES(ad_out_count),
							  delivered_count=VALUES(delivered_count),
							  earn=VALUES(earn)`, SupplierReportTableName, td, SupplierTableName, from, to)

	_, err := NewManager().GetRDbMap().Exec(q)
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
func (m *Manager) FillSupplierReport(p, c int, sort, order string, from, to int64, user *aaa.User) ([]SupplierReporter, int64) {
	var res []SupplierReporter
	var params []interface{}
	limit := c
	offset := (p - 1) * c
	params = append(params, from, to)
	countQuery := fmt.Sprintf("SELECT COUNT(sr.id) FROM %s AS sr "+
		"INNER JOIN %s AS s ON s.name=sr.supplier WHERE sr.target_date BETWEEN ? AND ? ", SupplierReportTableName, "suppliers")
	query := fmt.Sprintf("SELECT sr.* FROM %s AS sr "+
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
