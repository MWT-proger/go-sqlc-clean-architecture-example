package store

import (
	"context"
	"database/sql"

	"github.com/MWT-proger/go-sqlc-clean-architecture-example/configs"
	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/logger"
)

type UserStore struct {
}

func NewUserStore() *UserStore {

	return &UserStore{}
}

type Store struct {
	db *sql.DB
}

// NewStore создаёт и возвращает новый экземпляр хранилища.
func NewStore(ctx context.Context, conf *configs.Config) (*Store, error) {
	var storage = Store{}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := storage.Init(ctx, conf); err != nil {
		return nil, err
	}
	return &storage, nil
}

// Init вызывается при запуске программы,
// инициализирует соединение
// и возвращает ошибку в случае не удачи.
func (s *Store) Init(ctx context.Context, conf *configs.Config) error {
	logger.Log.Info("Хранилище: Подключение - ...")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var db, err = sql.Open("pgx", conf.DatabaseDSN)

	if err != nil {
		return err
	}

	s.db = db

	if err := s.ping(ctx); err != nil {
		return err
	}

	if err := s.migration(ctx); err != nil {
		return err
	}
	logger.Log.Info("Хранилище: Подключение - ок")

	return nil

}

func (s *Store) GetDB() *sql.DB {
	return s.db
}

// Close вызывается при завершение программы,
// закрывает соединение и возвращает ошибку в случае не удачи.
func (s *Store) Close() error {
	logger.Log.Info("Хранилище: Закрытие соединения - ...")

	if err := s.db.Close(); err != nil {
		return err
	}
	logger.Log.Info("Хранилище: Закрытие соединения - ок")

	return nil
}
