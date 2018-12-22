package server

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPort    = "4242"
	testSSLCert = "certificates/cert.pem"
	testSSLKey  = "certificates/key.pem"
)

var defaultEnv = map[string]string{
	"SERVER_PORT":            testPort,
	"SERVER_SSL_CERTIFICATE": testSSLCert,
	"SERVER_SSL_KEY":         testSSLKey,
}

func setenv() error {
	var err error
	for key, value := range defaultEnv {
		if err = os.Setenv(key, value); err != nil {
			return err
		}
	}
	return nil
}

func unsetenv() error {
	var err error
	for value := range defaultEnv {
		if err = os.Unsetenv(value); err != nil {
			return err
		}
	}
	return nil
}

func TestNewConfig(t *testing.T) {
	// should return an error because envconfig Process failed
	func() {
		// init
		appName := "WRONG_APP_NAME"
		expectedError := "required key WRONG_APP_NAME_SSL_CERTIFICATE missing value"
		c, err := NewConfig(appName)

		// assert
		if assert.Error(t, err) {
			assert.Zero(t, c)
			assert.Equal(t, expectedError, err.Error())
		}
	}()

	// default: should be ok
	func() {
		// init
		appName := "Server"
		assert.NoError(t, setenv())
		c, err := NewConfig(appName)
		assert.NoError(t, unsetenv())

		// assert
		if assert.NotNil(t, c) && assert.NoError(t, err) {
			assert.Equal(t, testPort, c.Port)
			assert.Equal(t, testSSLCert, c.SSL.Certificate)
			assert.Equal(t, testSSLKey, c.SSL.Key)
		}
	}()
}
