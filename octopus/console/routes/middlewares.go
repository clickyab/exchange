package routes

import (
	"net/http"

	"clickyab.com/exchange/octopus/console/internal/manager"
	"gopkg.in/labstack/echo.v3"
)

const userData = "__user_data__"
const tokenData = "__token__"

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")
		if token != "" {
			val := tokenStore.SubKey(token)
			if val != "" {
				user, err := manager.NewManager().GetUserByToken(val)
				if err == nil {
					c.Set(userData, user)
					c.Set(tokenData, token)
					return next(c)
				}
			}
		}
		return c.JSON(http.StatusUnauthorized, struct {
			error string
		}{
			error: "unauthorized",
		})
	}
}
