package core

import (
	"context"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"

	"errors"

	"clickyab.com/exchange/octopus/exchange/materialize"
	"github.com/clickyab/services/broker"
	"github.com/sirupsen/logrus"
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
		"track_id": imp.TrackID(),
		"type":     "provider",
	})
}

// Skip decide if provider should respond to demand or not
func (p *providerData) Skip() bool {
	x := atomic.AddInt64(&p.callRateTracker, 1)
	return x%100 >= int64(p.provider.CallRate())
}

func (p *providerData) watch(ctx context.Context, imp exchange.BidRequest) (res map[string]exchange.Advertise) {
	//in := time.Now()
	defer func() {
		//out := time.Since(in)
		jDem := materialize.DemandJob(
			imp,
			p.provider,
			res,
		)
		broker.Publish(jDem)
	}()

	log(imp).WithField("provider", p.provider.Name()).Debug("Watch IN for provider")
	defer log(imp).WithField("provider", p.provider.Name()).Debug("Watch OUT for provider")
	done := ctx.Done()
	assert.NotNil(done)

	res = make(map[string]exchange.Advertise)
	// the cancel is not required here. the parent is the hammer :)
	rCtx, _ := context.WithTimeout(ctx, p.timeout)

	chn := make(chan exchange.Advertise, len(imp.Slots()))
	go p.provider.Provide(rCtx, imp, chn)
	for {
		select {
		case <-done:
			// request is canceled
			return res
		case data, open := <-chn:
			if data != nil {
				res[data.SlotTrackID()] = data
			}
			if !open {
				return res
			}
		}
	}
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
func Call(ctx context.Context, imp exchange.BidRequest) map[string][]exchange.Advertise {
	rCtx, cnl := context.WithTimeout(ctx, maximumTimeout)
	defer cnl()

	wg := sync.WaitGroup{}
	l := len(allProviders)
	wg.Add(l)
	allRes := make(chan map[string]exchange.Advertise, l)
	lock.RLock()
	for i := range allProviders {
		go func(inner string) {
			defer wg.Done()
			if !demandIsAllowed(imp, allProviders[inner]) {
				return
			}
			p := allProviders[inner]
			res := p.watch(rCtx, imp)
			if res != nil {
				allRes <- res
			}
		}(i)
	}
	lock.RUnlock()

	wg.Wait()
	// The close is essential here.
	close(allRes)
	var limit int64
	if !imp.UnderFloor() {
		limit = imp.Source().FloorCPM()
	}
	log(imp).WithField("limit", limit).Debug("the limit")
	res := make(map[string][]exchange.Advertise)
	for provided := range allRes {
		log(imp).WithField("count", len(provided)).Debug("result from demand")
		for j := range provided {
			if provided[j].MaxCPM() >= limit {
				res[j] = append(res[j], provided[j])
			}
		}
	}

	return res
}

func demandIsAllowed(m exchange.BidRequest, d providerData) bool {
	for _, f := range filters {
		if f(m, d) {
			return false
		}
	}
	return true
}

func isSameProvider(impression exchange.BidRequest, data providerData) bool {
	return impression.Source().Name() == data.name
}

func notWhitelistCountries(impression exchange.BidRequest, data providerData) bool {
	if len(data.provider.WhiteListCountries()) == 0 {
		return false
	}
	return !contains(data.provider.WhiteListCountries(), impression.Location().Country().ISO)
}

func isExcludedDemands(impression exchange.BidRequest, data providerData) bool {
	return contains(impression.Source().Supplier().ExcludedDemands(), data.name)
}

func isNotSameMode(impression exchange.BidRequest, data providerData) bool {
	return impression.Source().Supplier().TestMode() != data.provider.TestMode()
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
