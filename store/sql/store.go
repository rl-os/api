package sql

import (
	osu "github.com/deissh/osu-go-client"
	"github.com/jmoiron/sqlx"
	"github.com/rl-os/api/config"
	"github.com/rl-os/api/store"
)

type SqlStore interface {
	GetMaster() *sqlx.DB
	GetConfig() *config.Config

	GetOsuClient() *osu.OsuAPI

	OAuth() store.OAuth
	Beatmap() store.Beatmap
	BeatmapSet() store.BeatmapSet
	User() store.User
	Friend() store.Friend
	Chat() store.Chat

	Close()
}
