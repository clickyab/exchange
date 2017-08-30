package commands

import (
	// Fluentd service to hook the logrus commands
	_ "github.com/clickyab/services/fluentd"
)

const (
	// AppName the application name
	AppName string = "exchange"
	// Organization the organization name
	Organization = "clickyab"
	// Prefix the prefix for config loader from env
	Prefix = "EXC"
)
