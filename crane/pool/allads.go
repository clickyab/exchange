package pool

import (
	"context"

	"time"

	"clickyab.com/exchange/crane/entity"
	"clickyab.com/exchange/crane/filter"
	"clickyab.com/exchange/crane/reducer"
	"github.com/Sirupsen/logrus"
)

var allFilters = []reducer.FilterFunc{
	filter.Category,
	filter.Country,
	filter.OS,
	filter.Province,
	filter.PublisherBlackList,
	filter.PublisherWhiteList,
	filter.AppSize,
	filter.WebSize,
	filter.VastSize,
	filter.Target,
}

// GetAllAds returns all ads after filtering them
func GetAllAds(ctx context.Context, imp entity.Impression) map[int][]entity.Advertise {
	fail := time.Since(lastTime) > 5*time.Minute

	if fail {
		logrus.Fatal("failed! restart the app please")
	}
	mixedFilters := reducer.Mix(allFilters...)
	return reducer.Apply(ctx, imp, adPool, mixedFilters)
}
