package generator

import (
	"context"

	"time"

	"clickyab.com/exchange/octopus/models"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/safe"
)

type mysqlInitializer struct {
}

func doJob() (err error) {
	m := models.NewManager()
	err = m.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if e := recover(); e != nil {
			m.Rollback()
			return
		}
		if err != nil {
			m.Rollback()
		} else {
			m.Commit()
		}
	}()
	last := time.Now().Add(-time.Hour * 24)
	now := time.Now()
	m.UpdateDemandRange(last, now)
	m.UpdateSupplierRange(last, now)
	m.UpdateExchangeRange(last, now)

	return err
}

func reportGenerator(context.CancelFunc) {
	doJob()
	for {
		t := time.After(time.Hour)
		select {
		case <-t:
			doJob()
		}
	}
}

func (mi *mysqlInitializer) Initialize() {
	safe.ContinuesGoRoutine(reportGenerator, time.Minute)

}

func init() {
	mysql.Register(&mysqlInitializer{})
}
