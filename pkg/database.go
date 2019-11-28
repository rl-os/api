package pkg

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/config/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// Current database connection
var Connection *sqlx.DB

// InitializeDB database connection
func InitializeDB() {
	db, err := sqlx.Connect(config.String("server.database.driver"), config.String("server.database.dsn"))
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("initialize database")
	}

	if err = db.Ping(); err != nil {
		log.Fatal().
			Err(err).
			Msg("database ping")
	}

	log.Debug().
		Msg("database connected")
}
