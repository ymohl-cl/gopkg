package server

import "github.com/kelseyhightower/envconfig"

type ssl struct {
	Certificate string `split_words:"true" required:"true"`
	Key         string `split_words:"true" required:"true"`
}

// Config support to the server
type Config struct {
	SSL  ssl    `split_words:"true" required:"true"`
	Port string `split_words:"true" required:"true"`
}

// NewConfig parse the environment values to return a initialized configuration
func NewConfig(appName string) (Config, error) {
	var err error
	var c Config

	if err = envconfig.Process(appName, &c); err != nil {
		return Config{}, err
	}
	return c, nil
}
