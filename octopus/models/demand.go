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
			Demand:      v.Demand,
			ID:          v.ID,
			AdOut:       v.AdOut,
			AdIn:        v.AdIn,
			SuccessRate: float64((v.AdIn * 100) / v.AdOut),
			AdDeliver:   v.AdDeliver,
			DeliverRate: float64((v.AdDeliver * 100) / v.AdIn),
			AdWin:       v.AdWin,
			WinRate:     float64((v.AdWin * 100) / v.AdIn),
			BidDeliver:  v.BidDeliver,
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
	return m.DemandByDateRangeNames(from, to)
}

// DemandByDateNames returns demand with specific date
func (m *Manager) DemandByDateNames(f time.Time, demands ...string) []DemandReport {
	return m.DemandByDateRangeNames(f, f, demands...)
}

// DemandByDateRangeNames returns demands with for range of dates
func (m *Manager) DemandByDateRangeNames(f time.Time, t time.Time, names ...string) []DemandReport {

	var a []DemandReport

	q := fmt.Sprintf(`SELECT
					id,
					demand,
					target_date,
					ad_out,
					ad_in,
					ad_win,
					ad_deliver,
					bid_deliver
				FROM demand_report where %s %s ORDER BY id DESC	`,
		demandTimePartial(true, f, t), demandPartial(false, names...))

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
					SUM(ad_out) as ad_out ,
					SUM(ad_in) as ad_in,
					SUM(ad_win) as ad_win,
					SUM(ad_deliver) as ad_deliver,
					SUM(bid_deliver) as bid_deliver
				FROM demand_report where %s %s GROUP BY demand`,
		demandTimePartial(true, f, t), demandPartial(false, demands...))

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
					SUM(ad_out) as ad_out ,
					SUM(ad_in) as ad_in,
					SUM(ad_win) as ad_win,
					SUM(ad_deliver) as ad_deliver,
					SUM(bid_deliver) as bid_deliver
				FROM demand_report where %s`,
		demandTimePartial(true, f, t))

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

func demandTimePartial(isFirst bool, from time.Time, to time.Time) (res string) {
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
func (m *Manager) FillDemandReport(p, c int, sort, order string, from, to string, user *aaa.User) ([]DemandReport, int64) {
	var res []DemandReport
	var params []interface{}
	limit := c
	offset := (p - 1) * c
	params = append(params, from, to)
	countQuery := fmt.Sprintf("SELECT COUNT(dr.id) FROM %s AS dr "+
		"INNER JOIN %s AS d ON d.name=dr.demand WHERE dr.target_date BETWEEN ? AND ? ", DemandReportTableName, "demands")
	query := fmt.Sprintf("SELECT dr.*,"+
		"CASE WHEN ad_in=0 THEN 0 ELSE ROUND(ad_deliver/ad_in,2) END AS deliver_rate,"+
		"CASE WHEN ad_out=0 THEN 0 ELSE ROUND(ad_in/ad_out,2) END AS success_rate,"+
		"CASE WHEN ad_in=0 THEN 0 ELSE ROUND(ad_win/ad_in,2) END AS win_rate"+
		" FROM %s AS dr "+
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

// updateDemandReport will update demand report in range of two date (inclusive)
func (m *Manager) updateDemandReport(t time.Time) {
	td := t.Format("2006-01-02")
	from, to := factTableRange(t)

	var q = fmt.Sprintf(`INSERT INTO demand_report (
								demand,
								target_date,
								ad_out,
								ad_in,
								ad_win,
								ad_deliver,
								bid_deliver,
								profit,
								click
								)

							SELECT demand,
							"%s",
							sum(ad_out),
							sum(ad_in),
							sum(ad_win),
							sum(ad_deliver),
							sum(bid_deliver),
							sum(profit),
							sum(click)
								FROM sup_dem_src WHERE time_id BETWEEN %d AND %d
							GROUP BY demand

							 ON DUPLICATE KEY UPDATE
							  demand=VALUES(demand),
							  target_date=VALUES(target_date),
							  ad_out=VALUES(ad_out),
							  ad_in=VALUES(ad_in),
							  ad_win=VALUES(ad_win),
							  ad_deliver=VALUES(ad_deliver),
							  bid_deliver=VALUES(bid_deliver),
							  profit=values(profit),
							  click=values(click)`, td, from, to)

	_, err := NewManager().GetWDbMap().Exec(q)
	assert.Nil(err)
}

// UpdateDemandRange will update demand report in range of two date (inclusive)
func (m *Manager) UpdateDemandRange(from time.Time, to time.Time) {
	if from.Unix() > to.Unix() {
		from, to = to, from
	}
	to = to.Add(24 * time.Hour)
	for from.Unix() < to.Unix() {
		m.updateDemandReport(from)
		from = from.Add(time.Hour * 24)
	}
}
