package mocks

import (
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/stretchr/testify/mock"
)

type Store struct {
	mock.Mock
}


// Beatmap provides a mocks function with given fields:
func (_m *Store) Beatmap() store.Beatmap {
	ret := _m.Called()

	var r0 store.Beatmap
	if rf, ok := ret.Get(0).(func() store.Beatmap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.Beatmap)
		}
	}

	return r0
}
