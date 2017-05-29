package routes

import (
	"context"
	"net/http"
	"services/config"
	"services/initializer"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/fzerorubigd/xhandler"
	"github.com/fzerorubigd/xmux"
)

type initConsole struct{}

var (
	listenAddress = config.RegisterString("console.router.listen", ":85", "exchnage router listen address")
)

func (initConsole) Initialize(ctx context.Context) {
	routes(ctx)
}

func routes(ctx context.Context) {
	mux := xmux.New()

	mux.GET("/login", xhandler.HandlerFuncC(loginGet))
	mux.POST("/login", xhandler.HandlerFuncC(loginPost))
	mux.GET("/logout", xhandler.HandlerFuncC(logout))

	srv := &http.Server{Addr: *listenAddress, Handler: xhandler.New(ctx, mux)}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Debug(err)
		}
	}()
	logrus.Debugf("Server started on %s", *listenAddress)
	go func() {
		done := ctx.Done()
		if done != nil {
			<-done
			s, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
			srv.Shutdown(s)
			logrus.Debug("Server stopped")
		}
	}()
}

func init() {
	initializer.Register(&initConsole{}, 118)
}
