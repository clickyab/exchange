package renderer

import (
	"net/http"

	"encoding/json"

	"clickyab.com/exchange/octopus/exchange"
)

type rtb struct {
	ID   string `json:"id"`
	Bids []bid  `json:"bids"`
}

type bid struct {
	ID       string `json:"id"`
	ImpID    string `json:"impid"`
	Price    int64  `json:"price"`
	WinURL   string `json:"nurl"`
	AdMarkup string `json:"adm"`
	Width    int    `json:"w"`
	Height   int    `json:"h"`
}

func (r rtb) Render(in exchange.BidResponse, w http.ResponseWriter) error {
	response := rtb{ID: in.ID()}
	for _, i := range in.Bids() {
		response.Bids = append(response.Bids, bid{
			ID:       i.ID(),
			WinURL:   i.WinURL(),
			Width:    i.AdWidth(),
			Height:   i.AdHeight(),
			ImpID:    i.ImpID(),
			Price:    i.Price(),
			AdMarkup: i.AdMarkup(),
		})
	}

	data, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	return err
}

// NewRestRenderer returns renderer response for restful
func NewRestRenderer() exchange.Renderer {
	return rtb{}
}
