package app

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/rl-os/api/entity"
	mock_repository "github.com/rl-os/api/repository/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBeatmapUseCase_GetBeatmap(t *testing.T) {
	t.SkipNow()

	t.Run("get beatmap by invalid id", func(t *testing.T) {
		ctx := context.TODO()
		ctrl := gomock.NewController(t)

		mock := mock_repository.NewMockBeatmap(ctrl)

		mock.EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("not exist"))

		a := NewBeatmapUseCase(nil, mock)

		_, err := a.Get(ctx, 1000)
		if err == nil {
			t.Error("must be error")
			return
		}

		assert.Contains(t, err.Error(), ErrNotFoundBM.Msg)
	})

	t.Run("get beatmap by valid id", func(t *testing.T) {
		ctx := context.TODO()
		ctrl := gomock.NewController(t)

		mock := mock_repository.NewMockBeatmap(ctrl)

		bm := &entity.SingleBeatmap{}
		bm.ID = 1000

		mock.EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(bm, nil)

		a := NewBeatmapUseCase(nil, mock)

		data, err := a.Get(ctx, 1000)

		assert.NoError(t, err)
		assert.Equal(t, data, bm)
	})
}

func TestBeatmapUseCase_LookupBeatmap(t *testing.T) {
	t.Run("lookup beatmap by invalid checksum", func(t *testing.T) {
		ctx := context.TODO()

		a := NewBeatmapUseCase(nil, nil)

		_, err := a.Lookup(ctx, 1000, "adsasd", "")
		if err == nil {
			t.Error("must be error")
			return
		}

		assert.Contains(t, err.Error(), ErrNotFoundBM.Msg)
	})

	t.Run("lookup beatmap by invalid filename", func(t *testing.T) {
		ctx := context.TODO()

		a := NewBeatmapUseCase(nil, nil)

		_, err := a.Lookup(ctx, 0, "", "adsasd")
		if err == nil {
			t.Error("must be error")
			return
		}

		assert.Contains(t, err.Error(), ErrNotFoundBM.Msg)
	})

	t.Run("lookup beatmap by invalid id", func(t *testing.T) {
		ctx := context.TODO()
		ctrl := gomock.NewController(t)

		mock := mock_repository.NewMockBeatmap(ctrl)

		mock.EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("not exist"))

		a := NewBeatmapUseCase(nil, mock)

		_, err := a.Lookup(ctx, 1000, "", "")
		if err == nil {
			t.Error("must be error")
			return
		}

		assert.Contains(t, err.Error(), ErrNotFoundBM.Msg)
	})

	t.Run("lookup beatmap by valid id", func(t *testing.T) {
		ctx := context.TODO()
		ctrl := gomock.NewController(t)

		mock := mock_repository.NewMockBeatmap(ctrl)

		bm := &entity.SingleBeatmap{}
		bm.ID = 1000

		mock.EXPECT().
			Get(gomock.Any(), gomock.Any()).
			Return(bm, nil)

		a := NewBeatmapUseCase(nil, mock)

		data, err := a.Lookup(ctx, 1000, "", "")

		assert.NoError(t, err)
		assert.Equal(t, data, bm)
	})
}
