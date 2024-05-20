package store

import (
	"context"
	"embed"
	"time"

	"github.com/pressly/goose/v3"

	"github.com/MWT-proger/go-sqlc-clean-architecture-example/internal/logger"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// migration вызывается при запуске программы,
// проверяет новые миграции
// и при неообходимости обновляет БД,
// возвращает ошибку в случае неудачи.
func (s *Store) migration(ctx context.Context) error {
	logger.Log.Info("Хранилище: Проверка и обновление миграций - ...")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.UpContext(ctx, s.db, "migrations"); err != nil {
		return err
	}

	logger.Log.Info("Хранилище: Проверка и обновление миграций - ок")

	return nil
}

// ping вызывается при запуске программы,
// проверяет соединение и возвращает ошибку в случае неудачи.
func (s *Store) ping(ctx context.Context) error {
	logger.Log.Info("Хранилище: Проверка соединения - ...")

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := s.db.PingContext(ctx); err != nil {
		return err
	}
	logger.Log.Info("Хранилище: Проверка соединения - ок")

	return nil
}
