package commands

import (
	// Fluentd service to hook the logrus commands
	_ "github.com/clickyab/services/fluentd"
	_ "github.com/clickyab/services/kv/redis"
)

const (
	// AppName the application name
	AppName string = "exchange"
	// Organization the organization name
	Organization = "clickyab"
	// Prefix the prefix for config loader from env
	Prefix = "EXC"
)
