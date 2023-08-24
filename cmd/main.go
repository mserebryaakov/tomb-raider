package main

import (
	"fmt"

	"github.com/mserebryaakov/tomb-raider/config"
	"github.com/mserebryaakov/tomb-raider/pkg/logger"
)

func main() {
	env := config.MustLoad()
	fmt.Printf("env: %+v\n", env)

	log := logger.New(env.Debug)

	log.Info("Hello from slog")
}
