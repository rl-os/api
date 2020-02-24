package utils

import "testing"

func TestCrypt(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Must return valid hashed password",
			args:    args{pwd: "password"},
			wantErr: false,
		},
		{
			name:    "Must return valid hashed password",
			args:    args{pwd: "SomeVeryStrongPassword"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHash(tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if CompareHash(got, tt.args.pwd) != true {
				t.Errorf("GetHash() got = %v, but CompareHash return false", got)
			}
		})
	}
}
