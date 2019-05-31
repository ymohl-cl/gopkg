package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// should return an error because config can't be loaded
	func() {
		// init
		appName := "WRONG_APP_NAME"
		expectedError := "required key WRONG_APP_NAME_SSL_ENABLE missing value"
		s, err := New(appName)

		// assert
		if assert.Error(t, err) {
			assert.Zero(t, s)
			assert.Equal(t, expectedError, err.Error())
		}
	}()

	// default: should be ok with tls configuration
	func() {
		// init
		appName := "Server"
		assert.NoError(t, setenv())
		s, err := New(appName)
		assert.NoError(t, unsetenv())

		// assert
		if assert.NotNil(t, s) && assert.NoError(t, err) {
			assert.Len(t, s.driver.Routes(), 1)
			assert.Equal(t, "GET", (s.driver.Routes())[0].Method)
			assert.Equal(t, "/ping", (s.driver.Routes())[0].Path)
		}
	}()

	// default: should be ok without tls configuration
	func() {
		// init
		appName := "Server"
		envs := map[string]string{"SERVER_PORT": testPort, "SERVER_SSL_ENABLE": "false"}
		defaultEnv = envs
		assert.NoError(t, setenv())
		s, err := New(appName)
		assert.NoError(t, unsetenv())
		defaultEnv = privateEnv

		// assert
		if assert.NotNil(t, s) && assert.NoError(t, err) {
			assert.Len(t, s.driver.Routes(), 1)
			assert.Equal(t, "GET", (s.driver.Routes())[0].Method)
			assert.Equal(t, "/ping", (s.driver.Routes())[0].Path)
		}
	}()
}
