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
	"github.com/mserebryaakov/tomb-raider/internal/httpserver/handlers"
	"github.com/mserebryaakov/tomb-raider/internal/httpserver/middleware/metrics"
	"github.com/mserebryaakov/tomb-raider/internal/httpserver/middleware/tracing"
	"github.com/mserebryaakov/tomb-raider/internal/httpserver/services"
	"github.com/mserebryaakov/tomb-raider/internal/storage/postgre"
	"github.com/mserebryaakov/tomb-raider/pkg/jaeger"
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

	// Создание jaeger
	jCfg := jaeger.Cfg{
		JaegerEnable: env.JaegerEnable,
		ServiceName:  env.JaegerService,
		SamplerType:  env.JaegerSamplerType,
		SamplerParam: env.JaegerSamplerParam,
		JaegerHost:   env.JaegerHost,
		JaegerPort:   env.JaegerPort,
	}
	_, closer, err := jaeger.NewTracer(jCfg)
	if err != nil {
		log.Error("Ошибка инициализации трейсинга", slog.String("error", err.Error()))
	}
	if closer != nil {
		defer closer.Close()
	}

	// Создание слоя storage
	storage, err := postgre.New(env.DSN)
	if err != nil {
		log.Error("Ошибка инициализации storage", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// Создание словия service
	service := services.New(storage)

	router := chi.NewRouter()
	// Создание requestID под каждый запрос
	router.Use(middleware.RequestID)
	// Логирование метрик
	router.Use(metrics.WithMetrics(log))
	// Трейсинг
	router.Use(tracing.WithTracing)
	// Восстановление после паники
	router.Use(middleware.Recoverer)
	// URLformat для удобного роутинга
	router.Use(middleware.URLFormat)

	// Создание слоя handlers
	handler := handlers.New(service)
	handler.Register(router)

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
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxShutdown); err != nil {
		log.Error("Ошибка остановки HTTP-сервера", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("Сервис остановлен")
}
