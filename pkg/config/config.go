package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ProviderSet = wire.NewSet(New)

func New(path string) (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)
	v.AddConfigPath(".")
	v.SetEnvPrefix("RL")
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	fmt.Printf("loaded config file -> %s\n", v.ConfigFileUsed())

	return v, err
}
