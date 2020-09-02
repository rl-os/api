package store

//go:generate mockgen -destination=./mocks/generated.go -source=store.go
//go:generate gowrap gen -g -p . -i OAuth -t layers/log.tmpl -o layers/log_oauth.go
//go:generate gowrap gen -g -p . -i Beatmap -t layers/log.tmpl -o layers/log_beatmap.go
//go:generate gowrap gen -g -p . -i BeatmapSet -t layers/log.tmpl -o layers/log_beatmapset.go
//go:generate gowrap gen -g -p . -i User -t layers/log.tmpl -o layers/log_user.go
//go:generate gowrap gen -g -p . -i Friend -t layers/log.tmpl -o layers/log_friends.go
//go:generate gowrap gen -g -p . -i Chat -t layers/log.tmpl -o layers/log_chat.go

import (
	"context"
	"github.com/rl-os/api/entity"
	"time"
)

type Store interface {
	OAuth() OAuth
	Beatmap() Beatmap
	BeatmapSet() BeatmapSet
	User() User
	Friend() Friend
	Chat() Chat

	Close()
}

type OAuth interface {
	CreateClient(ctx context.Context, name string, redirect string) (*entity.OAuthClient, error)
	GetClient(ctx context.Context, id uint, secret string) (*entity.OAuthClient, error)

	CreateToken(ctx context.Context, userId uint, clientID uint, clientSecret string, scopes string) (*entity.OAuthToken, error)
	RevokeToken(ctx context.Context, userId uint, accessToken string) error
	RefreshToken(ctx context.Context, refreshToken string, clientID uint, clientSecret string) (*entity.OAuthToken, error)
	ValidateToken(ctx context.Context, accessToken string) (*entity.OAuthToken, error)
}

type Beatmap interface {
	Create(ctx context.Context, from interface{}) (*entity.Beatmap, error)
	CreateBatch(ctx context.Context, from interface{}) (*[]entity.Beatmap, error)
	Update(ctx context.Context, id uint, from interface{}) (*entity.Beatmap, error)
	Delete(ctx context.Context, id uint) error

	Get(ctx context.Context, id uint) (*entity.SingleBeatmap, error)
	GetBySetId(ctx context.Context, beatmapsetId uint) []entity.Beatmap
}

type BeatmapSet interface {
	Create(ctx context.Context, from interface{}) (*entity.BeatmapSetFull, error)
	Update(ctx context.Context, id uint, from interface{}) (*entity.BeatmapSetFull, error)
	Delete(ctx context.Context, id uint) error

	Get(ctx context.Context, id uint) (*entity.BeatmapSetFull, error)
	GetAll(ctx context.Context, page int, limit int) (*[]entity.BeatmapSet, error)
	ComputeFields(ctx context.Context, set entity.BeatmapSetFull) (*entity.BeatmapSetFull, error)
	SetFavourite(ctx context.Context, userId uint, id uint) (uint, error)
	SetUnFavourite(ctx context.Context, userId uint, id uint) (uint, error)

	GetLatestId(ctx context.Context) (uint, error)
	GetIdsForUpdate(ctx context.Context, limit int) ([]uint, error)

	FetchFromBancho(ctx context.Context, id uint) (*entity.BeatmapSetFull, error)
}

type User interface {
	Create(ctx context.Context, name, email, pwd string) (*entity.User, error)
	Update(ctx context.Context, userId uint, from interface{}) (*entity.UserShort, error)

	Get(ctx context.Context, userId uint, mode string) (*entity.User, error)
	GetShort(ctx context.Context, userId uint, mode string) (*entity.UserShort, error)
	GetByBasic(ctx context.Context, login, pwd string) (*entity.UserShort, error)
	ComputeFields(ctx context.Context, user entity.User) (*entity.User, error)

	Activate(ctx context.Context, userId uint) error
	Deactivate(ctx context.Context, userId uint) error
	Ban(ctx context.Context, userId uint, time time.Duration) error
	UnBan(ctx context.Context, userId uint) error
	Mute(ctx context.Context, userId uint, time time.Duration) error
	UnMute(ctx context.Context, userId uint) error

	UpdateLastVisit(ctx context.Context, userId uint) error
}

type Friend interface {
	Add(ctx context.Context, userId, targetId uint) error
	Remove(ctx context.Context, userId, targetId uint) error
	GetSubscriptions(ctx context.Context, userId uint) (*[]entity.UserShort, error)
}

type Chat interface {
	CreatePM(ctx context.Context, userId, targetId uint, message string, isAction bool) (*entity.ChannelNewPm, error)

	Get(ctx context.Context, channelId uint) (*entity.Channel, error)
	GetOrCreatePm(ctx context.Context, userId, targetId uint) (*entity.Channel, error)
	GetPublic(ctx context.Context) (*[]entity.Channel, error)
	GetJoined(ctx context.Context, userId uint) (*[]entity.Channel, error)
	GetMessage(ctx context.Context, messageId uint) (*entity.ChatMessage, error)
	GetMessages(ctx context.Context, userId, since uint) (*[]entity.ChatMessage, error)
	GetUpdates(ctx context.Context, userId, since, channelId, limit uint) (*entity.ChannelUpdates, error)

	SendMessage(ctx context.Context, userId, channelId uint, content string, isAction bool) (*entity.ChatMessage, error)
	Join(ctx context.Context, userId, channelId uint) (*entity.Channel, error)
	Leave(ctx context.Context, userId, channelId uint) error
	ReadMessage()
}
