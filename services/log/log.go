package log

import (
	"github.com/google/wire"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	stdlog "log"
	"os"
	"time"
)

var ProviderSet = wire.NewSet(New, NewOptions)

// Options is log configuration struct
type Options struct {
	Level   string
	NoColor bool
}

func NewOptions(v *viper.Viper) (*Options, error) {
	options := Options{}

	if err := v.UnmarshalKey("log", &options); err != nil {
		return nil, err
	}

	return &options, nil
}

// New for init log library
func New(o *Options) (*zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(o.Level)
	if err != nil {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zerolog.SetGlobalLevel(level)

	logger := log.Output(
		zerolog.ConsoleWriter{
			NoColor:    o.NoColor,
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		},
	).With().Caller().Logger()

	stdlog.SetFlags(0)
	stdlog.SetOutput(logger)

	log.Logger = logger

	return &logger, err
}
