package users

import (
	"services/httplib/controller"

	echo2 "gopkg.in/labstack/echo.v3"
)

type User struct {
	controller.Base
}

func (User) Routes(r *echo2.Echo, mountPoint string) {
	r.GET("/login", nil, nil)
	r.POST("/login", nil, nil)
	r.GET("logout", nil, nil)
}

func init() {
	controller.Register(&User{})
}
