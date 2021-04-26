package app

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	ProviderAppSet,

	NewBeatmapUseCase,
	NewBeatmapSetUseCase,
	NewChatUseCase,
	NewFriendUseCase,
	NewUserUseCase,
	NewOAuthUseCase,
)
