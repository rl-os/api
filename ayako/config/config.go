package config

import (
	"time"
)

type Config struct {
	// In case of json files, this field will be used only when compiled with
	// go 1.10 or later.
	// This field will be ignored when compiled with go versions lower than 1.10.
	ErrorOnUnmatchedKeys bool

	Environment        string
	ENVPrefix          string
	AutoReload         bool
	AutoReloadInterval time.Duration
	AutoReloadCallback func(config interface{})
	configModTimes     map[string]time.Time

	// Configurations
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Database struct {
		DSN    string `json:"dsn"`
		Driver string `json:"driver"`
	}
	Mirror struct {
		S3 struct {
			SecretKey int64  `json:"secret_key"`
			Bucket    string `json:"bucket"`
			SecretID  int64  `json:"secret_id"`
		}
		Bancho struct {
			Username          string `json:"username"`
			ClientSecret      string `json:"client_secret"`
			Password          string `json:"password"`
			UsingRefreshToken bool   `json:"using_refresh_token"`
			ClientID          int64  `json:"client_id"`
		}
	}
}
