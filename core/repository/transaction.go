package repository

import (
	"context"
	"database/sql"
)

type TransactionalRepository interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (TransactionalRepository, error)
	Commit() error
	Rollback() error
	UseTx(tx TransactionalRepository)
}
