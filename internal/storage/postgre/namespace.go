package postgre

import (
	"context"

	"github.com/mserebryaakov/tomb-raider/internal/httpserver/model"
)

func (s *storage) CreateNamespace(namespace model.Namespace) (string, error) {
	var id string
	err := s.pool.QueryRow(context.Background(), "INSERT INTO namespace (name) VALUES ($1) RETURNING id", namespace.Name).Scan(&id)
	return id, err
}
