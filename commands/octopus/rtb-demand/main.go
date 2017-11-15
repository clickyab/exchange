package main

import (
	"clickyab.com/exchange/commands"

	"github.com/clickyab/services/config"
	_ "github.com/clickyab/services/fluentd"
	"github.com/clickyab/services/initializer"
	_ "github.com/clickyab/services/kv/redis"
	"github.com/clickyab/services/shell"
	"github.com/sirupsen/logrus"
)

const (
	exchange = "exchange.dev"
	host     = "demand.dev"
)

func main() {
	config.Initialize(commands.Organization, commands.AppName, commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()

	sig := shell.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
