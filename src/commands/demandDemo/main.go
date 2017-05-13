package demandDemo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"services/random"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3412", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	temp := &payload{}
	err := decoder.Decode(temp)
	if err != nil {
		w.Write([]byte("u suck!!!"))
	}

	var res response
	for i := range temp.Slots {
		a := singleResponse{
			ID:      <-random.ID,
			MaxCPM:  temp.Source.FloorCPM + 1,
			Width:   temp.Slots[i].Width,
			Height:  temp.Slots[i].Height,
			URL:     fmt.Sprintf("http://a.clickyab.com/ads/?a=4471405272967&width=%s&height=%s&slot=71634138754&domainname=p30download.com&eventpage=416310534&loc=http%3A%2F%2Fp30download.com%2Fagahi%2Fplan%2Fa1i.php&ref=http%3A%2F%2Fp30download.com%2F&adcount=1", temp.Slots[i].Width, temp.Slots[i].Height),
			Landing: "clickyab.com",
		}
		res = append(res, a)
	}

	w.Write([]byte(res))
}
