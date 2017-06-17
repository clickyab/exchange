package main

import (
	"clickyab.com/exchange/commands"
	_ "clickyab.com/exchange/octopus/console"
	"github.com/clickyab/services/config"

	"github.com/clickyab/services/initializer"

	"github.com/Sirupsen/logrus"
	_ "github.com/clickyab/services/eav/redis"
	"github.com/clickyab/services/shell"
	_ "golang.org/x/crypto/bcrypt"
)

func main() {
	config.Initialize(commands.Organization, commands.AppName, commands.Prefix, commands.DefaultConfig())
	defer initializer.Initialize()()

	sig := shell.WaitExitSignal()
	logrus.Debugf("%s received, exiting...", sig.String())
}
