package restful

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
	t := kv.NewAEAVStore(mp + name)
	t.IncSubKey("month", cpm)
	t.IncSubKey("month_count", 1)
	t.IncSubKey(dp, cpm)
	t.IncSubKey(dp+"_count", cpm)
	t.Save(month)
	t = kv.NewAEAVStore(wp + name)
	t.IncSubKey("week", cpm)
	t.IncSubKey(dp, cpm)
	t.Save(week)
	t = kv.NewAEAVStore(dp + name)
	t.IncSubKey("day", cpm)
	t.IncSubKey(hp, cpm)
	t.Save(day)
	t = kv.NewAEAVStore(hp + name)
	t.IncSubKey("hour", cpm)
	t.IncSubKey(ip, cpm)
	t.Save(hour)
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
	t := kv.NewAEAVStore(mp + name)
	m = t.SubKey("month")
	cc := t.SubKey("month_count")
	m = realVal(m, cc)

	t = kv.NewAEAVStore(wp + name)
	w = t.SubKey("week")
	cc = t.SubKey("week_count")
	w = realVal(w, cc)

	t = kv.NewAEAVStore(dp + name)
	d = t.SubKey("day")
	cc = t.SubKey("day_count")
	d = realVal(d, cc)

	t = kv.NewAEAVStore(hp + name)
	h = t.SubKey("hour")
	cc = t.SubKey("hour_count")
	h = realVal(h, cc)

	i = t.SubKey(ip)
	cc = t.SubKey(ip + "_count")
	i = realVal(i, cc)

	return
}
