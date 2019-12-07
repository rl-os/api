package pkg

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/gookit/config/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// Db connection
var Db *sqlx.DB

// InitializeDB database connection
func InitializeDB() {
	Db, err := sqlx.Connect(config.String("server.database.driver"), config.String("server.database.dsn"))
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("initialize database")
	}

	if err = Db.Ping(); err != nil {
		log.Fatal().
			Err(err).
			Msg("database ping")
	}

	stats := Db.Stats()

	log.Info().
		Str("driver", config.String("server.database.driver")).
		Int("open_connections", stats.OpenConnections).
		Int("max_open_connections", stats.MaxOpenConnections).
		Int("idle", stats.Idle).
		Int("in_use", stats.InUse).
		Msg("database connected")
}
