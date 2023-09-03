package services

import (
	"github.com/mserebryaakov/tomb-raider/internal/httpserver/model"
	"github.com/mserebryaakov/tomb-raider/internal/storage"
)

type service struct {
	storage storage.IStorage
}

type IService interface {
	CreateNamespace(namespace model.Namespace) (string, error)
}

func New(storage storage.IStorage) IService {
	return &service{
		storage,
	}
}
