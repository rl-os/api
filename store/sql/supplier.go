package sql

import (
	osu "github.com/deissh/osu-go-client"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rl-os/api/config"
	"github.com/rl-os/api/services"
	"github.com/rl-os/api/store"
	"github.com/rl-os/api/store/layers"
	"github.com/rs/zerolog/log"
)

type SupplierStores struct {
	beatmap    store.Beatmap
	beatmapSet store.BeatmapSet
	user       store.User
	oauth      store.OAuth
	friend     store.Friend
	chat       store.Chat
}

type Supplier struct {
	master *sqlx.DB
	cfg    *config.Config

	osuClient *osu.OsuAPI

	stores SupplierStores
}

// Init new store
// Using with DI
func Init(cfg *config.Config, services *services.Services) store.Store {
	log.Debug().Msg("Creating new SQL store")
	supplier := &Supplier{
		cfg:       cfg,
		osuClient: services.Bancho,
	}

	supplier.initConnection()

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
	supplier.stores.friend = layers.NewFriendWithLog(
		newSqlFriendStore(supplier),
	)
	supplier.stores.chat = layers.NewChatWithLog(
		newSqlChatStore(supplier),
	)

	return supplier
}

func (ss *Supplier) initConnection() {
	conn, err := sqlx.Connect(ss.cfg.Database.Driver, ss.cfg.Database.DSN)
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
		Str("driver", ss.cfg.Database.Driver).
		Int("open_connections", stats.OpenConnections).
		Int("max_open_connections", stats.MaxOpenConnections).
		Int("idle", stats.Idle).
		Int("in_use", stats.InUse).
		Msg("master database connected")
}

func (ss *Supplier) GetMaster() *sqlx.DB       { return ss.master }
func (ss *Supplier) GetConfig() *config.Config { return ss.cfg }

func (ss *Supplier) GetOsuClient() *osu.OsuAPI { return ss.osuClient }

func (ss *Supplier) Close() { _ = ss.master.Close() }

func (ss *Supplier) Beatmap() store.Beatmap       { return ss.stores.beatmap }
func (ss *Supplier) BeatmapSet() store.BeatmapSet { return ss.stores.beatmapSet }
func (ss *Supplier) User() store.User             { return ss.stores.user }
func (ss *Supplier) OAuth() store.OAuth           { return ss.stores.oauth }
func (ss *Supplier) Friend() store.Friend         { return ss.stores.friend }
func (ss *Supplier) Chat() store.Chat             { return ss.stores.chat }
