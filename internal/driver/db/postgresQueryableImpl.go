package db

import (
	"context"
	"github.com/jackc/pgx/v4"
)

//
// Here contains data query language or many kinds of one SQL instruction
//

func (pool *Pool) SelectRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	if tx := GetTransaction(ctx); tx != nil {
		return tx.QueryRow(ctx, query, args...)
	}

	return pool.Pool.QueryRow(ctx, query, args...)
}

func (pool *Pool) SelectRows(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	if tx := GetTransaction(ctx); tx != nil {
		return tx.Query(ctx, query, args...)
	}

	return pool.Pool.Query(ctx, query, args...)
}
