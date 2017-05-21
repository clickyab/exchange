package consule

import (
	"services/assert"
	"services/consul"
	"time"

	"github.com/hashicorp/consul/api"
)

type consul struct {
	ttl      time.Duration
	resource string
	lockCh   <-chan struct{}
	l        *api.Lock
}

func (c *consul) Lock() {
	var err error
	c.l, err = aconsul.Client.LockKey(c.resource)
	assert.Nil(err)
	stopCh := make(chan struct{})
	c.lockCh, err = c.l.Lock(stopCh)
	assert.Nil(err)
}

func (c *consul) Unlock() {
	err := c.l.Unlock()
	assert.Nil(err)
	<-c.lockCh
}

func (c *consul) Resource() string {
	return c.resource
}

func (c *consul) TTL() time.Duration {
	return c.ttl
}
