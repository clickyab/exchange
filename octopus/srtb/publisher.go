package srtb

// Inventory is the publisher type which contains app and site
type Publisher struct {
	ID     string   `json:"id"`
	Name   string   `json:"name,omitempty"`
	Domain string   `json:"domain,omitempty"`
	Ref    string   `json:"ref,omitempty"`
	Cat    []string `json:"cat"`
}

// App is the app type of publisher
type App struct {
	Publisher
	Bundle   string `json:"bundle,omitempty"`
	StoreURL string `json:"storeurl,omitempty"`
}

// Site is the site type publisher
type Site struct {
	Publisher
	Page string `json:"page,omitempty"`
	Ref  string `json:"ref,omitempty"`
}
