package main

import (
	"context"
	"fmt"
	"time"
	"github.com/xxxcrel/kafka-console/pkg/api"
	"github.com/xxxcrel/kafka-console/pkg/config"
	loggerpkg "github.com/xxxcrel/kafka-console/pkg/logger"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
	defaultLogger := loggerpkg.NewSlogLogger()

	cfg, err := config.LoadConfig(defaultLogger)
	if err != nil {
		loggerpkg.FatalStartup("failed to load config", err)
	}
	err = cfg.Validate()
	if err != nil {
		loggerpkg.FatalStartup("failed to validate config", err)
	}

	a, err := api.New(&cfg)
	if err != nil {
		loggerpkg.FatalStartup("failed to create API", err)
	}

	// Create startup context with timeout
	startupTimeout := 6*time.Second + cfg.Kafka.Startup.TotalMaxTime()
	ctx, cancel := context.WithTimeout(context.Background(), startupTimeout)
	defer cancel()

	if err := a.Start(ctx); err != nil {
		loggerpkg.FatalStartup("failed to start API", err)
	}
}

// Greet returns a greeting for the given name
func (app *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
