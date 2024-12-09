package app

import (
	"clicker/internal/requirements"
	"context"
)

type App struct {
	Context      *context.Context
	ServerArgs   *requirements.AppConfig
	PostgresArgs *requirements.DbConfig
}

// Instance
// Runs application's instance using known external
// options. If you don't know server/db arguments, use DefaultInstance
// function.
func Instance(
	ctx *context.Context,
	serverArgs *requirements.AppConfig,
	postgresArgs *requirements.DbConfig) *App {
	return &App{
		Context:      ctx,
		ServerArgs:   serverArgs,
		PostgresArgs: postgresArgs,
	}
}

// DefaultInstance
// Runs application's instance without external options or .env file
// based configuration.
func DefaultInstance(ctx *context.Context) *App {
	return &App{
		Context: ctx,
		ServerArgs: &requirements.AppConfig{
			Host: "localhost",
			Port: "5431",
		},
		PostgresArgs: &requirements.DbConfig{
			User:     "dbadmin",
			Password: "dbadmin",
			Name:     "clickerdb",
			Ssl:      "--verify-ca",
		},
	}
}

func (a *App) Run() {
	// run http adapter,
	// run ws adapter,
	// ws adapter -> models instance.

}
