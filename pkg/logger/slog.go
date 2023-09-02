package logger

import (
	"context"
	"log/slog"
	"os"
)

var (
	global *slog.Logger
)

type loggerContextKey struct{}

func init() {
	setLogger(new())
}

func setLogger(l *slog.Logger) {
	global = l
}

func new() *slog.Logger {
	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)
}

func FromContext(ctx context.Context) *slog.Logger {
	l := global
	if logger, ok := ctx.Value(loggerContextKey{}).(*slog.Logger); ok {
		l = logger
	}
	return l
}

func WithContext(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey{}, l)
}
