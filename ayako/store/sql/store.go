package sql

import (
	osu "github.com/deissh/osu-go-client"
	"github.com/deissh/rl/ayako/config"
	"github.com/deissh/rl/ayako/store"
	"github.com/jmoiron/sqlx"
)

type SqlStore interface {
	GetMaster() *sqlx.DB
	GetConfig() *config.Config

	GetOsuClient() *osu.OsuAPI

	OAuth() store.OAuth
	Beatmap() store.Beatmap
	BeatmapSet() store.BeatmapSet
	User() store.User
}
