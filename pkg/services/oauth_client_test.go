package services

import (
	"reflect"
	"testing"
)

func TestCreateOAuthClient(t *testing.T) {
	type args struct {
		userId   int
		name     string
		redirect string
	}
	tests := []struct {
		name       string
		args       args
		wantClient OAuthClient
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClient, err := CreateOAuthClient(tt.args.userId, tt.args.name, tt.args.redirect)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOAuthClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotClient, tt.wantClient) {
				t.Errorf("CreateOAuthClient() gotClient = %v, want %v", gotClient, tt.wantClient)
			}
		})
	}
}

func TestFindOAuthClient(t *testing.T) {
	type args struct {
		clientId     string
		clientSecret string
	}
	tests := []struct {
		name    string
		args    args
		wantOut OAuthToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := FindOAuthClient(tt.args.clientId, tt.args.clientSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOAuthClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("FindOAuthClient() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
