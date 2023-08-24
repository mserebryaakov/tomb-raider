package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Environment содержит общую конфигурацию, включая HTTP-сервер
type Environment struct {
	Debug bool `env:"DEBUG" env-required:"true" env-default:"false"`
	HTTPServer
}

// HTTPServer содержит конфигурацию для HTTP-сервера
type HTTPServer struct {
	Port        string        `env:"SERVER_PORT" env-required:"true"`
	Timeout     time.Duration `env:"SERVER_TIMEOUT" env-required:"true"`
	IdleTimeout time.Duration `env:"SERVER_IDLE_TIMEOUT" env-required:"true"`
}

// MustLoad загружает конфигурацию из переменных окружения, и в случае ошибки вызывает panic
func MustLoad() Environment {
	var env Environment

	if err := cleanenv.ReadEnv(&env); err != nil {
		log.Fatalf("Error reading environment: %v", err)
	}

	return env
}
