// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rl-os/api/api"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/repository/gorm"
	"github.com/rl-os/api/services"
	"github.com/rl-os/api/services/cache"
	"github.com/rl-os/api/services/config"
	"github.com/rl-os/api/services/log"
	"github.com/rl-os/api/services/redis"
	"github.com/rl-os/api/services/transports"
	"github.com/rl-os/api/services/transports/http"
	"github.com/rl-os/api/services/validator"
)

// Injectors from wire.go:

func Injector(configPath string) (transports.Server, error) {
	viper, err := config.New(configPath)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(logger, viper)
	if err != nil {
		return nil, err
	}
	redisOptions := cache.NewRedisOptions()
	options2, err := redis.NewOptions(logger, viper)
	if err != nil {
		return nil, err
	}
	pool := redis.New(logger, options2)
	cacheCache, err := cache.NewRedis(redisOptions, pool)
	if err != nil {
		return nil, err
	}
	appOptions, err := app.NewOptions(logger, viper)
	if err != nil {
		return nil, err
	}
	inst, err := validator.New()
	if err != nil {
		return nil, err
	}
	appApp := app.New(cacheCache, appOptions, inst)
	gormOptions, err := gorm.NewOptions(logger, viper)
	if err != nil {
		return nil, err
	}
	supplier, err := gorm.NewSupplier(logger, gormOptions)
	if err != nil {
		return nil, err
	}
	user := gorm.NewUserRepository(supplier)
	userUseCase := app.NewUserUseCase(appApp, user)
	userController := api.NewUserController(userUseCase, logger)
	chat := gorm.NewChatRepository(supplier)
	chatUseCase := app.NewChatUseCase(appApp, chat)
	chatController := api.NewChatController(chatUseCase, logger)
	friend := gorm.NewFriendRepository(supplier)
	friendUseCase := app.NewFriendUseCase(appApp, friend)
	friendController := api.NewFriendController(friendUseCase, logger)
	beatmap := gorm.NewBeatmapRepository(supplier)
	beatmapUseCase := app.NewBeatmapUseCase(appApp, beatmap)
	beatmapController := api.NewBeatmapController(beatmapUseCase, logger)
	beatmapSet := gorm.NewBeatmapSetRepository(supplier)
	beatmapSetUseCase := app.NewBeatmapSetUseCase(appApp, beatmapUseCase, beatmapSet)
	beatmapSetController := api.NewBeatmapSetController(beatmapSetUseCase, logger)
	currentUserController := api.NewCurrentUserController(userUseCase, logger)
	oAuthOptions, err := gorm.NewOAuthOptions(logger, viper)
	if err != nil {
		return nil, err
	}
	oAuth := gorm.NewOAuthRepository(oAuthOptions, supplier)
	oAuthUseCase := app.NewOAuthUseCase(appApp, oAuth, user)
	oAuthTokenController := api.NewOAuthTokenController(oAuthUseCase, logger)
	oAuthClientController := api.NewOAuthClientController(oAuthUseCase, logger)
	initControllers := api.CreateInitControllersFn(appApp, userController, chatController, friendController, beatmapController, beatmapSetController, currentUserController, oAuthTokenController, oAuthClientController)
	echo := http.NewRouter(httpOptions, logger, initControllers)
	server, err := http.New(httpOptions, logger, echo)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// wire.go:

var providerSet = wire.NewSet(services.ProviderSet, api.ProviderSet, app.ProviderSet, http.ProviderSet, gorm.ProviderSet, cache.ProviderSet)
