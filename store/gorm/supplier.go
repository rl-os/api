package sql

import (
	osu "github.com/deissh/osu-go-client"
	"github.com/jmoiron/sqlx"
	"github.com/rl-os/api/config"
	"github.com/rl-os/api/store"
	"github.com/rl-os/api/store/layers"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	master *gorm.DB
	cfg    *config.Config

	osuClient *osu.OsuAPI

	stores SupplierStores
}

// Init new store
// Using with DI
func Init(cfg *config.Config) store.Store {
	log.Debug().Msg("Creating new SQL store")
	supplier := &Supplier{
		cfg: cfg,
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
			Str("dsn", ss.cfg.Database.DSN).
			Msg("initialize database with ")
	}

	if err = conn.Ping(); err != nil {
		log.Fatal().
			Err(err).
			Msg("database ping")
	}

	stats := conn.Stats()
	log.Info().
		Str("driver", ss.cfg.Database.Driver).
		Int("open_connections", stats.OpenConnections).
		Int("max_open_connections", stats.MaxOpenConnections).
		Int("idle", stats.Idle).
		Int("in_use", stats.InUse).
		Msg("master database connected")

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn.DB,
	}), &gorm.Config{
		Logger: logger.New(&log.Logger, logger.Config{
			LogLevel: logger.Info,
		}),
	})

	ss.master = db
}

func (ss *Supplier) GetMaster() *gorm.DB       { return ss.master }
func (ss *Supplier) GetConfig() *config.Config { return ss.cfg }

func (ss *Supplier) GetOsuClient() *osu.OsuAPI { return ss.osuClient }

func (ss *Supplier) Close() { log.Warn().Msg("TODO") }

func (ss *Supplier) Beatmap() store.Beatmap       { return ss.stores.beatmap }
func (ss *Supplier) BeatmapSet() store.BeatmapSet { return ss.stores.beatmapSet }
func (ss *Supplier) User() store.User             { return ss.stores.user }
func (ss *Supplier) OAuth() store.OAuth           { return ss.stores.oauth }
func (ss *Supplier) Friend() store.Friend         { return ss.stores.friend }
func (ss *Supplier) Chat() store.Chat             { return ss.stores.chat }
