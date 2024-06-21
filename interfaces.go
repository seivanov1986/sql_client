package sql_client

import (
	"context"
	"database/sql"
	"io/fs"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

type DataBaseMethods interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	DeleteIn(ctx context.Context, query string, args ...interface{}) error
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
}

type DataBase interface {
	DataBaseMethods
	GetDB() *sqlx.DB
	NewTransaction() (*sqlxTransaction, error)
	RunMigrations(l goose.Logger, migrationFiles fs.FS) error
	Close() error
}

type Transaction interface {
	DataBaseMethods
	Rollback() error
	Commit() error
}

type TransactionManager interface {
	MakeTransaction(ctx context.Context, fn func(ctx context.Context) error) error
	FindTransaction(ctx context.Context) *sqlxTransaction
	DefaultTrOrDB(ctx context.Context, db DataBaseMethods) DataBaseMethods
}
