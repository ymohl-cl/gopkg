package httput

import (
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
)

// Context take an echo context and a httptest.Record to handlers tests
type Context struct {
	Input echo.Context
	Rec   *httptest.ResponseRecorder
}

// NewContext return a context to handlers tests
func NewContext(req *http.Request) Context {
	var c Context
	var e *echo.Echo

	e = echo.New()
	c.Rec = httptest.NewRecorder()
	c.Input = e.NewContext(req, c.Rec)
	return c
}
