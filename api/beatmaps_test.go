package api

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/entity"
	mock_store "github.com/rl-os/api/store/mocks"
	"net/http"
	"net/http/httptest"
	"net/url"
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
		{
			name: "Get beatmap with error in store",
			args: args{"112233"},
			prepare: func(s *mock_store.MockedStore) {
				s.BeatmapExpect().
					Get(gomock.Any(), gomock.Eq(uint(112233))).
					Return(nil, errors.New("internal error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			e := echo.New()
			mocked := mock_store.InitStore(ctrl)

			// trying create new api
			fakeApp := app.NewApp(mocked, nil)

			h := &BeatmapHandlers{fakeApp}

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

func TestBeatmapHandlers_Lookup(t *testing.T) {
	type args struct {
		id       string
		checksum string
		filename string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(s *mock_store.MockedStore)
		wantErr bool
	}{
		{
			name: "Lookup beatmap by id",
			args: args{
				id: "112233",
			},
			prepare: func(s *mock_store.MockedStore) {
				s.BeatmapExpect().
					Get(gomock.Any(), gomock.Eq(uint(112233))).
					Return(&entity.SingleBeatmap{}, nil)
			},
		},
		{
			name: "Lookup beatmap by invalid id",
			args: args{
				id: "notuint",
			},
			wantErr: true,
			prepare: func(s *mock_store.MockedStore) {},
		},
		{
			name: "Lookup beatmap by id with error",
			args: args{
				id: "112233",
			},
			wantErr: true,
			prepare: func(s *mock_store.MockedStore) {
				s.BeatmapExpect().
					Get(gomock.Any(), gomock.Eq(uint(112233))).
					Return(nil, errors.New("internal error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			e := echo.New()
			mocked := mock_store.InitStore(ctrl)

			// trying create new api
			fakeApp := app.NewApp(mocked, nil)

			h := &BeatmapHandlers{fakeApp}

			tt.prepare(&mocked)

			q := make(url.Values)
			q.Set("id", tt.args.id)
			q.Set("checksum", tt.args.checksum)
			q.Set("filename", tt.args.filename)

			req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			c := e.NewContext(req, httptest.NewRecorder())
			c.Set("current_user_id", uint(103))

			if err := h.Lookup(c); (err != nil) != tt.wantErr {
				t.Errorf("Lookup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
