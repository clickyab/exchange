package srtb

import "time"

// BidRequest is what is accepted from each publisher
type BidRequest struct {
	ID     string        `json:"id"`
	Imp    []Impression  `json:"imp"`
	Site   *Site         `json:"site,omitempty"`
	App    *App          `json:"app,omitempty"`
	Device Device        `json:"device"`
	UserID string        `json:"user_id,omitempty"`
	Test   bool          `json:"test"`
	TMax   time.Duration `json:"tmax,omitempty"`
	WLang  []string      `json:"wlang,omitempty"`
	BCat   []string      `json:"bcat,omitempty"`
	BAdv   []string      `json:"badv,omitempty"`
}
