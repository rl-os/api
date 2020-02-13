package oauth

import (
	"reflect"
	"testing"
)

func TestCreateOAuthToken(t *testing.T) {
	type args struct {
		userID   uint
		clientID uint
		secretID string
		scopes   string
	}
	tests := []struct {
		name           string
		args           args
		wantOAuthToken error
		wantErr        error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOAuthToken, gotErr := CreateOAuthToken(tt.args.userID, tt.args.clientID, tt.args.secretID, tt.args.scopes)
			if !reflect.DeepEqual(gotOAuthToken, tt.wantOAuthToken) {
				t.Errorf("CreateOAuthToken() gotOAuthToken = %v, want %v", gotOAuthToken, tt.wantOAuthToken)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("CreateOAuthToken() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestRefreshOAuthToken(t *testing.T) {
	type args struct {
		refreshToken string
		clientID     uint
		secretID     string
		scopes       string
	}
	tests := []struct {
		name           string
		args           args
		wantOAuthToken error
		wantErr        error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOAuthToken, gotErr := RefreshOAuthToken(tt.args.refreshToken, tt.args.clientID, tt.args.secretID, tt.args.scopes)
			if !reflect.DeepEqual(gotOAuthToken, tt.wantOAuthToken) {
				t.Errorf("RefreshOAuthToken() gotOAuthToken = %v, want %v", gotOAuthToken, tt.wantOAuthToken)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("RefreshOAuthToken() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestRevokeOAuthToken(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RevokeOAuthToken(); (err != nil) != tt.wantErr {
				t.Errorf("RevokeOAuthToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOAuthToken(t *testing.T) {
	type args struct {
		accessToken string
	}
	tests := []struct {
		name           string
		args           args
		wantOAuthToken error
		wantErr        error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOAuthToken, gotErr := ValidateOAuthToken(tt.args.accessToken)
			if !reflect.DeepEqual(gotOAuthToken, tt.wantOAuthToken) {
				t.Errorf("ValidateOAuthToken() gotOAuthToken = %v, want %v", gotOAuthToken, tt.wantOAuthToken)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("ValidateOAuthToken() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
