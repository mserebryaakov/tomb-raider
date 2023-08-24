package logger

import (
	"log/slog"
	"os"
)

func New(debug bool) *slog.Logger {
	handlerOptions := &slog.HandlerOptions{Level: slog.LevelInfo}

	if debug {
		handlerOptions.Level = slog.LevelDebug
	}

	handler := slog.NewJSONHandler(os.Stdout, handlerOptions)
	return slog.New(handler)
}
