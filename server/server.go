package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ymohl-cl/gopkg/server/config"
	"github.com/ymohl-cl/gopkg/server/handler"
	"github.com/ymohl-cl/gopkg/server/jsonvalidator"
)

// Server component
type Server struct {
	driver *echo.Echo
	config config.Config
}

// New server http
func New(appName string) (Server, error) {
	var err error
	var h handler.Handler

	s := Server{driver: echo.New()}
	if s.config, err = config.NewConfig(appName); err != nil {
		return Server{}, err
	}
	s.driver.Validator = jsonvalidator.New()
	if h, err = handler.New(appName, s.config.JWTKey); err != nil {
		return Server{}, err
	}
	s.driver.Use(middleware.Logger())
	s.driver.GET("/ping", handler.Ping)
	s.driver.POST("/login", h.Login)
	s.driver.POST("/register", h.NewUser)
	return s, nil
}

// SubRouter return with the prefix path specified on parameter
func (s Server) SubRouter(prefix string) *echo.Group {
	router := s.driver.Group(prefix)
	router.Use(middleware.JWT([]byte(s.config.JWTKey)))
	return router
}

// Start run the server
func (s Server) Start() error {
	var err error

	defer s.driver.Close()
	if s.config.SSL.Enable {
		if err = s.driver.StartTLS(":"+s.config.Port, s.config.SSL.Certificate, s.config.SSL.Key); err != nil {
			return err
		}
	} else {
		if err = s.driver.Start(":" + s.config.Port); err != nil {
			return err
		}
	}
	return nil
}
