package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// default: should be ok
	func() {
		// init
		expectedMessage := "Hello world !"
		message := Hello("world")

		// assert
		assert.Equal(t, expectedMessage, message)
	}()
}
