package scopes

import (
	"github.com/deissh/rl/api/pkg/services/oauth"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func next(_ echo.Context) error { return nil }

func TestRequired1(t *testing.T) {
	type args struct {
		required []string
	}
	type data struct {
		scopes string
		token  interface{}
	}

	// setting up echo test case
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	function := Required(oauth.ChatsScope)
	handler := function(next)

	tests := []struct {
		name    string
		args    args
		data    data
		wantErr bool
	}{
		{
			name: "Must allow perform request with all scopes",
			args: args{
				required: []string{oauth.ChatsScope},
			},
			data:    data{scopes: "*"},
			wantErr: false,
		},
		{
			name: "Must allow perform request with eq scope",
			args: args{
				required: []string{oauth.ChatsScope},
			},
			data:    data{scopes: oauth.ChatsScope},
			wantErr: false,
		},
		{
			name: "Must allow perform request with eq scopes",
			args: args{
				required: []string{oauth.ChatsScope},
			},
			data:    data{scopes: oauth.ChatsScope + "," + oauth.ProfileScope},
			wantErr: false,
		},
		{
			name: "Invalid current_user_token in reqest_context",
			args: args{
				required: []string{oauth.ChatsScope},
			},
			data:    data{scopes: "*", token: "invalid token"},
			wantErr: true,
		},
		{
			name: "User dont have required scopes",
			args: args{
				required: []string{oauth.ChatsScope},
			},
			data:    data{scopes: "profile"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setting up user token
			token, _ := oauth.TokenCreate(1, 2, tt.data.scopes)
			if tt.data.token != nil {
				c.Set("current_user_token", tt.data.token)
			} else {
				c.Set("current_user_token", token)
			}

			if err := handler(c); (err != nil) != tt.wantErr {
				t.Errorf("Required().Error = %v, want errors", err)
			}
		})
	}
}
