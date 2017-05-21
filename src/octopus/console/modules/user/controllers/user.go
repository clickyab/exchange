package users

import (
	"services/httplib/controller"

	echo2 "gopkg.in/labstack/echo.v3"
)

//User struct controller
type User struct {
	controller.Base
}

//Routes user
func (User) Routes(r *echo2.Echo, mountPoint string) {
	r.GET("/login", nil, nil)
	r.POST("/login", nil, nil)
	r.GET("logout", nil, nil)
}

func init() {
	controller.Register(&User{})
}
