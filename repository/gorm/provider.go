package gorm

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewOptions,
	NewSupplier,

	NewBeatmapRepository,

	NewBeatmapSetRepository,

	NewChatRepository,

	NewFriendRepository,

	NewOAuthRepository,

	NewUserRepository,
)
