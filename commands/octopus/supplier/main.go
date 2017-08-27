package main

import (
	"clickyab.com/exchange/commands"
	_ "clickyab.com/exchange/commands/octopus/supplier/logic"
	"github.com/Sirupsen/logrus"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/shell"
)

func main() {
	config.Initialize(commands.Organization, commands.AppName, commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()

	sig := shell.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
