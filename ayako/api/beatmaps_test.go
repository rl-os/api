package api

import (
	"github.com/deissh/osu-lazer/ayako/entity"
	mock_store "github.com/deissh/osu-lazer/ayako/store/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBeatmapHandlers_Show(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(s *mock_store.MockedStore)
		wantErr bool
	}{
		{
			name: "Get beatmap by valid id",
			args: args{"112233"},
			prepare: func(s *mock_store.MockedStore) {
				s.BeatmapExpect().
					Get(gomock.Any(), gomock.Eq(uint(112233))).
					Return(&entity.SingleBeatmap{}, nil)
			},
		},
		{
			name:    "Get beatmap by invalid id",
			args:    args{"notid"},
			prepare: func(s *mock_store.MockedStore) {},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			e := echo.New()
			mocked := mock_store.InitStore(ctrl)

			h := &BeatmapHandlers{mocked}

			tt.prepare(&mocked)

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id)
			c.Set("current_user_id", uint(103))

			if err := h.Show(c); (err != nil) != tt.wantErr {
				t.Errorf("Show() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
