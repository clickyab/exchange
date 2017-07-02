package main

import (
	"clickyab.com/exchange/commands"
	"github.com/clickyab/services/config"
	_ "github.com/clickyab/services/dset/redis"
	_ "github.com/clickyab/services/eav/redis"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/shell"
	_ "github.com/clickyab/services/statistic/redis"
	_ "github.com/clickyab/services/store/redis"

	"github.com/clickyab/services/dlock"
	"github.com/clickyab/services/dlock/mock"

	_ "clickyab.com/exchange/octopus/console/report/generator"
	"github.com/Sirupsen/logrus"
	_ "github.com/clickyab/services/mysql/connection/mysql"
)

func main() {

	// TODO : after implementing dlock backend remove the next line
	dlock.Register(mock.NewMockDistributedLocker)

	config.Initialize(commands.Organization, commands.AppName, commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()

	sig := shell.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
