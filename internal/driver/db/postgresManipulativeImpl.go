package db

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

//
// Here contains data manipulation language
// or many kinds of SQL Data manipulating instructions
//

// BeginContextAction
// Starts a transaction, executes an action (argument).
// Depending on the returned value of the function,
// commits or rolls back the transaction
func (pool *Pool) BeginContextAction(ctx context.Context, options pgx.TxOptions, action func(pgx.Tx) error) error {
	return pool.Pool.BeginTxFunc(ctx, options, action)
}

// Execute
// Executes query, without any results
func (pool *Pool) Execute(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	if tx := GetTransaction(ctx); tx != nil {
		return tx.Exec(ctx, sql, arguments...)
	}

	return pool.Pool.Exec(ctx, sql, arguments...)
}
