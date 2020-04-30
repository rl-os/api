package layers

import "github.com/deissh/osu-lazer/ayako/store"

type storeWithLog struct {
	beatmap    store.Beatmap
	beatmapSet store.BeatmapSet
}

func NewStoreWithLog(original store.Store) store.Store {
	return &storeWithLog{
		beatmap:    BeatmapWithLog{original.Beatmap()},
		beatmapSet: BeatmapSetWithLog{original.BeatmapSet()},
	}
}

func (ss *storeWithLog) Beatmap() store.Beatmap       { return ss.beatmap }
func (ss *storeWithLog) BeatmapSet() store.BeatmapSet { return ss.beatmapSet }
