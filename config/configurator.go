package config

import (
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

// Init initialize a Config
func Init(configPath string) *Config {
	config := &Config{}

	if config.AutoReload {
		config.AutoReloadInterval = time.Second * 10
		log.Info().
			Dur("auto_reload_interval", config.AutoReloadInterval).
			Msg("enabled auto reload")
	}

	if err := config.Load(configPath); err != nil {
		log.Panic().
			Err(err).
			Msgf("Failed to load configuration from %v", configPath)
	}

	return config
}

// GetEnvironment get environment
func (c *Config) GetEnvironment() string {
	if env := os.Getenv("ENV"); env == "" {
		return "development"
	} else {
		return env
	}
}

// GetErrorOnUnmatchedKeys returns a boolean indicating if an errors should be
// thrown if there are keys in the config file that do not correspond to the
// config struct
func (c *Config) GetErrorOnUnmatchedKeys() bool {
	return c.ErrorOnUnmatchedKeys
}
