package services

import "github.com/mserebryaakov/tomb-raider/internal/httpserver/model"

func (s *service) CreateNamespace(namespace model.Namespace) (string, error) {
	return s.storage.CreateNamespace(namespace)
}
