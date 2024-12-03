package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
)

// Here is seeking properties
// which detecting operations through application's context

type TransactionDetectable interface {
	GetTransaction(ctx context.Context) (pgx.Tx, error)
	SetTransaction(ctx context.Context, tx pgx.Tx) error
}

type TransactionKey struct{}

// SetTransaction
// Inserts transaction to application's context
func SetTransaction(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TransactionKey{}, tx)
}

// GetTransaction
// Seeking transaction using application's context
func GetTransaction(ctx context.Context) pgx.Tx {
	if tx, ok := ctx.Value(TransactionKey{}).(pgx.Tx); ok {
		return tx
	}
	return nil
}
