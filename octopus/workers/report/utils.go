package report

import (
	"time"

	"clickyab.com/exchange/octopus/workers/internal"
	"clickyab.com/exchange/services/assert"
)

func factTableID(tm time.Time) int64 {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2017-03-21T00:00:00.000Z"
	epoch, err := time.Parse(layout, str)
	assert.Nil(err)
	return int64(tm.Sub(epoch).Hours()) + 1
}

// FactTableYesterdayID is a helper function to get the fact table for yesterday id from time
func factTableYesterdayID(tm time.Time) (int64, int64) {
	y, m, d := tm.Date()
	from := time.Date(y, m, d, 0, 0, 1, 0, time.UTC)
	to := time.Date(y, m, d, 23, 59, 59, 0, time.UTC)
	return internal.FactTableID(from), internal.FactTableID(to)
}

func factTableRange(t time.Time) (int64, int64) {
	y, m, d := t.Date()
	from := time.Date(y, m, d, 0, 0, 1, 0, time.UTC)
	to := time.Date(y, m, d, 23, 59, 59, 0, time.UTC)
	return factTableID(from), factTableID(to)
}
