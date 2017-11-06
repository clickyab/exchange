package main

import (
	"clickyab.com/exchange/commands"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/shell"
	"github.com/sirupsen/logrus"
)

var exchangeURL = config.GetStringDefault("supplier.exchange.url", "http://exchange.dev/api/rest/get")
var prefix = "commands/octopus/rtb-supplier/static/template"

func main() {
	config.Initialize(commands.Organization, commands.AppName, commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()
	sig := shell.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
