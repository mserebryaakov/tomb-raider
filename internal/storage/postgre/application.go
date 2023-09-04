package postgre

import (
	"context"

	"github.com/mserebryaakov/tomb-raider/internal/httpserver/model"
)

func (s *storage) CreateApplication(app model.Application) (string, error) {
	var id string
	err := s.pool.QueryRow(context.Background(), "INSERT INTO application (name, namespace_id, code, data) VALUES ($1, $2, $3, $4) RETURNING id", app.Name, app.NamespaceID, app.Code, app.Data).Scan(&id)
	return id, err
}

func (s *storage) ReadApplication(id string) (model.Application, error) {
	var app model.Application
	app.ID = model.UUID(id)
	err := s.pool.QueryRow(context.Background(), "SELECT name, namespace_id, code, data FROM application WHERE id = $1", id).Scan(&app.Name, &app.NamespaceID, &app.Code, &app.Data)
	return app, err
}

func (s *storage) UpdateApplication(app model.Application) error {
	_, err := s.pool.Exec(context.Background(), "UPDATE application SET data = $2 WHERE id = $1", app.ID, app.Data)
	return err
}

func (s *storage) DeleteApplication(id string) error {
	_, err := s.pool.Exec(context.Background(), "DELETE FROM application WHERE id = $1", id)
	return err
}
