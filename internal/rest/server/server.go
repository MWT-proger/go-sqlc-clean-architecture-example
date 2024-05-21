package server

import (
	"net/http"

	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/config"
	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/logger"
	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/rest/handlers"
)

// Run запускает сервер и слушает его по указанному хосту.
func Run(h *handlers.APIHandler, conf *config.Config) error {
	r := router(h)
	logger.Log.Info("Run Server", logger.StringField("host", conf.HostServer))
	return http.ListenAndServe(conf.HostServer, r)
}
