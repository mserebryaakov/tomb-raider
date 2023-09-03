package tracing

import (
	"context"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
)

type contextKey string

const startKey contextKey = "start"

// Middleware трейсов
func WithTracing(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), startKey, time.Now())
		span := opentracing.StartSpan(r.RequestURI)
		ctx = opentracing.ContextWithSpan(ctx, span)
		r = r.WithContext(ctx)

		span.SetTag("http.method", r.Method)
		span.SetTag("http.url", r.URL.String())

		next.ServeHTTP(w, r)

		span.Finish()
	}

	return http.HandlerFunc(fn)
}
