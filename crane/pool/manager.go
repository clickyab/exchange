package pool

import (
	"context"
	"sync"
	"time"

	"clickyab.com/exchange/crane/entity"
	"clickyab.com/exchange/crane/pool/internal"
	"clickyab.com/exchange/services/config"
	"clickyab.com/exchange/services/initializer"
	"clickyab.com/exchange/services/safe"
	"github.com/Sirupsen/logrus"
)

var (
	adPool         []entity.Advertise
	adPoolCoolDown = config.GetDurationDefault("crane.adpool.cooldown", time.Second*30)
)

type constructor struct {
	locker *sync.RWMutex
}

func (c *constructor) Initialize(ctx context.Context) {
	m := internal.NewManager()
	safe.GoRoutine(func() {
		for {
			c.fillAds(m)
			logrus.Debug("ad pool updated")
			time.Sleep(adPoolCoolDown)
		}
	})
}

func (c *constructor) fillAds(m *internal.Manager) {
	ads, err := m.GetAllActiveAds()
	if err != nil {
		logrus.Warn(err.Error())
		return
	}

	c.locker.Lock()
	adPool = ads
	c.locker.Unlock()
}

func init() {
	initializer.Register(&constructor{locker: &sync.RWMutex{}}, 500)
}
