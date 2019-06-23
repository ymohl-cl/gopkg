package config

import "github.com/kelseyhightower/envconfig"

// Default driver name
const (
	DefaultDriverName = "postgres"
)

// DriverName to postgres sql
var DriverName = DefaultDriverName

// Config data to create a new postgres instance.
type Config struct {
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	DBName   string `envconfig:"POSTGRES_NAME" required:"true"`
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     string `envconfig:"POSTGRES_PORT" required:"true"`
}

// New parse the environment values to return a initialized configuration
func New(appName string) (Config, error) {
	var err error
	var c Config

	if err = envconfig.Process(appName, &c); err != nil {
		return Config{}, err
	}
	return c, nil
}
