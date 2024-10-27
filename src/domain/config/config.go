// config/config.go
package config

import (
	"github.com/spf13/viper" // Popular for config management
	"sabyabhoi.com/payment-settler/src/infrastructure/database"
)

type Config struct {
	// Server settings
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"` // dev, prod, etc

	// Database settings
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	// Auth settings
	JWTSecret string `mapstructure:"JWT_SECRET"`
	JWTExpiry int    `mapstructure:"JWT_EXPIRY"`

	// External services
	RedisURL  string `mapstructure:"REDIS_URL"`
	AwsRegion string `mapstructure:"AWS_REGION"`

	// API rate limiting
	RateLimit int `mapstructure:"RATE_LIMIT"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env") // or use yaml/json
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	return config, err
}

func (c *Config) ToDatabaseConfig() database.Config {
	return database.Config{
		Host:     c.DBHost,
		Port:     c.DBPort,
		User:     c.DBUser,
		Password: c.DBPassword,
		DBName:   c.DBName,
	}
}
