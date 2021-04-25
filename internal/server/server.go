package server

import (
	"github.com/faroukelkholy/bank/internal/service/account"
	"github.com/faroukelkholy/bank/internal/service/customer"
	"github.com/faroukelkholy/bank/internal/service/transaction"
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
	srv.echo.GET("/accounts/:id/transactions", GATsHandler(s))
}

//AddRoutesCS mount routes related to customer service
func (srv *Server) AddRoutesCS(s customer.Service) {
	srv.echo.POST("/customers/:id/accounts", CCAHandler(s))
}

//AddRoutesTS mount routes related to customer service
func (srv *Server) AddRoutesTS(s transaction.Service) {
	srv.echo.POST("/transactions", CTHandler(s))
}
