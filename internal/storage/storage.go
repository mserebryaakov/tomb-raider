package storage

import "github.com/mserebryaakov/tomb-raider/internal/httpserver/model"

type IStorage interface {
	Close()
	CreateNamespace(namespace model.Namespace) (string, error)
}
