package postgre

import "context"

type Namespace struct {
	ID   string
	Name string
}

func (s *storage) CreateNamespace(namespace Namespace) (string, error) {
	var id string
	err := s.pool.QueryRow(context.Background(), "INSERT INTO namespace (name) VALUES ($1) RETURNING id", namespace.Name).Scan(&id)
	return id, err
}
