package influx

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
)

type initInflux struct {
	client.Client

	lock  sync.Mutex
	done  bool
	bp    client.BatchPoints
	cfg   client.BatchPointsConfig
	count int
}

var (
	looper = &initInflux{}
)

func (inf *initInflux) createClient() (client.Client, error) {
	switch protocol.String() {
	case "http":
		return client.NewHTTPClient(
			client.HTTPConfig{
				Addr:     server.String(),
				Username: user.String(),
				Password: password.String(),
				// TODO : why not config?
				Timeout: time.Second,
			},
		)
	case "udp":
		// UDP type dose not support
		return client.NewUDPClient(
			client.UDPConfig{
				Addr: server.String(),
			},
		)
	default:
		return nil, fmt.Errorf("invalid client protocol in config : %s", protocol.String())
	}
}

func (inf *initInflux) flush() error {
	if inf.count == 0 {
		return nil
	}
	err := inf.Client.Write(inf.bp)
	if err == nil {
		inf.count = 0
		// As I can see error is only related to Precision, so its afe to assert
		// it here
		inf.bp, err = client.NewBatchPoints(inf.cfg)
		assert.Nil(err)
	}

	return err
}

// it was a loop :))
func (inf *initInflux) inputLoop(p *client.Point) error {
	inf.lock.Lock()
	defer inf.lock.Unlock()
	if inf.done == true {
		return errors.New("server already stoped. discarding")
	}

	assert.NotNil(inf.Client, "[BUG] the client is not initialized")

	inf.bp.AddPoint(p)
	inf.count++
	if inf.count > bufferSize.Int() {
		err := inf.flush()
		if err != nil && inf.count > bufferSize.Int()*5 {
			logrus.Panicf("to many error, last one was: %s", err)
		}
	}
	return nil
}

func (inf *initInflux) Initialize(ctx context.Context) {
	inf.lock.Lock()
	defer inf.lock.Unlock()

	done := ctx.Done()
	assert.NotNil(done, "[BUG] context is not cancelable")
	var err error
	inf.Client, err = inf.createClient()
	assert.Nil(err)
	_, _, err = inf.Client.Ping(time.Second)
	assert.Nil(err)

	// TODO : more config
	inf.cfg = client.BatchPointsConfig{
		Database: database.String(),
	}

	inf.bp, err = client.NewBatchPoints(inf.cfg)
	assert.Nil(err)

	go func() {
		<-done
		inf.lock.Lock()
		defer inf.lock.Unlock()

		inf.done = true
		inf.flush()
	}()
}

// AddPoint is a function to add a point to db
func AddPoint(
	name string,
	tags map[string]string,
	fields map[string]interface{},
	t ...time.Time,
) error {
	p, err := client.NewPoint(name, tags, fields, t...)
	if err != nil {
		return err
	}

	return looper.inputLoop(p)
}

func init() {
	initializer.Register(looper, 0)
}
