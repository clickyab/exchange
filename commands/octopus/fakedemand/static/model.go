package static

import "clickyab.com/exchange/octopus/exchange"

type request struct {
	IP          string `json:"ip"`
	Scheme      string `json:"scheme,omitempty"`
	PageTrackID string `json:"page_track_id"`
	UserTrackID string `json:"user_track_id"`
	Publisher   struct {
		PubName         string `json:"name"`
		PubFloorCPM     int64  `json:"floor_cpm"`
		PubSoftFloorCPM int64  `json:"soft_floor_cpm"`

		//sup   exchange.Supplier
		//rates []exchange.Rate
	} `json:"source"`
	Categories []exchange.Category `json:"categories"`
	Type       string              `json:"type"`
	UnderFloor bool                `json:"under_floor"`
	App        struct {
		OSVersion  string `json:"os_version,omitempty"`
		Operator   string `json:"operator,omitempty"`
		Brand      string `json:"brand,omitempty"`
		Model      string `json:"model,omitempty"`
		Language   string `json:"language,omitempty"`
		Network    string `json:"network,omitempty"`
		OSIdentity string `json:"os_identity,omitempty"`
		MCC        int64  `json:"mcc,omitempty"`
		MNC        int64  `json:"mnc,omitempty"`
		LAC        int64  `json:"lac,omitempty"`
		CID        int64  `json:"cid,omitempty"`
		UserAgent  string `json:"user_agent,omitempty"`
	} `json:"app,omitempty"`
	Web struct {
		Referrer  string `json:"referrer,omitempty"`
		Parent    string `json:"parent,omitempty"`
		UserAgent string `json:"user_agent,omitempty"`
	} `json:"web,omitempty"`
	Vast struct {
		Referrer  string `json:"referrer,omitempty"`
		Parent    string `json:"parent,omitempty"`
		UserAgent string `json:"user_agent,omitempty"`
	} `json:"vast,omitempty"`

	Slots []Slot `json:"slots"`
	Host  string
}

// Slot of request
type Slot struct {
	W           int               `json:"width"`
	H           int               `json:"height"`
	TID         string            `json:"track_id"`
	FallbackURL string            `json:"fallback_url"`
	FAttribute  map[string]string `json:"attributes"`
}

type response struct {
	TrackID     string `json:"track_id"`
	AdTrackID   string `json:"ad_track_id"`
	Winner      int64  `json:"winner"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Code        string `json:"code"`
	IsFilled    bool   `json:"is_filled"`
	Landing     string `json:"landing"`
	Description string `json:"description,omitempty"`
}
