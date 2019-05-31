package httput

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
)

// RequestJSON build a httptest request with the specific payload
func RequestJSON(method, path string, payload interface{}) (*http.Request, error) {
	var b []byte
	var err error

	if b, err = json.Marshal(payload); err != nil {
		return nil, err
	}
	req := httptest.NewRequest(method, path, bytes.NewBuffer(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req, nil
}
