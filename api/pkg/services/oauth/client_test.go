package oauth

import (
	"reflect"
	"testing"
)

func TestCreateOAuthClient(t *testing.T) {
	type args struct {
		userID   uint
		name     string
		redirect string
	}
	tests := []struct {
		name       string
		args       args
		wantClient Client
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClient, err := CreateOAuthClient(tt.args.userID, tt.args.name, tt.args.redirect)
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
		clientID     uint
		clientSecret string
	}
	tests := []struct {
		name    string
		args    args
		wantOut Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := FindOAuthClient(tt.args.clientID, tt.args.clientSecret)
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
