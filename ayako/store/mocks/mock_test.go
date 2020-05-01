package mock_store

import (
	"errors"
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMockStore(t *testing.T) {
	t.Run("get beatmap by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mocked := InitStore(ctrl)
		mocked.BeatmapExpect().
			Get(gomock.Any()).
			Return(&entity.SingleBeatmap{Beatmap: entity.Beatmap{ID: 123321}}, nil)

		var s store.Store
		s = mocked

		data, err := s.Beatmap().Get(1)
		assert.Nil(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, int64(123321), data.ID)
	})

	t.Run("get beatmap by id with error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		defError := errors.New("some internal error")

		mocked := InitStore(ctrl)
		mocked.BeatmapExpect().
			Get(gomock.Any()).
			Return(nil, defError)

		var s store.Store
		s = mocked

		data, err := s.Beatmap().Get(1)
		assert.Equal(t, err, defError)
		assert.Nil(t, data)
	})
}
