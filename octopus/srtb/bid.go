package srtb

// Bid part of every bidResponse
type Bid struct {
	ID       string `json:"id"`
	ImpID    string `json:"imp_id"`
	Price    int64  `json:"price"`
	AdMarkup string `json:"adm"`
	Width    int    `json:"w"`
	Height   int    `json:"h"`
}
