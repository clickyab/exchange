// Code generated build with router DO NOT EDIT.

package routes

import (
	"sync"

	"clickyab.com/exchange/octopus/console/user/authz"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/initializer"
)

var once = sync.Once{}

// Routes return the route registered with this
func (c *Controller) Routes(r framework.Mux) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{}

		group := r.NewGroup("/user")

		/* Route {
			"Route": "/login",
			"Method": "POST",
			"Function": "Controller.login",
			"RoutePkg": "routes",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "loginPayload",
			"Resource": "",
			"Scope": ""
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m0 = append(m0, middleware.PayloadUnMarshallerGenerator(loginPayload{}))
		group.POST("routes-Controller-login", "/login", framework.Mix(c.login, m0...))
		// End route with key 0

		/* Route {
			"Route": "/logout",
			"Method": "GET",
			"Function": "Controller.logout",
			"RoutePkg": "routes",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("routes-Controller-logout", "/logout", framework.Mix(c.logout, m1...))
		// End route with key 1

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
