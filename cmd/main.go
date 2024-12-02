package main

import (
	"clicker/internal/handler/httph"
	"clicker/internal/requirements"
	"context"
)

func main() {
	appContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		server   = requirements.Server{}
		postgres = requirements.Postgres{}
	)

	if err := server.LoadServer(); err != nil {
		panic(err.Error())
	}

	if err := postgres.LoadPostgres(); err != nil {
		panic(err.Error())
	}

	httph.Run(&appContext)
}
