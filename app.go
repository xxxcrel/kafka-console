package main

import (
	"context"
	"fmt"

	"github.com/cloudhut/common/rest"
	"github.com/xxxcrel/kafka-console/pkg/kconsole"
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

// Greet returns a greeting for the given name
func (app *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (app *App) GetClusterInfo() (*kconsole.ClusterInfo, error) {
	fmt.Printf("hello\n")
	return app.api.ConsoleSvc.GetClusterInfo(app.ctx)
}

func (app *App) GetBrokersWithLogDirs() ([]kconsole.BrokerWithLogDirs, error) {
	return app.api.ConsoleSvc.GetBrokersWithLogDirs(app.ctx)
}

func (app *App) GetBrokerConfig(brokerId int32) ([]kconsole.BrokerConfigEntry, *rest.Error) {
	return app.api.ConsoleSvc.GetBrokerConfig(app.ctx, brokerId)
}

func (app *App) GetEndpointCompatibility() (kconsole.EndpointCompatibility, error) {
	return app.api.ConsoleSvc.GetEndpointCompatibility(app.ctx)
}

func (app *App) GetConsumerGroupsOverview() ([]kconsole.ConsumerGroupOverview, *rest.Error) {
	return app.api.ConsoleSvc.GetConsumerGroupsOverview(app.ctx, nil)
}
