package main

import (
	"context"
	"time"

	"fmt"

	"clickyab.com/exchange/commands"
	"clickyab.com/exchange/services/broker"
	_ "clickyab.com/exchange/services/broker/selector"
	"clickyab.com/exchange/services/config"
	"clickyab.com/exchange/services/initializer"
	"clickyab.com/exchange/services/safe"
	"clickyab.com/exchange/services/shell"
	"github.com/Sirupsen/logrus"
)

type consumer struct {
}

func (c *consumer) Topic() string {
	return "test"
}

func (c *consumer) Queue() string {
	return "test_queue"
}

func (c *consumer) Consume() chan<- broker.Delivery {

	ch := make(chan broker.Delivery)
	safe.ContinuesGoRoutine(func(cnl context.CancelFunc) {
		var i int
		for job := range ch {
			fmt.Println("get a job")
			i++
			if i == 3 {
				logrus.Debug("rejecting only this")
				job.Reject(false)
				continue
			}

			if i == 5 {
				logrus.Debug("ack all")
				job.Nack(true, false)
				continue
			}
		}
	}, time.Second)
	return ch

}

func main() {
	config.Initialize("s", "s", "s", commands.DefaultConfig())
	defer initializer.Initialize()()
	c := &consumer{}

	broker.RegisterConsumer(c)

	shell.WaitExitSignal()
	logrus.Debug("bye")
}
