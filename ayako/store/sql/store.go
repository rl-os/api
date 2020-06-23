package sql

import (
	osu "github.com/deissh/osu-go-client"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/jmoiron/sqlx"
)

type SqlStore interface {
	GetMaster() *sqlx.DB

	GetOsuClient() *osu.OsuAPI

	Beatmap() store.Beatmap
	BeatmapSet() store.BeatmapSet
	User() store.User
}
