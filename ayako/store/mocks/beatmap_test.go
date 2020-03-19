package mocks

import (
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestStore_Beatmap(t *testing.T) {
	t.Run("Testing Beatmap mock", func(t *testing.T) {
		testCreateBeatmap(t)
	})
}

func testCreateBeatmap(t *testing.T) {
	type args struct {
		from interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Beatmap
		wantErr bool
	}{
		{
			name:    "Creating new Beatmap",
			args:    args{from: entity.Beatmap{Convert: true}},
			want:    &entity.Beatmap{Convert: true},
			wantErr: false,
		},
		{
			name:    "Creating new Beatmap with error",
			args:    args{from: entity.Beatmap{Convert: true}},
			want:    &entity.Beatmap{Convert: true},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_m := &BeatmapStore{mock.Mock{}}

			var _e error
			if tt.wantErr {
				_e = assert.AnError
			} else {
				_e = nil
			}

			_m.On("CreateBeatmap", tt.args.from).Return(tt.want, _e)

			got, err := _m.CreateBeatmap(tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBeatmap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBeatmap() got = %v, want %v", got, tt.want)
			}
		})
	}
}
