package main

import (
	"commands"
	_ "octopus/demands"
	_ "octopus/router/restful"
	"services/config"
	"services/initializer"
	_ "services/statistic/redis"

	"github.com/Sirupsen/logrus"
)

func main() {
	config.Initialize(commands.Organization, commands.AppName, commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()

	sig := commands.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
