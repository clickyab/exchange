package main

import (
	"commands"
	_ "octopus/console/models"
	_ "octopus/console/modules/user"
	"services/config"
	"services/initializer"

	"github.com/Sirupsen/logrus"
)

func main() {
	config.Initialize(commands.Organization, commands.AppName, commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()

	sig := commands.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
