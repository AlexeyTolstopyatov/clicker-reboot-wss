package main

import (
	"clicker/internal/app"
	"clicker/internal/requirements"
	"context"
	"fmt"
)

func main() {
	appContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		server   = requirements.AppConfig{}
		postgres = requirements.DbConfig{}
	)

	if err := server.LoadConfig(); err != nil {
		fmt.Printf("Failed to load app config: %v\n", err)
		fmt.Printf("Use command line arguments or environment variables (*.json)\n")
		fmt.Printf("Loading default configuration file\n")
		server = *app.DefaultInstance(&appContext).ServerArgs
	}

	if err := postgres.LoadConfig(); err != nil {
		fmt.Printf("Failed to load db config: %v\n", err)
		fmt.Printf("Use command line arguments or environment variables (*.json)\n")
		fmt.Printf("Loading default configuration file\n")
		postgres = *app.DefaultInstance(&appContext).PostgresArgs
	}

	instance := app.App{
		Context:      &appContext,
		ServerArgs:   &server,
		PostgresArgs: &postgres,
	}

	instance.Run()
}
