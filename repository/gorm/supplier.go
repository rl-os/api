package gorm

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Options struct {
	DSN    string `default:"postgres://postgres:postgres@/osuserver?sslmode=disable"`
	Driver string `default:"postgres"`
}

// NewOptions create and parse config from viper instance
func NewOptions(logger *zerolog.Logger, v *viper.Viper) (*Options, error) {
	o := Options{}

	logger.Debug().
		Msg("Loading config file")
	if err := v.UnmarshalKey("database", &o); err != nil {
		return nil, err
	}

	logger.Debug().
		Interface("database", o).
		Msg("Loaded config")

	return &o, nil
}

type Supplier struct {
	master *gorm.DB

	Logger  *zerolog.Logger
	Options *Options
}

// NewSupplier create new bd connection and return Supplier
func NewSupplier(logger *zerolog.Logger, options *Options) (*Supplier, error) {
	log := logger.With().
		Str("type", "repository.gorm").
		Logger()

	log.Debug().Msg("Creating new SQL repository")
	supplier := &Supplier{
		Logger:  &log,
		Options: options,
	}

	conn, err := supplier.initMasterConn()
	if err != nil {
		return nil, err
	}

	supplier.master = conn

	return supplier, nil
}

func (ss *Supplier) initMasterConn() (*gorm.DB, error) {
	conn, err := sqlx.Connect(ss.Options.Driver, ss.Options.DSN)
	if err != nil {
		ss.Logger.Fatal().
			Err(err).
			Str("dsn", ss.Options.DSN).
			Msg("initialize database with ")
	}

	if err = conn.Ping(); err != nil {
		ss.Logger.Fatal().
			Err(err).
			Msg("database ping")
	}

	stats := conn.Stats()
	ss.Logger.Info().
		Str("driver", ss.Options.Driver).
		Int("open_connections", stats.OpenConnections).
		Int("max_open_connections", stats.MaxOpenConnections).
		Int("idle", stats.Idle).
		Int("in_use", stats.InUse).
		Msg("master database connected")

	log := ss.Logger.With().
		Str("type", "gorm.debug").
		Str("connection", "master").
		Logger()

	return gorm.Open(postgres.New(postgres.Config{Conn: conn.DB}), &gorm.Config{
		Logger: logger.New(&log, logger.Config{
			LogLevel: logger.Info,
		}),
	})
}

func (ss *Supplier) GetMaster() *gorm.DB { return ss.master }
