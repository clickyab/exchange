package main

import (
	"net/http"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/shell"
)

var (
	port = config.RegisterString("test.config", ":3501", "desc")
)

func main() {
	// TODO : use framework
	config.Initialize("test", "test", "test")
	http.HandleFunc("/start", getSupplierDemo)
	http.HandleFunc("/send", postSupplierDemo)

	assert.Nil(http.ListenAndServe(port.String(), nil))
	shell.WaitExitSignal()
}
