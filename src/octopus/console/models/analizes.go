package models

// SupSrcDem supplier_source_demand
type SupSrcDem struct {
	Supplier   int64  `json:"supplier" db:"supplier"`
	Demand     int64  `json:"demand" db:"demand"`
	Source     string `json:"source" db:"source"`
	Time       int    `json:"time" db:"time"`
	Request    int    `json:"request" db:"request"`
	Impression int    `json:"impression" db:"impression"`
	Show       int    `json:"show" db:"show"`
	ImpBid     int    `json:"imp_bid" db:"imp_bid"`
	ShowBid    int    `json:"show_bid" db:"show_bid"`
	Win        int    `json:"win" db:"win"`
}

// SupSrc supplier_source
type SupSrc struct {
	Supplier   int64  `json:"supplier" db:"supplier"`
	Source     string `json:"source" db:"source"`
	Time       int    `json:"time" db:"time"`
	Request    int    `json:"request" db:"request"`
	Impression int    `json:"impression" db:"impression"`
	Show       int    `json:"show" db:"show"`
	ImpBid     int    `json:"imp_bid" db:"imp_bid"`
	ShowBid    int    `json:"show_bid" db:"show_bid"`
}

// DemSrc demand_source
type DemSrc struct {
	Demand     int64  `json:"demand" db:"demand"`
	Source     string `json:"source" db:"source"`
	Time       int    `json:"time" db:"time"`
	Request    int    `json:"request" db:"request"`
	Impression int    `json:"impression" db:"impression"`
	Show       int    `json:"show" db:"show"`
	ImpBid     int    `json:"imp_bid" db:"imp_bid"`
	ShowBid    int    `json:"show_bid" db:"show_bid"`
}
