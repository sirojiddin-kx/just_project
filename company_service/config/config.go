package config

import (
	"os"

	"github.com/spf13/cast"
)

var PerPageSize = 100

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PostgresHost     string
	PostgresPort     int
	PostgresPassword string
	PostgresUser     string
	PostgresDB       string
	LogLevel         string

	UserServiceHost string
	UserServicePort int

	ProducerServiceHost string
	ProducerServicePort int

	RPCPort string

	PasscodePool   string
	PasscodeLength int
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "company_service"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "proxima"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "proxima_ops"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":5005"))


	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

