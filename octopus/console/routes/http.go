package routes

import (
	"context"
	"time"

	"clickyab.com/exchange/services/config"
	"clickyab.com/exchange/services/initializer"

	"clickyab.com/exchange/services/eav"
	_ "clickyab.com/exchange/services/eav/redis"
	"github.com/Sirupsen/logrus"
	"gopkg.in/labstack/echo.v3"
)

type initConsole struct{}

var (
	listenAddress = config.RegisterString("console.router.listen", ":3500", "exchnage router listen address")
	tokenStore    eav.Kiwi
)

func (initConsole) Initialize(ctx context.Context) {
	routes(ctx)
}

func routes(ctx context.Context) {
	e := echo.New()

	e.GET("/login", loginGet)
	e.POST("/login", loginPost)
	e.GET("/logout", logout, auth)

	go func() {
		if err := e.Start(":3500"); err != nil {
			logrus.Debug(err)
		}
	}()

	logrus.Debugf("Server started on %s", *listenAddress)
	go func() {
		done := ctx.Done()
		if done != nil {
			<-done
			s, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
			e.Shutdown(s)
			logrus.Debug("Server stopped")
		}
	}()
}

func init() {
	tokenStore = eav.NewEavStore("token")
	initializer.Register(&initConsole{}, 118)
}
