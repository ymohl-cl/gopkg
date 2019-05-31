package errorx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	code := uint64(42)
	message := "error example"

	// default: should be ok
	func() {
		// init
		err := New(42, message)

		// assert
		if assert.NotNil(t, err) {
			assert.Equal(t, message, err.message)
			assert.Equal(t, code, err.code)
		}
	}()
}

func TestError(t *testing.T) {
	message := "error example"

	// default: should be ok
	func() {
		// init
		var err error
		err = New(uint64(0), message)

		// assert
		if assert.NotNil(t, err) {
			assert.Equal(t, message, err.Error())
		}
	}()
}

func TestWrap(t *testing.T) {
	code := uint64(0)
	message1 := "error example"
	message2 := "context error"

	// default: should be ok
	func() {
		// init
		err := New(code, message1)
		err.Wrap(message2)

		// assert
		assert.Equal(t, code, err.code)
		assert.Equal(t, message2+separator+message1, err.message)
	}()
}
