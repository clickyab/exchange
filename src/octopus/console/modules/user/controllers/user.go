package users

import (
	"services/httplib/controller"

	echo2 "gopkg.in/labstack/echo.v3"
)

//User struct controller
type User struct {
	controller.Base
}

// Routes adds routers to listener
func (u User) Routes(r *echo2.Echo, mountPoint string) {
	r.GET("/login", nil)
	r.POST("/login", login)
	r.GET("/logout", nil)
}

func init() {
	controller.Register(&User{})
}
