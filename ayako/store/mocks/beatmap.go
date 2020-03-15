package mocks

import (
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/stretchr/testify/mock"
)

type BeatmapStore struct {
	mock.Mock
}

// Get provides a mocks function with given fields: user_id, offset, limit
func (_m *BeatmapStore) GetBeatmap(id uint) (*entity.Beatmap, error) {
	ret := _m.Called(id)

	var r0 *entity.Beatmap
	if rf, ok := ret.Get(0).(func(uint) *entity.Beatmap); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Beatmap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

func (_m *BeatmapStore) CreateBeatmap(from interface{}) (*entity.Beatmap, error) {
	ret := _m.Called(from)

	var r0 *entity.Beatmap
	if rf, ok := ret.Get(0).(func(interface{}) *entity.Beatmap); ok {
		r0 = rf(from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Beatmap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(from)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

func (_m *BeatmapStore) UpdateBeatmap(id uint, from interface{}) (*entity.Beatmap, error) {
	panic("implement me")
}

func (_m *BeatmapStore) DeleteBeatmap(id uint) error {
	panic("implement me")
}