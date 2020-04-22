package mock_store

import (
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/golang/mock/gomock"
)

type MockStore struct {
	ctrl *gomock.Controller

	beatmap    *MockBeatmap
	beatmapSet *MockBeatmapSet
}

func NewMockStore(ctrl *gomock.Controller) MockStore {
	return MockStore{
		ctrl:       ctrl,
		beatmap:    NewMockBeatmap(ctrl),
		beatmapSet: NewMockBeatmapSet(ctrl),
	}
}

func (ss MockStore) Beatmap() store.Beatmap       { return ss.beatmap }
func (ss MockStore) BeatmapSet() store.BeatmapSet { return ss.beatmapSet }

func (ss MockStore) BeatmapExpect() *MockBeatmapMockRecorder       { return ss.beatmap.EXPECT() }
func (ss MockStore) BeatmapSetExpect() *MockBeatmapSetMockRecorder { return ss.beatmapSet.EXPECT() }
