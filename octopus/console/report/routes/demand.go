package routes

import (
	"context"
	"net/http"

	"errors"
	"strings"
	"time"

	"clickyab.com/exchange/octopus/console/user/aaa"
	"clickyab.com/exchange/octopus/console/user/authz"
	"clickyab.com/exchange/octopus/models"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/framework"
	"github.com/rs/xmux"
)

type demandReportResponse struct {
	Data  []models.DemandReport `json:"data"`
	Count int64                 `json:"count"`
}

// demand report in system
// @Route {
// 		url = /demand/:from/:to
//		method = get
//		middleware = authz.Authenticate
//		_sort_ = string, the sort and order like id:asc or id:desc available column "id","created_at","updated_at"
//		_c_ = integer , count per page
//		_p_ = integer , page number
//		400 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
//		200 = demandReportResponse
// }
func (c Controller) demand(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	currentUser := authz.MustGetUser(ctx)
	var toTime time.Time
	var res demandReportResponse
	p, count := framework.GetPageAndCount(r, false)
	from := xmux.Param(ctx, "from")
	if from == "" {
		c.BadResponse(w, errors.New("start date not valid"))
		return
	}
	to := xmux.Param(ctx, "to")
	fromTime, err := time.Parse("20060102", from)
	if err != nil {
		c.BadResponse(w, errors.New("start date not valid"))
		return
	}
	toTime, err = time.Parse("20060102", to)
	if err != nil {
		toTime = fromTime.AddDate(0, 0, 1)
	}
	fromTimeString := fromTime.Format("2006-01-02")
	toTimeString := toTime.Format("2006-01-02")
	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "id") {
		sort = ""
	}
	order := strings.ToUpper(parts[1])
	if !array.StringInArray(order, "ASC", "DESC") {
		order = aaa.DefaultOrder
	}
	result, num := models.NewManager().FillDemandReport(p, count, sort, order, fromTimeString, toTimeString, currentUser)
	res.Data = result
	res.Count = num
	c.OKResponse(w, res)
}
