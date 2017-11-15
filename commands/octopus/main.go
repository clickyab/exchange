package main

import (
	"clickyab.com/exchange/commands"
	_ "clickyab.com/exchange/octopus/console/report/routes"
	_ "clickyab.com/exchange/octopus/console/user/aaa"
	_ "clickyab.com/exchange/octopus/console/user/routes"
	_ "clickyab.com/exchange/octopus/demands"
	_ "clickyab.com/exchange/octopus/router"
	_ "github.com/clickyab/services/broker/selector"
	"github.com/clickyab/services/config"
	_ "github.com/clickyab/services/kv/redis"

	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/shell"

	_ "github.com/clickyab/services/fluentd"
	_ "github.com/clickyab/services/mysql/connection/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	config.Initialize(commands.Organization, commands.AppName, commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()

	sig := shell.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
