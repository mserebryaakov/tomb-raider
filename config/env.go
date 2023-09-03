package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Environment содержит общую конфигурацию, включая HTTP-сервер
type Environment struct {
	HTTPServer
	Storage
	Jaeger
}

// HTTPServer содержит конфигурацию для HTTP-сервера
type HTTPServer struct {
	Port        string        `env:"SERVER_PORT" env-required:"true"`
	Timeout     time.Duration `env:"SERVER_TIMEOUT" env-required:"true"`
	IdleTimeout time.Duration `env:"SERVER_IDLE_TIMEOUT" env-required:"true"`
}

// DSN строка подключения в БД
type Storage struct {
	DSN string `env:"DSN" env-required:"true"`
}

// Jaeger трейсинг
type Jaeger struct {
	JaegerEnable       bool    `env:"JAEGER_ENABLE" env-required:"true"`
	JaegerService      string  `env:"JAEGER_SERVICE_NAME" env-required:"true"`
	JaegerHost         string  `env:"JAEGER_AGENT_HOST" env-required:"true"`
	JaegerPort         string  `env:"JAEGER_AGENT_PORT" env-required:"true"`
	JaegerSamplerType  string  `env:"JAEGER_SAMPLER_TYPE" env-required:"true"`
	JaegerSamplerParam float64 `env:"JAEGER_SAMPLER_PARAM" env-required:"true"`
}

// MustLoad загружает конфигурацию из переменных окружения, и в случае ошибки вызывает panic
func MustLoad() Environment {
	var env Environment

	if err := cleanenv.ReadEnv(&env); err != nil {
		log.Fatalf("Error reading environment: %v", err)
	}

	return env
}
