package pkg

import (
	"github.com/gookit/config/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/redis.v5"
)

var Rb *redis.Client

// InitializeDB database connection
func InitializeRedis() {
	conn := redis.NewClient(&redis.Options{
		Addr:     config.String("server.redis.host"),
		Password: config.String("server.redis.password"),
		DB:       config.Int("server.redis.db"),
	})

	if err := Rb.Ping().Err(); err != nil {
		log.Fatal().
			Err(err).
			Msg("redis ping")
	}

	Rb = conn
	stats := Rb.PoolStats()

	log.Info().
		Uint32("total_conns", stats.TotalConns).
		Uint32("free_conns", stats.FreeConns).
		Uint32("requests", stats.Requests).
		Uint32("timeouts", stats.Timeouts).
		Uint32("hits", stats.Hits).
		Msg("redis connected")
}
