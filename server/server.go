package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server component
type Server struct {
	driver *echo.Echo
	config Config
}

// New server http
func New(appName string) (Server, error) {
	var err error

	s := Server{driver: echo.New()}
	if s.config, err = NewConfig(appName); err != nil {
		return Server{}, err
	}
	s.driver.Use(middleware.Logger())
	s.driver.GET("/ping", Ping)
	return s, nil
}

// SubRouter return with the prefix path specified on parameter
func (s Server) SubRouter(prefix string) *echo.Group {
	return s.driver.Group(prefix)
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
