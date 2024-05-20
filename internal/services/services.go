package services

import (
	"context"

	"github.com/MWT-proger/go-sqlc-clean-architecture-example/configs"
)

type UserStorer interface {
}

type UserService struct {
	conf      *configs.Config
	userStore UserStorer
}

func NewUserService(s UserStorer, config *configs.Config) *UserService {

	return &UserService{
		conf:      config,
		userStore: s,
	}
}

func (s *UserService) UserLogin(ctx context.Context, login string, password string) (string, error) {
	return "", nil
}

func (s *UserService) UserRegister(ctx context.Context, login string, password string) (string, error) {
	return "", nil
}
