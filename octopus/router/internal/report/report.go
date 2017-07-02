package report

import (
	"context"
	"net/http"
	"time"

	"clickyab.com/exchange/octopus/models"
	"github.com/rs/xmux"
)

// Update is the demand status route
func Update(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ts := xmux.Param(ctx, "date")
	if ts == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, e := time.Parse("20060102", ts)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if t.Unix() > time.Now().Unix() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m := models.NewManager()
	m.UpdateDemandRange(t, t)
	m.UpdateSupplierRange(t, t)
	m.UpdateExchangeRange(t, t)
}

// UpdateRange is the demand status route
func UpdateRange(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	sf := xmux.Param(ctx, "from")
	st := xmux.Param(ctx, "to")

	if sf == "" || st == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tf, ef := time.Parse("20060102", sf)
	tt, et := time.Parse("20060102", st)

	if ef != nil || et != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if tf.Unix() > time.Now().Unix() || tt.Unix() > time.Now().Unix() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := models.NewManager()
	m.UpdateDemandRange(tf, tt)
	m.UpdateSupplierRange(tf, tt)
	m.UpdateExchangeRange(tf, tt)
}
