package app

import (
	"clicker/internal/requirements"
	"context"
)

type App struct {
	Context      *context.Context
	ServerArgs   *requirements.Server
	PostgresArgs *requirements.Postgres
}

// Instance
// Runs application's instance using known external
// options. If you don't know server/db arguments, use DefaultInstance
// function.
func Instance(
	ctx *context.Context,
	serverArgs *requirements.Server,
	postgresArgs *requirements.Postgres) *App {
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
		ServerArgs: &requirements.Server{
			Host: "localhost",
			Port: "5431",
		},
		PostgresArgs: &requirements.Postgres{
			User:     "postgres",
			Password: "postgres",
			Name:     "clickerdb",
			Ssl:      false,
		},
	}
}

func (a *App) Run() {

}
