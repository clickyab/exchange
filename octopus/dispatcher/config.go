package dispatcher

import (
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
)

// I do not use the Register mode since I need to handle the zero and very large value.
// do not like the idea, but its ok for now.
var maximumTimeout time.Duration

type cfgInitializer struct {
}

func (ci *cfgInitializer) Initialize() config.DescriptiveLayer {
	l := config.NewDescriptiveLayer()
	l.Add("maximum time to wait for demands to respond", "exchange.core.maximum_timeout", time.Second)
	return l
}

func (ci *cfgInitializer) Loaded() {
	maximumTimeout = config.GetDuration("exchange.core.maximum_timeout")
	assert.True(maximumTimeout > 0)
	assert.True(maximumTimeout < 10*time.Second)
}

func init() {
	config.Register(&cfgInitializer{})
}
