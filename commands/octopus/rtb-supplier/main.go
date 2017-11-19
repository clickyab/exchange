package main

import (
	"clickyab.com/exchange/commands"
	"github.com/clickyab/services/config"
	_ "github.com/clickyab/services/fluentd"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/shell"
	"github.com/sirupsen/logrus"
)

var exchangeURL = config.RegisterString("supplier.exchange.url", "http://exchange.dev/api/rest/get", "gi")
var prefix = "commands/octopus/rtb-supplier/static/template"

func main() {
	config.Initialize(commands.Organization, commands.AppName, commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()
	sig := shell.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
