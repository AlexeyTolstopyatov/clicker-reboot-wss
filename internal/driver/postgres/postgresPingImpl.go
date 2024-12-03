package postgres

import (
	"clicker/internal/app"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

// Implementation of 'Connect/Disconnect'
//

// Connect
// Trying to ping PostgresSQL, returns <nil> if result OK,
// overwise error{}
func (pool *Pool) Connect(appInst *app.App) error {
	return pool.Pool.Ping(*appInst.Context)
}

// Disconnect
// Trying to close connection with PostgresSQL
// returns <nil> if result ok
// overwize throws error{} (maybe kill the server...)
func (pool *Pool) Disconnect() {
	pool.Pool.Close()
}

// Reconnect
// Trying to restore connection with PostgresSQL
// throws error{} if connection unrecoverable
// (maybe kill the server...)
func (pool *Pool) Reconnect(appInst *app.App) error {
	for {
		if err := pool.Connect(appInst); err != nil {
			pool.Disconnect()

			if pool.Pool != nil {
				newPool, err := pgxpool.Connect(*appInst.Context, (*appInst).PostgresArgs.ToString())
				if err != nil {
					// error while reconnection attempt was
					time.Sleep(2 * time.Second)
					continue
				}
				pool.Pool = newPool
			}
		}
		time.Sleep(1 * time.Second)
	}
}
