package pool

import (
	"context"
	"sync"
	"time"

	"clickyab.com/exchange/crane/entity"
	"clickyab.com/exchange/crane/pool/internal"

	"github.com/clickyab/services/config"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/safe"

	"github.com/Sirupsen/logrus"
)

var (
	adPool         []entity.Advertise
	adPoolCoolDown = config.GetDurationDefault("crane.adpool.cooldown", time.Minute)
	lastTime       time.Time
)

type constructor struct {
	locker *sync.RWMutex
}

func (c *constructor) Initialize(ctx context.Context) {
	lastTime = time.Now()
	m := internal.NewManager()
	safe.GoRoutine(func() {
		for {
			ap, err := m.GetAllActiveAds()
			if err != nil {
				time.Sleep(adPoolCoolDown)
				continue
			}
			logrus.Debug("ad pool updated")
			c.locker.Lock()
			adPool = ap
			lastTime = time.Now()
			c.locker.Unlock()
			time.Sleep(adPoolCoolDown)

		}
	})
}

func init() {
	initializer.Register(&constructor{locker: &sync.RWMutex{}}, 500)
}
