package models

import (
	"fmt"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
)

// Manager is the model manager
type Manager struct {
	mysql.Manager
}

// Initialize the manager. nothing to do, just keep it in interface shape
func (m *Manager) Initialize() {
	m.AddTableWithName(
		SupplierSourceDemand{},
		"sup_dem_src",
	).SetKeys(
		true,
		"ID",
	)
	m.AddTableWithName(
		SupplierSource{},
		"sup_src",
	).SetKeys(
		true,
		"ID",
	)
	m.AddTableWithName(
		SupplierSource{},
		"exchange_report",
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(DemandReport{}, "demand_report").
		SetKeys(true, "ID")

	m.AddTableWithName(
		SupplierReporter{},
		"supplier_report",
	).SetKeys(
		true,
		"ID")

}

// UpdateDemandReport will update demand report in range of two date (inclusive)
func (m *Manager) UpdateDemandReport(t time.Time) {
	td := t.Format("2006-01-02")
	from, to := factTableRange(t)

	var q = fmt.Sprintf(`INSERT INTO demand_report (
								demand,
								target_date,
								request_out_count,
								ad_in_count,
								imp_out_count,
								ad_out_count,
								ad_out_bid,
								deliver_count,
								deliver_bid,
								profit
								)

							SELECT demand,
							"%s",
							sum(request_out_count),
							sum(ad_in_count),
							sum(imp_out_count),
							sum(ad_out_count),
							sum(ad_out_bid),
							sum(deliver_count),
							sum(deliver_bid),
							sum(profit)
								FROM sup_dem_src WHERE time_id BETWEEN %d AND %d
							GROUP BY demand

							 ON DUPLICATE KEY UPDATE
							  demand=VALUES(demand),
							  target_date=VALUES(target_date),
							  request_out_count=VALUES(request_out_count),
							  ad_in_count=VALUES(ad_in_count),
							  imp_out_count=VALUES(imp_out_count),
							  ad_out_count=VALUES(ad_out_count),
							  ad_out_bid=VALUES(ad_out_bid),
							  deliver_count=VALUES(deliver_count),
							  deliver_bid=VALUES(deliver_bid),
							  profit=values(profit)`, td, from, to)

	_, err := NewManager().GetRDbMap().Exec(q)
	assert.Nil(err)
}

// UpdateDemandRange will update demand report in range of two date (inclusive)
func (m *Manager) UpdateDemandRange(from time.Time, to time.Time) {
	if from.Unix() > to.Unix() {
		from, to = to, from
	}
	to = to.Add(24 * time.Hour)
	for from.Unix() < to.Unix() {
		m.UpdateDemandReport(from)
		from = from.Add(time.Hour * 24)
	}
}

// NewManager return a new manager object
func NewManager() *Manager {
	return &Manager{}
}

func init() {
	mysql.Register(NewManager())
}
