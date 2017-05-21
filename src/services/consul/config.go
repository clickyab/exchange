package aconsul

import "services/config"

var (
	port *string
	host *string
)

func init() {

	port = config.RegisterString("services.consul.port", "8500", "Consul port")
	host = config.RegisterString("services.consul.host", "127.0.0.1", "Consul port")

}
