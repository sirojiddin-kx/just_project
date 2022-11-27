package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PositionServiceHost string
	PositionServicePort int

	CompanyServiceHost string
	CompanyServicePort int

	MinioAccessKeyID string
	MinioSecretKey   string
	MinioEndpoint    string
	MinioBucketName  string
	MinioLocation    string
	MinioHost        string

	LogLevel string
	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))

	config.PositionServiceHost = cast.ToString(getOrReturnDefault("POSITION_SERVICE_HOST", "localhost"))
	config.PositionServicePort = cast.ToInt(getOrReturnDefault("POSITION_SERVICE_PORT", 5004))

	config.CompanyServiceHost = cast.ToString(getOrReturnDefault("COMPANY_SERVICE_HOST", "localhost"))
	config.CompanyServicePort = cast.ToInt(getOrReturnDefault("COMPANY_SERVICE_PORT", 5005))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
