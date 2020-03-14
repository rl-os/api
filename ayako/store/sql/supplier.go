package sql

import (
	"github.com/deissh/osu-lazer/ayako/config"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type SupplierStores struct {
	beatmap    store.Beatmap
	beatmapSet store.BeatmapSet
}

type Supplier struct {
	master *sqlx.DB

	stores SupplierStores
}

// Init new store
// Using with DI
func Init(cfg *config.Config) store.Store {
	log.Debug().Msg("Creating new SQL store")
	supplier := &Supplier{}

	supplier.initConnection(cfg)

	supplier.stores.beatmap = newSqlBeatmapStore(supplier)
	supplier.stores.beatmapSet = newSqlBeatmapSetStore(supplier)

	return supplier
}

func (ss *Supplier) initConnection(cfg *config.Config) {
	conn, err := sqlx.Connect(cfg.Database.Driver, cfg.Database.DSN)
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
		Str("driver", cfg.Database.Driver).
		Int("open_connections", stats.OpenConnections).
		Int("max_open_connections", stats.MaxOpenConnections).
		Int("idle", stats.Idle).
		Int("in_use", stats.InUse).
		Msg("master database connected")
}

func (ss *Supplier) GetMaster() *sqlx.DB { return ss.master }

func (ss *Supplier) Beatmap() store.Beatmap       { return ss.stores.beatmap }
func (ss *Supplier) BeatmapSet() store.BeatmapSet { return ss.stores.beatmapSet }
