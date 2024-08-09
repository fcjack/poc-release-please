package config

import (
	"fmt"
	"log/slog"
	"os"
)

type Config struct {
	// LogLevel is the log level.
	LogLevel string

	// Env is the environment.
	Env string

	// DatabaseUser is the database user.
	DatabaseUser string

	// DatabasePassword is the database password.
	DatabasePassword string

	// DatabaseHost is the database host.
	DatabaseHost string

	// DatabasePort is the database port.
	DatabasePort string

	// DatabaseName is the database name.
	DatabaseName string

	// DatabaseCryptKey is the database crypt key.
	DatabaseCryptKey string
}

func ConfigFromEnv() (Config, error) {
	c := Config{
		LogLevel:         os.Getenv("LOG_LEVEL"),
		Env:              os.Getenv("ENV"),
		DatabaseUser:     os.Getenv("MYSQL_USER"),
		DatabasePassword: os.Getenv("MYSQL_PASSWORD"),
		DatabaseHost:     os.Getenv("DB_HOST"),
		DatabasePort:     os.Getenv("DB_PORT"),
		DatabaseName:     os.Getenv("DB_NAME"),
		DatabaseCryptKey: os.Getenv("DB_CRYPT_KEY"),
	}

	if err := c.Parse(); err != nil {
		return Config{}, err
	}

	return c, nil
}

// Parse parses the config and returns an error if anything is missing or invalid.
func (c *Config) Parse() error {
	if c.Env == "" || (c.Env != "dev" && c.Env != "ops" && c.Env != "prod") {
		return fmt.Errorf("ENV must be either dev or prod (not %q)", c.Env)
	}

	if c.DatabaseHost == "" {
		return fmt.Errorf("env DB_HOST must be set")
	}

	if c.DatabasePort == "" {
		return fmt.Errorf("env DB_PORT must be set")
	}

	if c.DatabaseUser == "" {
		return fmt.Errorf("env MYSQL_USER must be set")
	}

	if c.DatabasePassword == "" {
		return fmt.Errorf("env MYSQL_PASSWORD must be set")
	}

	if c.DatabaseName == "" {
		return fmt.Errorf("env DB_NAME must be set")
	}

	if c.DatabaseCryptKey == "" {
		return fmt.Errorf("env DB_CRYPT_KEY must be set")
	}

	return nil
}

func (c *Config) ParseLogLevel() slog.Level {
	switch c.LogLevel {
	case "debug", "DEBUG":
		return slog.LevelDebug
	case "info", "INFO":
		return slog.LevelInfo
	case "warn", "WARN":
		return slog.LevelWarn
	case "error", "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
