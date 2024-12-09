package txhost

import (
	"clicker/internal/driver/db"
	"github.com/jackc/pgx/v4"
	"golang.org/x/net/context"
)

//
// Transaction's Host
// Uses server's PostgreSQL driver ".../internal/driver/db" API for controlling transactions
//

type (
	Transactable interface {
		WithinTransaction(ctx context.Context, transactAction func(context context.Context) error) error
	}

	TransactHost struct {
		client db.Pool
	}
)

func Instance(client db.Pool) *TransactHost {
	return &TransactHost{client: client}
}

func (t *TransactHost) WithinTransaction(ctx context.Context, transactAction func(context context.Context) error) error {
	return t.client.BeginContextAction(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return transactAction(db.SetTransaction(ctx, tx))
	})
}
