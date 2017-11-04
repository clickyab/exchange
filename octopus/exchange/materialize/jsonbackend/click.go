package jsonbackend

import (
	"encoding/json"

	"time"

	"github.com/clickyab/services/broker"
	"github.com/sirupsen/logrus"
)

type click struct {
	data map[string]interface{}
	key  string

	src []byte
}

// Encode encodes the data in map
func (c *click) Encode() ([]byte, error) {
	if len(c.src) == 0 {
		var err error
		c.src, err = json.Marshal(c.data)
		if err != nil {
			return nil, err
		}
	}

	return c.src, nil
}

// Length represent the len of data
func (c *click) Length() int {
	x, _ := c.Encode()
	return len(x)
}

// Topic shows the topic for
func (*click) Topic() string {
	return "click"
}

// Key returns the job key
func (c *click) Key() string {
	return c.key
}

// Report reports error
func (*click) Report() func(error) {
	return func(err error) {
		if err != nil {
			logrus.Warn(err)
		}
	}
}

// ClickJob return a broker job
func ClickJob(source, supplier, demand, ip string) broker.Job {
	return &click{
		data: map[string]interface{}{
			"source":   source,
			"supplier": supplier,
			"demand":   demand,
			"time":     time.Now(),
		},
		key: ip,
	}
}
