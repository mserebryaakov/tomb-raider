package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mserebryaakov/tomb-raider/internal/httpserver/model"
	"github.com/opentracing/opentracing-go"
)

func (h *handler) CreateNamespace(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	parentSpan := opentracing.SpanFromContext(ctx)

	span := opentracing.StartSpan("CreateNamespace", opentracing.ChildOf(parentSpan.Context()))
	defer span.Finish()

	var namespace model.Namespace
	err := json.NewDecoder(r.Body).Decode(&namespace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateNamespace(namespace)
	if err != nil {
		span.SetTag("success", "false")
		span.SetTag("error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	span.SetTag("success", "true")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
