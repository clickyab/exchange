package core

import (
	"context"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"
	"github.com/sirupsen/logrus"

	"errors"

	"clickyab.com/exchange/octopus/exchange/materialize"
	"github.com/clickyab/services/broker"
)

var (
	allProviders = make(map[string]providerData)
	lock         = &sync.RWMutex{}
	filters      []func(exchange.BidRequest, providerData) bool
)

type providerData struct {
	name            string
	provider        exchange.Demand
	timeout         time.Duration
	callRateTracker int64
}

func log(imp exchange.BidRequest) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"track_id": imp.ID(),
		"type":     "provider",
	})
}

// Skip decide if provider should respond to demand or not
func (p *providerData) Skip() bool {
	x := atomic.AddInt64(&p.callRateTracker, 1)
	return x%100 >= int64(p.provider.CallRate())
}

func (p *providerData) watch(ctx context.Context, bq exchange.BidRequest) exchange.BidResponse {
	var data exchange.BidResponse
	chn := make(chan exchange.BidResponse, 1)
	defer func() {
		//out := time.Since(in)
		if data != nil && len(data.Bids()) != 0 {
			jDem := materialize.DemandJob(
				bq,
				<-chn,
			)
			broker.Publish(jDem)
		}
	}()

	log(bq).WithField("provider", p.provider.Name()).Debug("Watch IN for provider")
	defer log(bq).WithField("provider", p.provider.Name()).Debug("Watch OUT for provider")
	done := ctx.Done()
	assert.NotNil(done)

	// the cancel is not required here. the parent is the hammer :)
	rCtx, _ := context.WithTimeout(ctx, p.timeout)

	go p.provider.Provide(rCtx, bq, chn)
	select {
	case <-done:
		// request is canceled
		return nil
	case x, open := <-chn:
		if x != nil && open {
			data = x
			return data
		}
	}
	assert.NotNil(data)
	return data
}

// Register is used to handle new layer in system
func Register(provider exchange.Demand, timeout time.Duration) {
	lock.Lock()
	defer lock.Unlock()
	name := provider.Name()

	_, ok := allProviders[name]
	assert.False(ok, "[BUG] provider is already registered")

	allProviders[name] = providerData{
		name:            name,
		provider:        provider,
		timeout:         timeout,
		callRateTracker: rand.Int63n(100),
	}

	logrus.WithField("type", "register_demand").Debugf("demand with name %s is registered", name)
}

// ResetProviders remove all providers
func ResetProviders() {
	lock.Lock()
	defer lock.Unlock()

	allProviders = make(map[string]providerData)
}

// Call is for getting the current ads for this imp
func Call(ctx context.Context, req exchange.BidRequest) []exchange.BidResponse {
	rCtx, cnl := context.WithTimeout(ctx, maximumTimeout)
	defer cnl()

	wg := sync.WaitGroup{}
	l := len(allProviders)
	wg.Add(l)
	var allRes = make(chan exchange.BidResponse, l)
	lock.RLock()
	for i := range allProviders {
		go func(inner string) {
			defer wg.Done()
			if !demandIsAllowed(req, allProviders[inner]) {
				return
			}
			p := allProviders[inner]
			res := p.watch(rCtx, req)
			if res != nil {
				allRes <- res
			}
		}(i)
	}
	lock.RUnlock()

	wg.Wait()
	// The close is essential here.
	close(allRes)

	response := []exchange.BidResponse{}
	for i := range allRes {
		response = append(response, i)
	}
	return response
}

func demandIsAllowed(m exchange.BidRequest, d providerData) bool {
	for _, f := range filters {
		if f(m, d) {
			return false
		}
	}
	return true
}

func isSameProvider(bq exchange.BidRequest, data providerData) bool {
	return bq.Inventory().Name() == data.name
}

func notWhitelistCountries(bq exchange.BidRequest, data providerData) bool {
	if len(data.provider.WhiteListCountries()) == 0 {
		return false
	}
	return !contains(data.provider.WhiteListCountries(), bq.Device().Geo().Country().ISO)
}

func isExcludedDemands(bq exchange.BidRequest, data providerData) bool {
	return contains(bq.WhiteList(), data.name)
}

func isNotSameMode(bq exchange.BidRequest, data providerData) bool {
	return bq.Test() != data.provider.TestMode()
}

func contains(s []string, t string) bool {
	for _, a := range s {
		if a == t {
			return true
		}
	}
	return false
}

func init() {
	filters = []func(exchange.BidRequest, providerData) bool{
		isNotSameMode,
		isSameProvider,
		notWhitelistCountries,
		isExcludedDemands,
	}
}

// GetDemand return demand by its name
func GetDemand(name string) (exchange.Demand, error) {
	lock.RLock()
	defer lock.RUnlock()
	val, found := allProviders[name]
	if !found {
		return nil, errors.New("demand not found")
	}
	return val.provider, nil
}
