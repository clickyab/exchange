package demandDemo

type payload struct {
	TrackID   string `json:"track_id"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`

	Source struct {
		Name         string                 `json:"name"`
		FloorCPM     int                    `json:"floor_cpm"`
		SoftFloorCPM int                    `json:"soft_floor_cpm"`
		Attributes   map[string]interface{} `json:"attributes"`
	} `json:"source"`

	Location struct {
		Country struct {
			Valid bool   `json:"valid"`
			Name  string `json:"name"`
			ISO   string `json:"iso"`
		} `json:"country"`
		Province struct {
			Valid bool   `json:"valid"`
			Name  string `json:"name"`
		} `json:"province"`
		LatLon struct {
			Valid bool    `json:"valid"`
			Lat   float64 `json:"lat"`
			long  float64 `json:"long"`
		} `json:"latlon"`
	} `json:"location"`

	Attributes map[string]interface{} `json:"attributes"`
	Slots      []struct {
		Width   int    `json:"width"`
		Height  int    `json:"height"`
		TrackID string `json:"track_id"`
	} `json:"slots"`

	Category []struct {
		Name string `json:"name"`
	} `json:"category"`

	Platform   string `json:"platform"`
	Underfloor bool   `json:"underfloor"`
}

type singleResponse struct {
	ID      string `json:"id"`
	MaxCPM  int    `json:"max_cpm"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	URL     string `json:"url"`
	Landing string `json:"landing"`
}

type response []singleResponse
