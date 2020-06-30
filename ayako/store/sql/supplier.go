package sql

import (
	osu "github.com/deissh/osu-go-client"
	"github.com/deissh/rl/ayako/config"
	"github.com/deissh/rl/ayako/store"
	"github.com/deissh/rl/ayako/store/layers"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type SupplierStores struct {
	beatmap    store.Beatmap
	beatmapSet store.BeatmapSet
	user       store.User
	oauth      store.OAuth
}

type Supplier struct {
	master *sqlx.DB
	cfg    *config.Config

	osuClient *osu.OsuAPI

	stores SupplierStores
}

// Init new store
// Using with DI
func Init(cfg *config.Config, osuClient *osu.OsuAPI) store.Store {
	log.Debug().Msg("Creating new SQL store")
	supplier := &Supplier{
		cfg:       cfg,
		osuClient: osuClient,
	}

	supplier.initConnection(cfg)

	supplier.stores.beatmap = layers.NewBeatmapWithLog(
		newSqlBeatmapStore(supplier),
	)
	supplier.stores.beatmapSet = layers.NewBeatmapSetWithLog(
		newSqlBeatmapSetStore(supplier),
	)
	supplier.stores.user = layers.NewUserWithLog(
		newSqlUserStore(supplier),
	)
	supplier.stores.oauth = layers.NewOAuthWithLog(
		newSqlOAuthStore(supplier),
	)

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

func (ss *Supplier) GetMaster() *sqlx.DB       { return ss.master }
func (ss *Supplier) GetConfig() *config.Config { return ss.cfg }

func (ss *Supplier) GetOsuClient() *osu.OsuAPI { return ss.osuClient }

func (ss *Supplier) Beatmap() store.Beatmap       { return ss.stores.beatmap }
func (ss *Supplier) BeatmapSet() store.BeatmapSet { return ss.stores.beatmapSet }
func (ss *Supplier) User() store.User             { return ss.stores.user }
func (ss *Supplier) OAuth() store.OAuth           { return ss.stores.oauth }
