package srtb

import "time"

// BidRequest is what is accepted from each publisher
type BidRequest struct {
	ID     string        `json:"id"`
	Imp    []Impression  `json:"Imp"`
	Site   *Site         `json:"Site,omitempty"`
	App    *App          `json:"App,omitempty"`
	Device Device        `json:"Device"`
	UserID string        `json:"User"`
	Test   bool          `json:"test"`
	TMax   time.Duration `json:"tmax,omitempty"`
	WLang  []string      `json:"wlang,omitempty"`
	BCat   []string      `json:"bcat,omitempty"`
	BAdv   []string      `json:"badv,omitempty"`
}
