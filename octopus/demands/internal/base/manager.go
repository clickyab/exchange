package base

import (
	"bytes"
	"context"
	"net/http"

	"clickyab.com/exchange/octopus/exchange"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/xlog"
	"github.com/sirupsen/logrus"
)

// Manager is the model manager
type Manager struct {
	mysql.Manager
}

// Initialize the manager. nothing to do, just keep it in interface shape
func (m *Manager) Initialize() {

}

// NewManager return a new manager object
func NewManager() *Manager {
	return &Manager{}
}

func init() {
	mysql.Register(&Manager{})
}

// Provide provide ad for specified bid request
func Provide(ctx context.Context, dem exchange.Demand, bq exchange.BidRequest, ch chan exchange.BidResponse) {
	defer close(ch)
	if !dem.HasLimits() {
		return
	}
	buf := &bytes.Buffer{}

	header := dem.RenderBidRequest(ctx, buf, bq)
	req, err := http.NewRequest("POST", dem.EndPoint(), buf)
	req.Header = header
	if err != nil {
		logrus.Debug(err)
		return
	}

	xlog.Get(ctx).WithField("key", dem.Name()).Debug("calling demand")
	resp, err := dem.Client().Do(req.WithContext(ctx))
	if err != nil {
		logrus.Debug(err)
		return
	}

	ch <- dem.GetBidResponse(ctx, resp, bq.Inventory().Supplier())
}
