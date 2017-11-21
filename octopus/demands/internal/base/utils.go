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

	hiShare float64 = 100
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

func incCPM(name string, cpm float64) {
	mp := getMonthlyPattern()
	wp := getWeeklyPattern()
	dp := getDailyPattern()
	hp := getHourlyPattern()
	ip := getMinutlyPattern()
	t := kv.NewAEAVStore(mp+name, month)
	t.IncSubKey("month", int64(cpm*hiShare))
	t.IncSubKey("month_count", 1)
	t.IncSubKey(dp, int64(cpm*hiShare))
	t.IncSubKey(dp+"_count", int64(cpm*hiShare))
	t = kv.NewAEAVStore(wp+name, week)
	t.IncSubKey("week", int64(cpm*hiShare))
	t.IncSubKey(dp, int64(cpm*hiShare))

	t = kv.NewAEAVStore(dp+name, day)
	t.IncSubKey("day", int64(cpm*hiShare))
	t.IncSubKey(hp, int64(cpm*hiShare))

	t = kv.NewAEAVStore(hp+name, hour)
	t.IncSubKey("hour", int64(cpm*hiShare))
	t.IncSubKey(ip, int64(cpm*hiShare))

}

func realVal(all float64, count int64) float64 {
	if count > 0 {
		return all / float64(count)
	}
	return 0
}

func getCPM(name string) (m, w, d, h, i float64) {
	mp := getMonthlyPattern()
	wp := getWeeklyPattern()
	dp := getDailyPattern()
	hp := getHourlyPattern()
	ip := getMinutlyPattern()
	t := kv.NewAEAVStore(mp+name, 0)
	m = float64(t.SubKey("month")) / hiShare
	cc := t.SubKey("month_count")
	m = realVal(m, cc)

	t = kv.NewAEAVStore(wp+name, 0)
	w = float64(t.SubKey("week")) / hiShare
	cc = t.SubKey("week_count")
	w = realVal(w, cc)

	t = kv.NewAEAVStore(dp+name, 0)
	d = float64(t.SubKey("day")) / hiShare
	cc = t.SubKey("day_count")
	d = realVal(d, cc)

	t = kv.NewAEAVStore(hp+name, 0)
	h = float64(t.SubKey("hour")) / hiShare
	cc = t.SubKey("hour_count")
	h = realVal(h, cc)

	i = float64(t.SubKey(ip)) / hiShare
	cc = t.SubKey(ip + "_count")
	i = realVal(i, cc)

	return
}
