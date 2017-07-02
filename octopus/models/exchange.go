package models

import (
	"time"

	"fmt"

	"clickyab.com/exchange/octopus/console/user/aaa"
	"github.com/clickyab/services/assert"
)

// fetchDemand select demand side
func (m *Manager) fetchDemand(start int64, end int64) *ExchangeReport {
	ex := ExchangeReport{}
	q := fmt.Sprintf(`SELECT COALESCE(SUM(ad_out_count),0) AS demand_impression_in,
	COALESCE(SUM(imp_out_count),0) AS demand_impression_out,
	COALESCE(SUM(deliver_bid),0) AS earn
	FROM %s
	WHERE time_id >= ?
	AND time_id <= ?`, DemandTableName)
	err := m.GetRDbMap().SelectOne(&ex, q, start, end)
	assert.Nil(err)
	return &ex
}

// fetchSupplier select demand side
func (m *Manager) fetchSupplier(start int64, end int64) *ExchangeReport {
	ex := ExchangeReport{}
	q := fmt.Sprintf(`SELECT COALESCE(SUM(request_in_count),0) AS supplier_impression_in,
	COALESCE(SUM(deliver_count),0) AS supplier_impression_out,
	COALESCE(SUM(deliver_bid),0) AS spent
	FROM %s
	WHERE time_id >= ?
	AND time_id <= ?`, SupplierTableName)
	err := m.GetRDbMap().SelectOne(&ex, q, start, end)
	assert.Nil(err)
	return &ex
}

// updateExchangeReport will update demand report (inclusive)
func (m *Manager) updateExchangeReport(t time.Time) {
	td := t.Format("2006-01-02")
	from, to := factTableRange(t)
	dem := m.fetchDemand(from, to)
	sup := m.fetchSupplier(from, to)
	q := fmt.Sprintf(`INSERT INTO %s
				(target_date,
				supplier_impression_in,
				supplier_impression_out,
				demand_impression_in,
				demand_impression_out,
				earn,
				spent,
				income)
				VALUES(?,?,?,?,?,?,?,?)
				ON DUPLICATE KEY UPDATE
				supplier_impression_in = VALUES(supplier_impression_in),
				supplier_impression_out = VALUES(supplier_impression_out),
				demand_impression_in = VALUES(demand_impression_in),
				demand_impression_out = VALUES(demand_impression_out),
				earn = VALUES(earn),
				spent = VALUES(spent),
				income = VALUES(income)
				`, ExchangeReportTableName)
	_, err := m.GetRDbMap().Exec(q, td, sup.SupplierImpressionIN,
		sup.SupplierImpressionOUT, dem.DemandImpressionIN, dem.DemandImpressionOUT,
		sup.Spent, dem.Earn, dem.Earn-sup.Spent)
	assert.Nil(err)
}

// UpdateExchangeRange cron worker report exchange
func (m *Manager) UpdateExchangeRange(from time.Time, to time.Time) {
	if from.Unix() > to.Unix() {
		from, to = to, from
	}
	to = to.Add(24 * time.Hour)
	for from.Unix() < to.Unix() {
		m.updateExchangeReport(from)
		from = from.Add(time.Hour * 24)
	}

}

// FillExchangeReport exchange report
func (m *Manager) FillExchangeReport(p, c int, sort, order string, from, to int64, user *aaa.User) ([]ExchangeReport, int64) {
	var res []ExchangeReport
	var params []interface{}
	limit := c
	offset := (p - 1) * c
	params = append(params, from, to)
	countQuery := fmt.Sprintf("SELECT COUNT(er.id) FROM %s AS er "+
		"WHERE er.target_date BETWEEN ? AND ? ", ExchangeReportTableName)
	query := fmt.Sprintf("SELECT er.* FROM %s AS er "+
		"WHERE er.target_date BETWEEN ? AND ? ", ExchangeReportTableName)

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
