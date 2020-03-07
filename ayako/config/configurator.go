package config

import (
	"os"
	"regexp"
	"time"
)

// Init initialize a Config
func Init() *Config {
	config := &Config{}

	if config.AutoReload && config.AutoReloadInterval == 0 {
		config.AutoReloadInterval = time.Second
	}

	return config
}

var testRegexp = regexp.MustCompile("_test|(\\.test$)")

// GetEnvironment get environment
func (c *Config) GetEnvironment() string {
	if c.Environment == "" {
		if testRegexp.MatchString(os.Args[0]) {
			return "test"
		}

		return "development"
	}
	return c.Environment
}

// GetErrorOnUnmatchedKeys returns a boolean indicating if an error should be
// thrown if there are keys in the config file that do not correspond to the
// config struct
func (c *Config) GetErrorOnUnmatchedKeys() bool {
	return c.ErrorOnUnmatchedKeys
}
