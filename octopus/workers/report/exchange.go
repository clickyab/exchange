package report

import (
	"time"

	"fmt"

	"clickyab.com/exchange/octopus/models"
	"clickyab.com/exchange/services/assert"
)

// FetchDemand select demand side
func FetchDemand(start int64, end int64) *models.ExchangeReport {
	ex := models.ExchangeReport{}
	q := fmt.Sprintf(`SELECT SUM(imp_in_count) AS demand_impression_in,
	SUM(imp_out_count) AS demand_impression_out,
	SUM(deliver_bid) AS earn
	FROM %s
	WHERE time_id >= ?
	AND time_id <= ?`, models.SupplierDemandTableName)
	m := models.NewManager()
	_, err := m.GetRDbMap().Select(ex, q, start, end)
	assert.Nil(err)
	return &ex
}

// FetchSupplier select demand side
func FetchSupplier(start int64, end int64) *models.ExchangeReport {
	ex := models.ExchangeReport{}
	q := fmt.Sprintf(`SELECT SUM(request_in_count) AS supplier_impression_in,
	SUM(deliver_count) AS supplier_impression_out,
	SUM(deliver_bid) AS spent
	FROM %s
	WHERE time_id >= ?
	AND time_id <= ?`, models.SupplierTableName)
	m := models.NewManager()
	_, err := m.GetRDbMap().Select(ex, q, start, end)
	assert.Nil(err)
	return &ex
}

func updateExchangeReport(t time.Time) {
	start, end := factTableYesterdayID(t)
	dem := FetchDemand(start, end)
	sup := FetchSupplier(start, end)
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
				`, models.ExchangeReportTableName)
	m := models.NewManager()
	_, err := m.GetRDbMap().Exec(q, t, sup.SupplierImpressionIN,
		sup.SupplierImpressionOUT, dem.DemandImpressionIN, dem.DemandImpressionOUT,
		sup.Earn, dem.Spent, sup.Earn-sup.Spent)
	assert.Nil(err)
}

// UpdateExchangeRange cron worker report exchange
func UpdateExchangeRange(from, to time.Time) {
	if from.Unix() > to.Unix() {
		from, to = to, from
	}
	to = to.Add(24 * time.Hour)
	for from.Unix() < to.Unix() {
		updateExchangeReport(from)
		from = from.Add(time.Hour * 24)
	}
}
