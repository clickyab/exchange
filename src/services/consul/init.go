package aconsul

import (
	"context"
	"services/assert"

	"github.com/Sirupsen/logrus"
	"github.com/hashicorp/consul/api"
)

var (
	// Client the actual pool to use with redis
	Client *api.Client
)

type initConsul struct {
}

// Initialize try to create a redis pool
func (initConsul) Initialize(ctx context.Context) {
	var err error
	Client, err = api.NewClient(&api.Config{Address: *host + *port})
	assert.Nil(err)
	// PING the server to make sure every thing is fine
	logrus.Debug("consul is ready.")
}
