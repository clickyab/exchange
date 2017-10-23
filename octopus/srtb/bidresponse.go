package srtb

// bidResponse is the object that exchange return to publisher
type bidResponse struct {
	ID   string `json:"id"`
	Bids []Bid  `json:"bids"`
}
