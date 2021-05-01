package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"time"
)

var ProviderSet = wire.NewSet(NewOptions, New)

type Options struct {
	URL       string
	MaxIdle   int
	MaxActive int
}

type Pool struct {
	*redis.Pool

	logger *zerolog.Logger
}

// NewOptions create and parse config from viper instance
func NewOptions(logger *zerolog.Logger, v *viper.Viper) (*Options, error) {
	o := Options{}

	logger.Debug().
		Msg("Loading config file")
	if err := v.UnmarshalKey("redis", &o); err != nil {
		return nil, err
	}

	logger.Debug().
		Interface("redis", o).
		Msg("Loaded config")

	return &o, nil
}

func New(logger *zerolog.Logger, o *Options) *Pool {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(o.URL)
		},
		MaxIdle:     o.MaxIdle,
		MaxActive:   o.MaxActive,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	stats := pool.Stats()
	logger.Debug().
		Int("total_connections", stats.ActiveCount).
		Int("idle_connections", stats.IdleCount).
		Dur("timeouts", stats.WaitDuration).
		Msg("redis connected")

	return &Pool{
		pool,
		logger,
	}
}
