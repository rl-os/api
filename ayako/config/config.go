package config

import (
	"time"
)

type Config struct {
	// In case of json files, this field will be used only when compiled with
	// go 1.10 or later.
	// This field will be ignored when compiled with go versions lower than 1.10.
	ErrorOnUnmatchedKeys bool

	AutoReload         bool
	AutoReloadInterval time.Duration
	AutoReloadCallback func(config *Config)
	configModTimes     map[string]time.Time

	// Configurations
	Server struct {
		Host          string `default:"0.0.0.0"`
		Port          string `default:"2400"`
		EnableProfile bool   `default:"false"`
		EnableJobs    bool   `default:"false"`
	}
	JWT struct {
		Secret            string `default:"somesupersecretstring"`
		HoursBeforeRevoke int    `default:"1"`
	}
	Database struct {
		DSN    string `default:"postgres://postgres:postgres@/osuserver?sslmode=disable"`
		Driver string `default:"postgres"`
	}
	Mirror struct {
		S3 struct {
			SecretKey int64
			Bucket    string
			SecretID  int64
		}
		Bancho struct {
			Username string
			Password string
		}
	}
}
