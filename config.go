package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config defines the structure of the configuration for the application.
type Config struct {
	Server struct {
		Port string `envconfig:"SERVER_PORT" default:"8080"`
		Host string `envconfig:"SERVER_HOST" default:"localhost"`
	}
	Database struct {
		URI      string `envconfig:"MONGODB_URI" required:"true"`
		DBName   string `envconfig:"MONGODB_DB_NAME" required:"true"`
		Username string `envconfig:"DB_USERNAME"`
		Password string `envconfig:"DB_PASSWORD"`
	}
	APIKeys map[string]string `envconfig:"API_KEYS"`
}

// LoadConfig reads configuration from environment variables and returns a Config struct.
func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// MustLoadConfig is similar to LoadConfig but logs a fatal error and exits if the configuration cannot be loaded.
func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	return cfg
}

func main() {
	// Load the configuration when the application starts.
	config := MustLoadConfig()

	// Use the loaded configuration.
	log.Printf("Server will start at %s:%s", config.Server.Host, config.Server.Port)
}
