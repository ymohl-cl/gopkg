package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ymohl-cl/gopkg/httput"
)

func TestPing(t *testing.T) {
	// default: shoudl be ok
	func() {
		// init
		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		context := httput.NewContext(req)

		// assert
		if assert.NoError(t, Ping(context.Input)) {
			assert.Equal(t, http.StatusOK, context.Rec.Code)
			assert.Equal(t, `{"pong":true}`, context.Rec.Body.String())
		}
	}()
}
