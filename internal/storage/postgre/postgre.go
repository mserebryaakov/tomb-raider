package postgre

import (
	"context"

	istorage "github.com/mserebryaakov/tomb-raider/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type storage struct {
	pool *pgxpool.Pool
}

func New(dsn string) (istorage.IStorage, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return &storage{}, err
	}

	return &storage{
		pool,
	}, nil
}

func (s *storage) Close() {
	s.pool.Close()
}
