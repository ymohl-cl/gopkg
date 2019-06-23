package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymohl-cl/gopkg/server/model"
	"github.com/ymohl-cl/gopkg/server/transactor"
)

const (
	usernameKey = "username"
	passwordKey = "password"
)

// Handler manager to provide any driver using in process
type Handler struct {
	JWTKey     string
	transactor transactor.Transactor
	//	postgresDriver postgres.Postgres
	//	cassandra Cassandra
}

// New return a fresh handler
func New(appName, jwtKey string) (Handler, error) {
	var h Handler
	var err error

	h.JWTKey = jwtKey
	if h.transactor, err = transactor.New(appName); err != nil {
		return Handler{}, err
	}
	return h, nil
}

// Ping method http GET
func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &model.Pong{Pong: true})
}

// Login endpoint api
func (h Handler) Login(c echo.Context) error {
	var user model.User
	var token string
	var err error

	username := c.FormValue(usernameKey)
	password := c.FormValue(passwordKey)
	// find user
	if user, err = h.transactor.ConnectUser(username, password); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	if user.Identifier == "" {
		return echo.ErrUnauthorized
	}
	// token generation
	if token, err = h.generateToken(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

// NewUser endpoint api
func (h Handler) NewUser(c echo.Context) error {
	var user model.User
	var err error

	if err = c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err = c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if user, err = h.transactor.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user)
}
