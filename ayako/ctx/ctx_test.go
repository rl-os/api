package ctx

import (
	"context"
	"testing"
)

func TestGetRequestId(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "With empty context value",
			args: args{
				ctx: context.TODO(),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Without context value with selected key",
			args: args{
				ctx: Pipe(context.TODO(), SetUserToken("valid_value")),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "With invalid value type",
			args: args{
				ctx: context.WithValue(context.TODO(), RequestId, -1),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "With valid value type",
			args: args{
				ctx: Pipe(context.TODO(), SetRequestID("valid_value")),
			},
			want:    "valid_value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRequestId(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequestId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRequestId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "With empty context value",
			args: args{
				ctx: context.TODO(),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Without context value with selected key",
			args: args{
				ctx: Pipe(context.TODO(), SetRequestID("valid_value")),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "With invalid value type",
			args: args{
				ctx: context.WithValue(context.TODO(), UserId, "not_uint"),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "With valid value type",
			args: args{
				ctx: Pipe(context.TODO(), SetUserID(101112)),
			},
			want:    101112,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserID(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserToken(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "With empty context value",
			args: args{
				ctx: context.TODO(),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Without context value with selected key",
			args: args{
				ctx: Pipe(context.TODO(), SetRequestID("valid_value")),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "With invalid value type",
			args: args{
				ctx: context.WithValue(context.TODO(), RequestId, -1),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "With valid value type",
			args: args{
				ctx: Pipe(context.TODO(), SetUserToken("valid_value")),
			},
			want:    "valid_value",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserToken(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
