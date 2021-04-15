package mock_repository

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMockStore(t *testing.T) {
	t.Run("get beatmap by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mocked := NewMockBeatmap(ctrl)
		mocked.EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(&entity.SingleBeatmap{Beatmap: entity.Beatmap{ID: 123321}}, nil)

		var s repository.Beatmap
		s = mocked

		data, err := s.Get(context.TODO(), 123)
		assert.Nil(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, int64(123321), data.ID)
	})

	t.Run("get beatmap by id with errors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		defError := errors.New("some internal errors")

		mocked := NewMockBeatmap(ctrl)
		mocked.EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(nil, defError)

		var s repository.Beatmap
		s = mocked

		data, err := s.Get(context.TODO(), 1)
		assert.Equal(t, err, defError)
		assert.Nil(t, data)
	})
}
