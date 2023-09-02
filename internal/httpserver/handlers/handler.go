package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/mserebryaakov/tomb-raider/internal/httpserver/services"
)

type handler struct {
	service services.IService
}

func New(service services.IService) *handler {
	return &handler{
		service,
	}
}

func (h *handler) Register(r chi.Router) {
	r.Route("/namespace", func(r chi.Router) {
		r.Post("/", h.CreateNamespace())
	})
}
