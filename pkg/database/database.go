package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// Initialize database connection
func Initialize() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Initialize database error")
	}

	if err = db.Ping(); err != nil {
		log.Fatal().
			Err(err).
			Msg("Database connecting fail")
	}

	return db, err
}
