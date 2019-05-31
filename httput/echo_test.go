package httput

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	// Default: should be ok
	func() {
		var c Context
		var request *http.Request
		var err error
		payload := map[string]interface{}{
			"Name": "Toto",
		}

		if request, err = RequestJSON(http.MethodPost, "/path", &payload); err != nil {
			t.Error(err)
		}
		c = NewContext(request)
		assert.NotNil(t, c.Input)
		assert.NotNil(t, c.Rec)
	}()
}
