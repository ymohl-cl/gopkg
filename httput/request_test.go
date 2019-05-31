package httput

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	path = "/path_to_"
)

func TestRequestJSON(t *testing.T) {
	// Default: should be ok
	func() {
		var request *http.Request
		var err error
		p := map[string]interface{}{
			"Name": "Toto",
		}

		request, err = RequestJSON(http.MethodPost, "/path", &p)
		if assert.NoError(t, err) {
			assert.NotNil(t, request)
		}
	}()

	// Should return an error because payload isn't marshalable
	func() {
		var request *http.Request
		var err error
		p := map[string]interface{}{
			"Name": make(chan int),
		}

		request, err = RequestJSON(http.MethodPost, "/path", p)
		if assert.Error(t, err) {
			assert.Nil(t, request)
		}
	}()
}
