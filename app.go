package main

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/xxxcrel/kafka-console/pkg/api"
	"github.com/xxxcrel/kafka-console/pkg/config"
)

// App struct
type App struct {
	ctx context.Context
	api *api.API
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
	startupLogger := zap.NewExample()

	cfg, err := config.LoadConfig(startupLogger)
	if err != nil {
		startupLogger.Fatal("failed to load config", zap.Error(err))
	}
	err = cfg.Validate()
	if err != nil {
		startupLogger.Fatal("failed to validate config", zap.Error(err))
	}

	a := api.New(&cfg)
	a.Start()
	app.api = a
	fmt.Printf("startup: %s", app.api.Cfg.Kafka.Brokers)
}
