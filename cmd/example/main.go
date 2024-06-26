package main

import (
	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/config"
	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/logger"
	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/rest/handlers"
	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/rest/server"
	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/services"
	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/store"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	if err := logger.Initialize(conf.LogLevel); err != nil {
		panic(err)
	}

	userstore := store.NewUserStore()
	userservice := services.NewUserService(userstore, conf)
	handlers, err := handlers.NewAPIHandler(userservice)

	if err != nil {
		panic(err)
	}

	if err := server.Run(handlers, conf); err != nil {
		panic(err)
	}
}
