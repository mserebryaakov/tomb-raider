package services

import "github.com/mserebryaakov/tomb-raider/internal/storage"

type servive struct {
	storage storage.IStorage
}

type IService interface {
}

func New(storage storage.IStorage) IService {
	return &servive{
		storage,
	}
}
