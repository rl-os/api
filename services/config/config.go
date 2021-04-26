package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"os"
)

var ProviderSet = wire.NewSet(New)

// New viper instance with preload config
// from file if path is not empty
// also load from env vars with ENV_PREFIX
func New(path string) (*viper.Viper, error) {
	var (
		err       error
		v         = viper.New()
		envPrefix = os.Getenv("ENV_PREFIX")
	)

	v.AddConfigPath(".")
	v.SetEnvPrefix(envPrefix)

	v.AutomaticEnv()

	if path != "" {
		v.SetConfigFile(path)

		err := v.ReadInConfig()
		if err != nil {
			return nil, err
		}
	}

	fmt.Printf("loaded config file -> %s\n", v.ConfigFileUsed())

	return v, err
}
