package sql

import (
	"github.com/deissh/osu-lazer/ayako/store"
)

type SqlStore interface {
	Beatmap() store.Beatmap
}
