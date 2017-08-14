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
	"github.com/Sirupsen/logrus"
	"github.com/clickyab/services/broker"
)

var (
	allProviders = make(map[string]providerData)
	lock         = &sync.RWMutex{}
	filters      []func(exchange.Impression, providerData) bool
)

type providerData struct {
	name            string
	provider        exchange.Demand
	timeout         time.Duration
	callRateTracker int64
}

// Skip decide if provider should respond to demand or not
func (p *providerData) Skip() bool {
	x := atomic.AddInt64(&p.callRateTracker, 1)
	return x%100 >= int64(p.provider.CallRate())
}

func (p *providerData) watch(ctx context.Context, imp exchange.Impression) (res map[string]exchange.Advertise) {
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

	logrus.Debugf("Watch in for %s", p.provider.Name())
	defer logrus.Debugf("Watch out for %s", p.provider.Name())
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

	logrus.Debugf("demand with name %s is registered", name)
}

// ResetProviders remove all providers
func ResetProviders() {
	lock.Lock()
	defer lock.Unlock()

	allProviders = make(map[string]providerData)
}

// Call is for getting the current ads for this imp
func Call(ctx context.Context, imp exchange.Impression) map[string][]exchange.Advertise {
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
	logrus.Debugf("the limit is %d", limit)
	res := make(map[string][]exchange.Advertise)
	for provided := range allRes {
		logrus.Debugf("get a result contain %d ads ", len(provided))
		for j := range provided {
			if provided[j].MaxCPM() >= limit {
				res[j] = append(res[j], provided[j])
			}
		}
	}

	return res
}

func demandIsAllowed(m exchange.Impression, d providerData) bool {
	for _, f := range filters {
		if f(m, d) {
			return false
		}
	}
	return true
}

func isSameProvider(impression exchange.Impression, data providerData) bool {
	return impression.Source().Name() == data.name
}

func notWhitelistCountries(impression exchange.Impression, data providerData) bool {
	if len(data.provider.WhiteListCountries()) == 0 {
		return false
	}
	return !contains(data.provider.WhiteListCountries(), impression.Location().Country().ISO)
}

func isExcludedDemands(impression exchange.Impression, data providerData) bool {
	return contains(impression.Source().Supplier().ExcludedDemands(), data.name)
}

func isNotSameMode(impression exchange.Impression, data providerData) bool {
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
	filters = []func(exchange.Impression, providerData) bool{
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
