package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mserebryaakov/tomb-raider/config"
	httplogger "github.com/mserebryaakov/tomb-raider/internal/httpserver/middleware/logger"
	"github.com/mserebryaakov/tomb-raider/internal/storage/postgre"
	"github.com/mserebryaakov/tomb-raider/pkg/logger"
)

func main() {
	// Загрузка переменных окружения
	env := config.MustLoad()

	// Создание контекста
	ctx := context.Background()
	// Получение логера
	log := logger.FromContext(ctx)

	log.Info("Старт сервиса", slog.Any("env", env))

	// Создание слоя storage
	storage, err := postgre.New(env.DSN)
	if err != nil {
		log.Error("Ошибка инициализации storage", slog.String("error", err.Error()))
		os.Exit(1)
	}

	_ = storage

	router := chi.NewRouter()
	// Создание requestID под каждый запрос
	router.Use(middleware.RequestID)
	// Логирование запроса
	router.Use(httplogger.NewMiddleware(log))
	// Восстановление после паники
	router.Use(middleware.Recoverer)
	// URLformat для удобного роутинга
	router.Use(middleware.URLFormat)

	server := &http.Server{
		Addr:    env.Port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Ошибка запуска HTTP-сервера", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	// Gracefull shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Error("Ошибка остановки HTTP-сервера", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("Сервис остановлен")
}
