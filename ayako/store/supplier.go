package store

import (
	"github.com/gookit/config/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type SupplierStores struct {
	beatmap    Beatmap
	beatmapSet BeatmapSet
}

type Supplier struct {
	master *sqlx.DB

	stores SupplierStores
}

// Init new store
// Using with DI
func Init() Store {
	supplier := &Supplier{}

	supplier.initConnection()

	return &supplier.stores
}

func (ss *Supplier) initConnection() {
	conn, err := sqlx.Connect(config.String("server.database.driver"), config.String("server.database.dsn"))
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("initialize database")
	}

	if err = conn.Ping(); err != nil {
		log.Fatal().
			Err(err).
			Msg("database ping")
	}

	ss.master = conn

	stats := ss.master.Stats()
	log.Info().
		Str("driver", config.String("server.database.driver")).
		Int("open_connections", stats.OpenConnections).
		Int("max_open_connections", stats.MaxOpenConnections).
		Int("idle", stats.Idle).
		Int("in_use", stats.InUse).
		Msg("master database connected")
}

func (s *SupplierStores) Beatmap() Beatmap       { return s.beatmap }
func (s *SupplierStores) BeatmapSet() BeatmapSet { return s.beatmapSet }
