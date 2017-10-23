package srtb

// Impression each bid request contains some impression which has data about requested ad
type Impression struct {
	ID       string  `json:"id"`
	lid      string  `json:"-"`
	Banner   *Banner `json:"Banner"`
	BidFloor float64 `json:"bid_floor"`
	Secure   int     `json:"secure"`
}
