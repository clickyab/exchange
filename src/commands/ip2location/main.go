package main

import (
	"commands"
	"services/config"
	"services/initializer"

	"net/http"

	"strings"

	"services/ip2location"

	"encoding/json"

	"services/assert"

	"github.com/Sirupsen/logrus"
)

func main() {
	config.Initialize(commands.Organization, commands.AppName, commands.Prefix)
	defer initializer.Initialize()()
	ip2location.Open()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Trim(r.URL.String(), "/ ")
		tmp := ip2location.Get_all(ip)
		dec := json.NewEncoder(w)
		assert.Nil(dec.Encode(tmp))
	})
	go func() {
		// TODO : load from config
		http.ListenAndServe(":8190", nil)
	}()

	sig := commands.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}