package main

import (
	"clickyab.com/exchange/commands"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/shell"
	"github.com/sirupsen/logrus"

	_ "clickyab.com/exchange/demos"
)

func main() {
	config.Initialize(commands.Organization, "demandDemo", commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()

	sig := shell.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
