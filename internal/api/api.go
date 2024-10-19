package api

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/api/routes"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/config"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract"
)

type HTTPServer struct {
	e *echo.Echo
	c *contract.Contract
}

func NewHTTPServer(c *contract.Contract) HTTPServer {
	e := echo.New()

	return HTTPServer{
		e: e,
		c: c,
	}
}

func (s *HTTPServer) RunHTTPServer() {
	v1 := routes.InitVersion(s.e, "/api/v1", s.c)
	routes.RouteV1(&v1)

	port := fmt.Sprintf(":%d", config.Config.App.Port)
	if err := s.e.Start(port); err != nil {
		log.Fatalf("Failed to running server: %v", err)
	}
}
