// Code generated build with router DO NOT EDIT.

package routes

import (
	"sync"

	"clickyab.com/exchange/octopus/console/user/authz"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/initializer"
)

var once = sync.Once{}

// Routes return the route registered with this
func (c *Controller) Routes(r framework.Mux) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{}

		group := r.NewGroup("/report")

		/* Route {
			"Route": "/demand/:from/:to",
			"Method": "GET",
			"Function": "Controller.demand",
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
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/demand/:from/:to", framework.Mix(c.demand, m0...))
		// End route with key 0

		/* Route {
			"Route": "/exchange/:from/:to",
			"Method": "GET",
			"Function": "Controller.exchange",
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

		group.GET("/exchange/:from/:to", framework.Mix(c.exchange, m1...))
		// End route with key 1

		/* Route {
			"Route": "/supplier/:from/:to",
			"Method": "GET",
			"Function": "Controller.supplier",
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
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/supplier/:from/:to", framework.Mix(c.supplier, m2...))
		// End route with key 2

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
