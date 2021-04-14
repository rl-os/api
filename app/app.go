package app

import (
	"context"
	"github.com/google/wire"
	"github.com/rl-os/api/repository"
	"github.com/rl-os/api/services/bancho"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// ProviderSet provide DI
var ProviderSet = wire.NewSet(New, NewOptions)

// Options is app configuration struct
type Options struct {
	JWT struct {
		Secret string
	}
}

type App struct {
	Context context.Context
	Options *Options

	BeatmapRepository    repository.Beatmap
	BeatmapSetRepository repository.BeatmapSet
	ChatRepository       repository.Chat
	FriendRepository     repository.Friend
	OAuthRepository      repository.OAuth
	UserRepository       repository.User

	BanchoClient *bancho.Client
}

// NewOptions create and parse config from viper instance
func NewOptions(logger *zerolog.Logger, v *viper.Viper) (*Options, error) {
	o := Options{}

	logger.Debug().
		Msg("Loading config file")
	if err := v.UnmarshalKey("app", &o); err != nil {
		return nil, err
	}

	logger.Debug().
		Interface("app", o).
		Msg("Loaded config")

	return &o, nil
}

// New with DI
func New(
	options *Options,
	bancho *bancho.Client,

	BeatmapRepository repository.Beatmap,
	BeatmapSetRepository repository.BeatmapSet,
	ChatRepository repository.Chat,
	FriendRepository repository.Friend,
	OAuthRepository repository.OAuth,
	UserRepository repository.User,
) *App {
	app := &App{
		Options:              options,
		BanchoClient:         bancho,
		BeatmapRepository:    BeatmapRepository,
		BeatmapSetRepository: BeatmapSetRepository,
		ChatRepository:       ChatRepository,
		FriendRepository:     FriendRepository,
		OAuthRepository:      OAuthRepository,
		UserRepository:       UserRepository,
	}

	return app
}

func (a *App) SetContext(ctx context.Context) {
	a.Context = ctx
}
