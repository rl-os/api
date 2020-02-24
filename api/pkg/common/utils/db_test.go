package utils

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestNullString_MarshalJSON(t *testing.T) {
	type fields struct {
		NullString sql.NullString
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Must decode valid string",
			fields: fields{
				NullString: sql.NullString{
					String: "test case 1", Valid: true,
				},
			},
			want:    []byte{34, 116, 101, 115, 116, 32, 99, 97, 115, 101, 32, 49, 34},
			wantErr: false,
		},
		{
			name: "Must return nil",
			fields: fields{
				NullString: sql.NullString{
					String: "test case 1", Valid: false,
				},
			},
			want:    []byte{110, 117, 108, 108},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NullString{
				NullString: tt.fields.NullString,
			}
			got, err := v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullString_UnmarshalJSON(t *testing.T) {
	type fields struct {
		NullString sql.NullString
	}
	type args struct {
		data []byte
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "Must decode to string",
			fields: fields{},
			args: args{
				data: []byte{34, 116, 101, 115, 116, 32, 99, 97, 115, 101, 32, 49, 34},
			},
			want:    "test case 1",
			wantErr: false,
		},
		{
			name:   "Must decode to string",
			fields: fields{},
			args: args{
				data: []byte{34, 101, 109, 112, 116, 121, 32, 115, 116, 114, 105, 110, 103, 34},
			},
			want:    "empty string",
			wantErr: false,
		},
		{
			name:   "Must decode null string",
			fields: fields{},
			args: args{
				data: []byte{110, 117, 108, 108},
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &NullString{
				NullString: tt.fields.NullString,
			}
			if err := v.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if (v.Valid == true && tt.want != v.String) ||
				(v.Valid == false && tt.want != "") {
				t.Errorf("NullString.String = %v, want %v", v.String, tt.want)
			}
		})
	}
}
