package models

import (
	"time"

	"github.com/clickyab/services/assert"
)

const (
	// ExchangeReportTableName table name exchange
	ExchangeReportTableName = "exchange_report"
	// DemandTableName table name demand
	DemandTableName = "sup_dem_src"
	// SupplierTableName table supplier
	SupplierTableName = "sup_src"
	// DemandReportTableName demand report table
	DemandReportTableName = "demand_report"
	// SupplierReportTableName table supplier report
	SupplierReportTableName = "supplier_report"
)

// SupplierSourceDemand supplier_source_demand
type SupplierSourceDemand struct {
	ID         int64   `json:"id" db:"id"`
	Demand     string  `json:"demand" db:"demand"`
	Supplier   string  `json:"supplier" db:"supplier"`
	Source     string  `json:"source" db:"source"`
	TimeID     int64   `json:"time_id" db:"time_id"`
	RequestOut int64   `json:"request_out" db:"request_out"`
	AdIn       int64   `json:"ad_in" db:"ad_in"`
	AdOut      int64   `json:"ad_out" db:"ad_out"`
	BidWin     float64 `json:"bid_win" db:"bid_win"`
	AdWin      int64   `json:"ad_win" db:"ad_win"`
	AdDeliver  int64   `json:"ad_deliver" db:"ad_deliver"`
	BidDeliver int64   `json:"bid_deliver" db:"bid_deliver"`
	Profit     int64   `json:"profit" db:"profit"`
	Click      int64   `json:"click" db:"click"`
}

// SupplierSource supplier_source
type SupplierSource struct {
	ID         int64   `json:"id" db:"id"`
	Supplier   string  `json:"supplier" db:"supplier"`
	Source     string  `json:"source" db:"source"`
	TimeID     int64   `json:"time_id" db:"time_id"`
	AdIn       int64   `json:"ad_in" db:"ad_in"`
	AdOut      int64   `json:"ad_out" db:"ad_out"`
	AdWin      int64   `json:"ad_win" db:"ad_win"`
	AdDeliver  int64   `json:"ad_deliver" db:"ad_deliver"`
	BidDeliver float64 `json:"bid_deliver" db:"bid_deliver"`
	Profit     int64   `json:"profit" db:"profit"`
	Click      int64   `json:"click" db:"click"`
}

// TimeTable TimeTable
type TimeTable struct {
	ID     int64 `json:"id" db:"id"`
	Year   int64 `json:"year" db:"year"`
	Month  int64 `json:"month" db:"month"`
	Day    int64 `json:"day" db:"day"`
	Hour   int64 `json:"hour" db:"hour"`
	JYear  int64 `json:"j_year" db:"j_year"`
	JMonth int64 `json:"j_month" db:"j_month"`
	JDay   int64 `json:"j_day" db:"j_day"`
}

// ExchangeReport exchange_report
type ExchangeReport struct {
	ID            int64     `json:"id" db:"id"`
	Date          time.Time `json:"target_date" db:"target_date"`
	SupplierAdIN  int64     `json:"supplier_ad_in" db:"supplier_ad_in"`
	SupplierAdOUT int64     `json:"supplier_ad_out" db:"supplier_ad_out"`
	DemandAdIN    int64     `json:"demand_ad_in" db:"demand_ad_in"`
	DemandAdOUT   int64     `json:"demand_ad_out" db:"demand_ad_out"`
	Earn          float64   `json:"earn" db:"earn"`
	Spent         float64   `json:"spent" db:"spent"`
	Income        float64   `json:"income" db:"income"`
	Click         int64     `json:"click" db:"click"`
}

// Parts is a multi query trick
type Parts struct {
	Query  string
	Params []interface{}
}

// DemandReport demand_report
type DemandReport struct {
	ID          int64     `json:"id" db:"id"`
	Demand      string    `json:"demand" db:"demand"`
	TargetDate  time.Time `json:"target_date" db:"target_date"`
	AdOut       int64     `json:"ad_out" db:"ad_out"`
	AdIn        int64     `json:"ad_in" db:"ad_in"`
	AdWin       int64     `json:"ad_win" db:"ad_win"`
	AdDeliver   int64     `json:"ad_deliver" db:"ad_deliver"`
	BidDeliver  int64     `json:"bid_deliver" db:"bid_deliver"`
	Profit      int64     `json:"profit" db:"profit"`
	SuccessRate float64   `json:"success_rate" db:"success_rate"`
	DeliverRate float64   `json:"deliver_rate" db:"deliver_rate"`
	WinRate     float64   `json:"win_rate" db:"win_rate"`
	Click       int64     `json:"click" db:"click"`
}

// SupplierReporter table
type SupplierReporter struct {
	ID          int64     `json:"id" db:"id"`
	Supplier    string    `json:"supplier" db:"supplier"`
	Date        time.Time `json:"target_date" db:"target_date"`
	AdIn        int64     `json:"ad_in" db:"ad_in"`
	AdOut       int64     `json:"ad_out" db:"ad_out"`
	AdDeliver   int64     `json:"ad_deliver" db:"ad_deliver"`
	Earn        int64     `json:"earn" db:"earn"`
	SuccessRate float64   `json:"success_rate" db:"success_rate"`
	DeliverRate float64   `json:"deliver_rate" db:"deliver_rate"`
	Click       int64     `json:"click" db:"click"`
}

// MultiQuery is a hack to run multiple query in one transaction
func (m *Manager) MultiQuery(parts ...Parts) (err error) {
	err = m.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			err = m.Commit()
		}
	}()

	for _, q := range parts {
		_, err = m.GetProperDBMap().Exec(q.Query, q.Params...)
		if err != nil {
			return
		}
	}

	return nil
}
