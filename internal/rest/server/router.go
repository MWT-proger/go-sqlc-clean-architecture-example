package server

import (
	"github.com/go-chi/chi/v5"

	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/rest/handlers"
)

// router Перенаправляет запросы на необходимые хендлеры.
func router(h *handlers.APIHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/user/register", h.UserRegister)
	r.Post("/user/login", h.UserLogin)

	return r
}
