package dispatcher

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/assert"
	"github.com/sirupsen/logrus"

	"clickyab.com/exchange/octopus/exchange/materialize"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/safe"
	"github.com/clickyab/services/xlog"
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

func log(ctx context.Context, imp exchange.BidRequest) *logrus.Entry {
	return xlog.GetWithFields(ctx, logrus.Fields{
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
	defer safe.GoRoutine(func() {
		//out := time.Since(in)
		if data != nil && len(data.Bids()) != 0 {
			jDem := materialize.DemandJob(
				bq,
				data,
				p.name,
			)
			broker.Publish(jDem)
		}
	})

	log(ctx, bq).WithField("provider", p.provider.Name()).Debug("Watch IN for provider")
	defer log(ctx, bq).WithField("provider", p.provider.Name()).Debug("Watch OUT for provider")
	// the cancel is not required here. the parent is the hammer :)
	rCtx, _ := context.WithTimeout(ctx, p.timeout)

	go p.provider.Provide(rCtx, bq, chn)
	select {
	case <-rCtx.Done():
		// request is canceled
		return nil
	case x, open := <-chn:
		if x != nil && open {
			data = x
			return data
		}

		return nil
	}
}

// Register add a new demand to system. the timeout is the maximum timeout for this demand.
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

// ResetProviders remove all providers, for reload it again
func ResetProviders() {
	lock.Lock()
	defer lock.Unlock()

	allProviders = make(map[string]providerData)
}

// Call is the main function. calling this create a sequence of calls and all demands are called
// one by one. there is a maximum timeout, but if you need a certain cancellation method, use the
// context for that.
// since the result is from multiple demands, the result is an array
func Call(ctx context.Context, req exchange.BidRequest) []exchange.BidResponse {
	rCtx, cnl := context.WithTimeout(ctx, maximumTimeout)
	defer cnl()

	wg := sync.WaitGroup{}
	l := len(allProviders)
	wg.Add(l)
	var allRes = make(chan exchange.BidResponse, l)
	lock.RLock()
	for i := range allProviders {
		go func(inner providerData) {
			defer wg.Done()
			if !demandIsAllowed(req, inner) {
				return
			}
			res := inner.watch(rCtx, req)
			if res != nil {
				allRes <- res
			}
		}(allProviders[i])
	}
	lock.RUnlock()

	wg.Wait()
	// The close is essential here.
	close(allRes)

	var response []exchange.BidResponse
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

func acceptCurrnecy(bq exchange.BidRequest, data providerData) bool {
	currencies := data.provider.Currencies()
	if len(currencies) == 0 {
		return false
	}
	c := bq.Imp()[0].Currency()
	for _, v := range currencies {
		if c == v {
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
	// TODO : change test mode to integer and change it to group functionality
	// if we use number we can create multiple separate network and its a cool
	// functionality, but not yet.
	return data.provider.TestMode() != func() bool {
		if t := bq.Inventory().Supplier().TestMode(); t {
			return t
		}
		return bq.Test()
	}()
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
		acceptCurrnecy,
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
