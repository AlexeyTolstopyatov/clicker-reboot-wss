package postgres

import (
	"clicker/internal/app"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Pool struct {
	Pool *pgxpool.Pool
}

type Postgres interface {
	Connect(appInst *app.App) error
	Disconnect()
	Reconnect(appInst *app.App) error
	// Here is main driver's function
	// for controlling a connection with
	// database layer in OS
}

type DataQueryable interface {
	// Here contains data query language
	// or many kinds of one SQL instruction
	// [SELECT]
}

type DataManipulative interface {
	// Here is data manipulation language
	// or kinds of data-manipulating instruction
	// adopted.
	// INSERT UPDATE DELETE etc.
}

// Instance
// Return new Instance of PostgresSQL driver
// and provides API through Pool structure
func Instance(appInst *app.App) *Pool {
	return &Pool{}
}
