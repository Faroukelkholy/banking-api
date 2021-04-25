package server

import (
	"github.com/faroukelkholy/bank/internal/service/account"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Server struct holds echo instance that holds all mounted routes
type Server struct {
	echo *echo.Echo
}

//New Return Server instance
func New() *Server {
	srv := new(Server)
	srv.echo = echo.New()
	srv.init()
	return srv
}

//Start bind echo to a port and run echo server
func (srv *Server) Start(port string) error {
	return srv.echo.Start(port)
}

//init holds all middleware that can be useful
func (srv *Server) init() {
	srv.echo.Use(middleware.Logger())
}

//AddRoutesAS mount routes related to account service
func (srv *Server) AddRoutesAS(s account.Service) {
	srv.echo.GET("/accounts/:id", GAHandler(s))
}
