package base

import (
	"fmt"
	"time"

	"github.com/clickyab/services/kv"
)

const (
	hour  = time.Hour
	day   = 24 * hour
	week  = 7 * day
	month = 31 * day
)

func getMonthlyPattern() string {
	return time.Now().Format("200601")
}

func getWeeklyPattern() string {
	t := time.Now()
	_, w := t.ISOWeek()

	return time.Now().Format("2006w") + fmt.Sprintf("%02d", w)

}

func getDailyPattern() string {
	return time.Now().Format("20060102")
}

func getHourlyPattern() string {
	return time.Now().Format("2006010203")
}

func getMinutlyPattern() string {
	return time.Now().Format("200601020304")
}

func incCPM(name string, cpm int64) {
	mp := getMonthlyPattern()
	wp := getWeeklyPattern()
	dp := getDailyPattern()
	hp := getHourlyPattern()
	ip := getMinutlyPattern()
	t := kv.NewAEAVStore(mp+name, month)
	t.IncSubKey("month", cpm)
	t.IncSubKey("month_count", 1)
	t.IncSubKey(dp, cpm)
	t.IncSubKey(dp+"_count", cpm)
	t = kv.NewAEAVStore(wp+name, week)
	t.IncSubKey("week", cpm)
	t.IncSubKey(dp, cpm)

	t = kv.NewAEAVStore(dp+name, day)
	t.IncSubKey("day", cpm)
	t.IncSubKey(hp, cpm)

	t = kv.NewAEAVStore(hp+name, hour)
	t.IncSubKey("hour", cpm)
	t.IncSubKey(ip, cpm)

}

func realVal(all, count int64) int64 {
	if count > 0 {
		return all / count
	}
	return 0
}

func getCPM(name string) (m, w, d, h, i int64) {
	mp := getMonthlyPattern()
	wp := getWeeklyPattern()
	dp := getDailyPattern()
	hp := getHourlyPattern()
	ip := getMinutlyPattern()
	t := kv.NewAEAVStore(mp+name, 0)
	m = t.SubKey("month")
	cc := t.SubKey("month_count")
	m = realVal(m, cc)

	t = kv.NewAEAVStore(wp+name, 0)
	w = t.SubKey("week")
	cc = t.SubKey("week_count")
	w = realVal(w, cc)

	t = kv.NewAEAVStore(dp+name, 0)
	d = t.SubKey("day")
	cc = t.SubKey("day_count")
	d = realVal(d, cc)

	t = kv.NewAEAVStore(hp+name, 0)
	h = t.SubKey("hour")
	cc = t.SubKey("hour_count")
	h = realVal(h, cc)

	i = t.SubKey(ip)
	cc = t.SubKey(ip + "_count")
	i = realVal(i, cc)

	return
}
