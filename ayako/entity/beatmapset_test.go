package entity

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"testing"
)

func TestAvailability_Scan(t *testing.T) {
	type fields struct {
		DownloadDisabled bool
		MoreInformation  interface{}
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    fields
		wantErr bool
	}{
		{
			name: "scan from string",
			args: args{
				value: []byte("{\"download_disabled\":false,\"more_information\":\"None\"}"),
			},
			want: fields{
				DownloadDisabled: false,
				MoreInformation:  "None",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Availability{
				DownloadDisabled: true,
				MoreInformation:  "None",
			}
			if err := c.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAvailability_Value(t *testing.T) {
	type fields struct {
		DownloadDisabled bool
		MoreInformation  interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name:    "json value",
			fields:  fields{true, "None"},
			want:    []byte("{\"download_disabled\":true,\"more_information\":\"None\"}"),
			wantErr: false,
		},
		{
			name:    "json value with array",
			fields:  fields{true, []string{"first", "second"}},
			want:    []byte("{\"download_disabled\":true,\"more_information\":[\"first\",\"second\"]}"),
			wantErr: false,
		},
		{
			name: "json value with dict",
			fields: fields{true, struct {
				Value string `json:"value"`
			}{"None"}},
			want:    []byte("{\"download_disabled\":true,\"more_information\":{\"value\":\"None\"}}"),
			wantErr: false,
		},
		{
			name:    "json value with null",
			fields:  fields{true, nil},
			want:    []byte("{\"download_disabled\":true,\"more_information\":null}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Availability{
				DownloadDisabled: tt.fields.DownloadDisabled,
				MoreInformation:  tt.fields.MoreInformation,
			}
			got, err := c.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCovers_Value(t *testing.T) {
	tests := []struct {
		name    string
		fields  Covers
		wantErr bool
	}{
		{
			name: "json value",
			fields: Covers{
				Cover:       "https://assets.ppy.sh/beatmaps/365924/covers/cover.jpg?0",
				Cover2X:     "https://assets.ppy.sh/beatmaps/365924/covers/cover@2x.jpg?0",
				Card:        "https://assets.ppy.sh/beatmaps/365924/covers/card.jpg?0",
				Card2X:      "https://assets.ppy.sh/beatmaps/365924/covers/card@2x.jpg?0",
				List:        "https://assets.ppy.sh/beatmaps/365924/covers/list.jpg?0",
				List2X:      "https://assets.ppy.sh/beatmaps/365924/covers/list@2x.jpg?0",
				Slimcover:   "https://assets.ppy.sh/beatmaps/365924/covers/slimcover.jpg?0",
				Slimcover2X: "https://assets.ppy.sh/beatmaps/365924/covers/slimcover@2x.jpg?0",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Covers{
				Cover:       tt.fields.Cover,
				Cover2X:     tt.fields.Cover2X,
				Card:        tt.fields.Card,
				Card2X:      tt.fields.Card2X,
				List:        tt.fields.List,
				List2X:      tt.fields.List2X,
				Slimcover:   tt.fields.Slimcover,
				Slimcover2X: tt.fields.Slimcover2X,
			}
			got, err := c.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}

			b, _ := json.Marshal(tt.fields)

			if !reflect.DeepEqual(got, b) {
				t.Errorf("Value() got = %v, want %v", got, b)
			}
		})
	}
}

func TestCurrentUserAttributes_Scan(t *testing.T) {
	type fields struct {
		CanDelete     bool
		CanHype       bool
		CanHypeReason interface{}
		CanLove       bool
		IsWatching    bool
		NewHypeTime   interface{}
		RemainingHype int64
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CurrentUserAttributes{
				CanDelete:     tt.fields.CanDelete,
				CanHype:       tt.fields.CanHype,
				CanHypeReason: tt.fields.CanHypeReason,
				CanLove:       tt.fields.CanLove,
				IsWatching:    tt.fields.IsWatching,
				NewHypeTime:   tt.fields.NewHypeTime,
				RemainingHype: tt.fields.RemainingHype,
			}
			if err := c.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() errors = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCurrentUserAttributes_Value(t *testing.T) {
	type fields struct {
		CanDelete     bool
		CanHype       bool
		CanHypeReason interface{}
		CanLove       bool
		IsWatching    bool
		NewHypeTime   interface{}
		RemainingHype int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrentUserAttributes{
				CanDelete:     tt.fields.CanDelete,
				CanHype:       tt.fields.CanHype,
				CanHypeReason: tt.fields.CanHypeReason,
				CanLove:       tt.fields.CanLove,
				IsWatching:    tt.fields.IsWatching,
				NewHypeTime:   tt.fields.NewHypeTime,
				RemainingHype: tt.fields.RemainingHype,
			}
			got, err := c.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDescription_Scan(t *testing.T) {
	type fields struct {
		Description string
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Description{
				Description: tt.fields.Description,
			}
			if err := c.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() errors = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDescription_Value(t *testing.T) {
	type fields struct {
		Description string
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Description{
				Description: tt.fields.Description,
			}
			got, err := c.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenre_Scan(t *testing.T) {
	type fields struct {
		ID   int64
		Name string
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Genre{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			if err := c.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() errors = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenre_Value(t *testing.T) {
	type fields struct {
		ID   int64
		Name string
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Genre{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			got, err := c.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHype_Scan(t *testing.T) {
	type fields struct {
		Current  int64
		Required int64
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Hype{
				Current:  tt.fields.Current,
				Required: tt.fields.Required,
			}
			if err := c.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() errors = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHype_Value(t *testing.T) {
	type fields struct {
		Current  int64
		Required int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Hype{
				Current:  tt.fields.Current,
				Required: tt.fields.Required,
			}
			got, err := c.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Scan(t *testing.T) {
	type fields struct {
		ID            int64
		Username      string
		ProfileColour interface{}
		AvatarURL     string
		CountryCode   string
		DefaultGroup  DefaultGroup
		IsActive      bool
		IsBot         bool
		IsOnline      bool
		IsSupporter   bool
		LastVisit     *string
		PmFriendsOnly bool
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &User{
				ID:            tt.fields.ID,
				Username:      tt.fields.Username,
				ProfileColour: tt.fields.ProfileColour,
				AvatarURL:     tt.fields.AvatarURL,
				CountryCode:   tt.fields.CountryCode,
				DefaultGroup:  tt.fields.DefaultGroup,
				IsActive:      tt.fields.IsActive,
				IsBot:         tt.fields.IsBot,
				IsOnline:      tt.fields.IsOnline,
				IsSupporter:   tt.fields.IsSupporter,
				LastVisit:     tt.fields.LastVisit,
				PmFriendsOnly: tt.fields.PmFriendsOnly,
			}
			if err := c.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() errors = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_Value(t *testing.T) {
	type fields struct {
		ID            int64
		Username      string
		ProfileColour interface{}
		AvatarURL     string
		CountryCode   string
		DefaultGroup  DefaultGroup
		IsActive      bool
		IsBot         bool
		IsOnline      bool
		IsSupporter   bool
		LastVisit     *string
		PmFriendsOnly bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := User{
				ID:            tt.fields.ID,
				Username:      tt.fields.Username,
				ProfileColour: tt.fields.ProfileColour,
				AvatarURL:     tt.fields.AvatarURL,
				CountryCode:   tt.fields.CountryCode,
				DefaultGroup:  tt.fields.DefaultGroup,
				IsActive:      tt.fields.IsActive,
				IsBot:         tt.fields.IsBot,
				IsOnline:      tt.fields.IsOnline,
				IsSupporter:   tt.fields.IsSupporter,
				LastVisit:     tt.fields.LastVisit,
				PmFriendsOnly: tt.fields.PmFriendsOnly,
			}
			got, err := c.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}
