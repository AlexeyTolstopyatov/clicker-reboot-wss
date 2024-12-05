package db

import (
	"clicker/internal/app"
	"clicker/internal/util"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Pool struct {
	Pool *pgxpool.Pool
}

type DataFundamentals interface {
	Connect(appInst *app.App) error
	Disconnect()
	Reconnect(appInst *app.App) error
	// Here is main driver's function
	// for controlling a connection with
	// database layer in OS
}

type DataQueryable interface {
	SelectRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	SelectRows(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	// Here contains data query language
	// or many kinds of one SQL instruction
	// [SELECT]
}

type DataController interface {
	Execute(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	BeginContextAction(ctx context.Context, options pgx.TxOptions, action func(tx pgx.Tx) error) error
	// Here is data manipulation language
	// or kinds of data-manipulating instruction
	// adopted.
	// INSERT UPDATE DELETE etc.
}

// Instance
// Return new Instance of PostgresSQL driver
// and provides API through Pool structure
func Instance(appInst *app.App, attempts int) (*Pool, error) {
	pool := pgxpool.Pool{}
	err := util.ExecuteRepeat(func() error {
		actionctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(3)*time.Second)
		defer cancel() // lambda executes independ of cancel()

		pool, err := pgxpool.Connect(
			actionctx,
			appInst.PostgresArgs.ToString())

		if err != nil {
			return err // CONNECT error
		}

		if err := pool.Ping(actionctx); err != nil {
			return err // PING error
		}

		return nil
	}, attempts, time.Second*2)

	return &Pool{Pool: &pool}, err
}
