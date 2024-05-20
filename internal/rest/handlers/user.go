package handlers

import (
	"context"
	"net/http"

	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/logger"
)

type UserServicer interface {
	UserLogin(ctx context.Context, login string, password string) (string, error)
	UserRegister(ctx context.Context, login string, password string) (string, error)
}

func (h *APIHandler) UserRegister(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("APIHandler - UserRegister")

}

func (h *APIHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("APIHandler - UserLogin")

}
