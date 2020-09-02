package mock_store

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMockStore(t *testing.T) {
	t.Run("get beatmap by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mocked := InitStore(ctrl)
		mocked.BeatmapExpect().
			Get(gomock.Any(), gomock.Any()).
			Return(&entity.SingleBeatmap{Beatmap: entity.Beatmap{ID: 123321}}, nil)

		var s store.Store
		s = mocked

		data, err := s.Beatmap().Get(context.TODO(), 123)
		assert.Nil(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, int64(123321), data.ID)
	})

	t.Run("get beatmap by id with errors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		defError := errors.New("some internal errors")

		mocked := InitStore(ctrl)
		mocked.BeatmapExpect().
			Get(gomock.Any(), gomock.Any()).
			Return(nil, defError)

		var s store.Store
		s = mocked

		data, err := s.Beatmap().Get(context.TODO(), 1)
		assert.Equal(t, err, defError)
		assert.Nil(t, data)
	})
}
