package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"os"
)

var ProviderSet = wire.NewSet(New)

func New(path string) (*viper.Viper, error) {
	var (
		err       error
		v         = viper.New()
		envPrefix = os.Getenv("ENV_PREFIX")
	)

	v.AddConfigPath(".")
	v.SetEnvPrefix(envPrefix)
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	fmt.Printf("loaded config file -> %s\n", v.ConfigFileUsed())

	return v, err
}
