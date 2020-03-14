package sql

import (
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/jmoiron/sqlx"
)

type SqlStore interface {
	GetMaster() *sqlx.DB

	Beatmap() store.Beatmap
}
