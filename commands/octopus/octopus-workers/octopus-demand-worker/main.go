package main

import (
	"clickyab.com/exchange/commands"
	_ "clickyab.com/exchange/octopus/console/report/generator"
	_ "clickyab.com/exchange/octopus/workers/demand"
	_ "clickyab.com/exchange/octopus/workers/manager"
	_ "github.com/clickyab/services/broker/selector"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/initializer"
	_ "github.com/clickyab/services/kv/redis"
	_ "github.com/clickyab/services/mysql/connection/mysql"
	"github.com/clickyab/services/shell"
	"github.com/sirupsen/logrus"
)

func main() {
	config.Initialize(commands.Organization, commands.AppName, commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()

	sig := shell.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
