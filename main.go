package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/rl-os/api/config"
	"github.com/rl-os/api/services"
	sql "github.com/rl-os/api/store/gorm"
)

func main() {
	cfg := config.Init("config.yaml")

	store := sql.Init(cfg, &services.Services{Bancho: nil})

	user, err := store.
		User().
		Get(context.TODO(), 100, "")

	log.Info(user.Country, err)
}
