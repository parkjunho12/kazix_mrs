package service

import (
	"sync"

	"github.com/gorilla/sessions"
	"github.com/kazix-mrs/factory"
	"github.com/kazix-mrs/service/webservice"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	MainServer struct {
		Fac        *factory.Factory
		mutex      sync.Mutex
		webHandler *webservice.WebHandler
	}
)

func (self *MainServer) StartService() {

	self.webHandler = webservice.NewWebHandler(self.Fac)

	e := echo.New()
	store := sessions.NewCookieStore([]byte("MRS-REST"))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))

	e.Use(middleware.CORS())
}
