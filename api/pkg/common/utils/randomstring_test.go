package utils

import (
	"testing"
)

func TestGenerateRandomBytes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "random 6 bytes", args: args{n: 6}, wantErr: false},
		{name: "random 256 bytes", args: args{n: 256}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := GenerateRandomBytes(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRandomBytes() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(bytes) != tt.args.n {
				t.Errorf("GenerateRandomBytes len = %v, want %v", len(bytes), tt.args.n)
			}
		})
	}
}

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "random 6 bytes", args: args{s: 6}},
		{name: "random 256 bytes", args: args{s: 256}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = GenerateRandomString(tt.args.s)
		})
	}
}
