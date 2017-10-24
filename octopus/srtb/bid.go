package srtb

// Bid part of every bidrequest
type Bid struct {
	ID       string `json:"id"`
	ImpID    string `json:"imp_id"`
	Price    int64  `json:"price"`
	WinURL   string `json:"nurl"`
	AdMarkup string `json:"adm"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}
