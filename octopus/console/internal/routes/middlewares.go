package user

import (
	"context"
	"net/http"

	"clickyab.com/exchange/octopus/console/internal/aaa"
	"github.com/clickyab/services/eav"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/trans"
)

type dataType string

type data struct {
	user  *aaa.User
	token string
}

const dataKey dataType = "__user_data__"

// Authenticate is a middleware used for authorization in exchange console
func Authenticate(next framework.Handler) framework.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token != "" {
			val := eav.NewEavStore(token).SubKey("token")
			if val != "" {
				usr, err := aaa.NewAaaManager().GetUserByToken(val)
				if err == nil {
					ctx = context.WithValue(ctx, dataKey, data{
						token: token,
						user:  usr,
					})
					next(ctx, w, r)
					return
				}
			}
		}
		framework.JSON(w, http.StatusUnauthorized, controller.ErrorResponseSimple{Error: trans.E("Unauthorized")})
	}
}
