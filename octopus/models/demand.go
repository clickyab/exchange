package models

import (
	"fmt"
	"time"

	"clickyab.com/exchange/octopus/console/user/aaa"
	"github.com/clickyab/services/assert"
)

func calculator(a []DemandReport) []DemandReport {
	res := make([]DemandReport, 0)

	for _, v := range a {
		res = append(res, DemandReport{
			Demand:          v.Demand,
			ID:              v.ID,
			ImpOutCount:     v.ImpOutCount,
			RequestOutCount: v.AdInCount,
			SuccessRate:     (v.ImpOutCount * 100) / v.AdInCount,
			DeliverCount:    v.DeliverCount,
			DeliverRate:     (v.DeliverCount * 100) / v.AdOutCount,
			AdOutCount:      v.AdOutCount,
			WinRate:         (v.AdOutCount * 100) / v.AdInCount,
			DeliverBid:      v.DeliverBid,
		})
	}

	return res
}

// DemandByDate returns list of demand for specific date
func (m *Manager) DemandByDate(t time.Time) []DemandReport {
	return m.DemandByDateRange(t, t)
}

// DemandByDateRange returns list of demand for range of dates
func (m *Manager) DemandByDateRange(from time.Time, to time.Time) []DemandReport {
	return m.DeamandByDateRangeNames(from, to)
}

// DemandByDateNames returns demand with specific date
func (m *Manager) DemandByDateNames(f time.Time, demands ...string) []DemandReport {
	return m.DeamandByDateRangeNames(f, f, demands...)
}

// DeamandByDateRangeNames returns demands with for range of dates
func (m *Manager) DeamandByDateRangeNames(f time.Time, t time.Time, names ...string) []DemandReport {

	var a []DemandReport

	q := fmt.Sprintf(`SELECT
					id,
					demand,
					target_date,
					request_out_count,
					ad_in_count,
					imp_out_count,
					ad_out_count,
					ad_out_bid,
					deliver_count,
					deliver_bid
				FROM demand_report where %s %s ORDER BY id DESC	`,
		demandtimePartial(true, f, t), demandPartial(false, names...))

	_, err := NewManager().GetRDbMap().Select(&a, q)
	assert.Nil(err)

	return calculator(a)
}

// DemandAggregateByDate returns list of demand for specific date
func (m *Manager) DemandAggregateByDate(t time.Time) []DemandReport {
	return m.DemandAggregateDemandsByDateRange(t, t)
}

// DemandAggregateByDateRange return list of demand for range of dates
func (m *Manager) DemandAggregateByDateRange(f time.Time, t time.Time) []DemandReport {
	return m.DemandAggregateDemandsByDateRange(f, t)

}

// DemandAggregateDemandsByDate return demand with specific date
func (m *Manager) DemandAggregateDemandsByDate(f time.Time, demands ...string) []DemandReport {
	return m.DemandAggregateDemandsByDateRange(f, f, demands...)
}

// DemandAggregateDemandsByDateRange return demands with for range of dates
func (m *Manager) DemandAggregateDemandsByDateRange(f time.Time, t time.Time, demands ...string) []DemandReport {

	var a []DemandReport

	q := fmt.Sprintf(`SELECT
					demand,
					target_date,
					SUM(request_out_count) as request_out_count ,
					SUM(ad_in_count) as ad_in_count,
					SUM(imp_out_count) as imp_out_count,
					SUM(ad_out_count) as win_count,
					SUM(ad_out_bid) as win_bid,
					SUM(deliver_count) as deliver_count,
					SUM(deliver_bid) as deliver_bid
				FROM demand_report where %s %s GROUP BY demand`,
		demandtimePartial(true, f, t), demandPartial(false, demands...))

	_, err := NewManager().GetRDbMap().Select(&a, q)
	assert.Nil(err)

	return calculator(a)
}

// DemandAggregateAllByDate return all with for range of dates
func (m *Manager) DemandAggregateAllByDate(t time.Time) []DemandReport {
	return m.DemandAggregateAllByDateRange(t, t)
}

// DemandAggregateAllByDateRange return demands with for range of dates
func (m *Manager) DemandAggregateAllByDateRange(f time.Time, t time.Time) []DemandReport {

	var a []DemandReport

	q := fmt.Sprintf(`SELECT
					"All",
					target_date,
					SUM(request_out_count) as request_out_count ,
					SUM(ad_in_count) as ad_in_count,
					SUM(imp_out_count) as imp_out_count,
					SUM(ad_out_count) as ad_out_count,
					SUM(ad_out_bid) as ad_out_bid,
					SUM(deliver_count) as deliver_count,
					SUM(deliver_bid) as deliver_bid
				FROM demand_report where %s`,
		demandtimePartial(true, f, t))

	_, err := NewManager().GetRDbMap().Select(&a, q)
	assert.Nil(err)

	return calculator(a)
}

func demandPartial(isFirst bool, names ...string) (res string) {
	if len(names) == 0 {
		return
	}
	if isFirst {
		res = " demand = "
	} else {
		res = "AND demand = "
	}

	for i := range names {
		res += fmt.Sprintf(`"%s"`, names[i])
		if len(names) < i+1 {
			res += " OR "
		}
	}
	return
}

func demandtimePartial(isFirst bool, from time.Time, to time.Time) (res string) {
	if isFirst {
		res = "target_date  "
	} else {
		res = " AND target_date  "
	}
	if from.Unix() > to.Unix() {
		from, to = to, from
	}
	f, e := from.Format("2006-01-02"), to.Format("2006-01-02")
	if f == e {
		res += fmt.Sprintf(` = "%s"`, f)
	} else {
		res += fmt.Sprintf(` BETWEEN "%s" AND "%s"`, f, e)
	}
	return
}

// FillDemandReport demand report
func (m *Manager) FillDemandReport(p, c int, sort, order string, from, to int64, user *aaa.User) ([]DemandReport, int64) {
	var res []DemandReport
	var params []interface{}
	limit := c
	offset := (p - 1) * c
	params = append(params, from, to)
	countQuery := fmt.Sprintf("SELECT COUNT(dr.id) FROM %s AS dr "+
		"INNER JOIN %s AS d ON d.name=dr.demand WHERE dr.target_date BETWEEN ? AND ? ", DemandReportTableName, "demands")
	query := fmt.Sprintf("SELECT dr.* FROM %s AS dr "+
		"INNER JOIN %s AS d ON d.name=dr.demand WHERE dr.target_date BETWEEN ? AND ? ", DemandReportTableName, "demands")
	//check user perm
	if user.UserType != aaa.AdminUserType {
		countQuery += "AND d.user_id = ? "
		query += "AND d.user_id = ? "
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
