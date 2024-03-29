package materialize

import (
	"clickyab.com/exchange/octopus/exchange"
	"clickyab.com/exchange/octopus/exchange/materialize/jsonbackend"
	"github.com/clickyab/services/broker"
	"github.com/clickyab/services/config"

	"github.com/sirupsen/logrus"
)

const (
	jsonDriver  string = "json"
	emptyDriver        = "empty"
)

var (
	driver = config.RegisterString("octopus.exchange.materialize.driver", jsonDriver, "driver for materialize")
)

// DemandJob returns a job for demand
// TODO : add a duration to this. for better view this is important
func DemandJob(rq exchange.BidRequest, resp exchange.BidResponse, demand string) broker.Job {
	switch driver.String() {
	case jsonDriver:
		return jsonbackend.DemandJob(rq, resp, demand)
	case emptyDriver:
		return job{
			data:  []byte("demand job"),
			key:   "key",
			topic: "demand_job",
		}
	default:
		logrus.Panicf("invalid driver %s", driver.String())
		return nil
	}
}

// TODO change advertise to winner (new interface)

// WinnerJob return a broker job
func WinnerJob(bq exchange.BidRequest, bid exchange.Bid) broker.Job {
	switch driver.String() {
	case jsonDriver:
		return jsonbackend.WinnerJob(bq, bid)
	case emptyDriver:
		return job{
			data:  []byte("winner job"),
			key:   "key",
			topic: "winner_job",
		}
	default:
		logrus.Panicf("invalid driver %s", driver.String())
		return nil
	}
}

// ClickJob return a broker job
func ClickJob(source, supplier, demand, ip string) broker.Job {
	switch driver.String() {
	case jsonDriver:
		return jsonbackend.ClickJob(source, supplier, demand, ip)
	case emptyDriver:
		return job{
			data:  []byte("winner job"),
			key:   "key",
			topic: "winner_job",
		}
	default:
		logrus.Panicf("invalid driver %s", driver.String())
		return nil
	}
}

// ShowJob return a broker job
func ShowJob(demand string, IP string, winner float64, t string, supplier string, publisher string, profit float64) broker.Job {
	switch driver.String() {
	case jsonDriver:
		return jsonbackend.ShowJob(demand, IP, winner, t, supplier, publisher, profit)
	case emptyDriver:
		return job{
			data:  []byte("show job"),
			key:   "key",
			topic: "show_job",
		}
	default:
		logrus.Panicf("invalid driver %s", driver.String())
		return nil
	}
}

// ImpressionJob return a broker job
func ImpressionJob(imp exchange.BidRequest) broker.Job {
	switch driver.String() {
	case jsonDriver:
		return jsonbackend.ImpressionJob(imp)
	case emptyDriver:
		return job{
			data:  []byte("impression job"),
			key:   "key",
			topic: "impression_job",
		}
	default:
		logrus.Panicf("invalid driver %s", driver.String())
		return nil
	}
}
