package sql_client

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type key string

const (
	trx key = "trx"
)

func (d *DataBaseImpl) NewTransaction() (*sqlxTransaction, error) {
	tx, _ := d.DB.Beginx()
	return &sqlxTransaction{tx}, nil
}

func (d *DataBaseImpl) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.DB.SelectContext(ctx, dest, query, args...)
}

func (d *DataBaseImpl) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.DB.ExecContext(ctx, query, args...)
}

func (d *DataBaseImpl) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return d.DB.NamedExecContext(ctx, query, arg)
}

func (d *DataBaseImpl) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.DB.GetContext(ctx, dest, query, args...)
}

func (d *DataBaseImpl) DeleteIn(ctx context.Context, query string, args ...interface{}) error {
	query, inArgs, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}

	_, err = d.DB.ExecContext(ctx, query, inArgs...)
	return err
}

func (t *sqlxTransaction) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.TX.SelectContext(ctx, dest, query, args...)
}

func (t *sqlxTransaction) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return t.TX.ExecContext(ctx, query, args...)
}

func (t *sqlxTransaction) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return t.TX.NamedExecContext(ctx, query, arg)
}

func (t *sqlxTransaction) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.TX.GetContext(ctx, dest, query, args...)
}

func (t *sqlxTransaction) DeleteIn(ctx context.Context, query string, args ...interface{}) error {
	query, inArgs, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}

	_, err = t.TX.ExecContext(ctx, query, inArgs...)
	return err
}

func (tr *transactionManager) MakeTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	transaction, err := tr.db.NewTransaction()
	if err != nil {
		return err
	}

	err = fn(context.WithValue(ctx, trx, transaction))
	if err != nil {
		transaction.TX.Rollback()
		return err
	}

	return transaction.TX.Commit()
}

func (tr *transactionManager) FindTransaction(ctx context.Context) *sqlxTransaction {
	transaction := ctx.Value(trx)
	result, ok := transaction.(*sqlxTransaction)
	if !ok {
		return nil
	}

	return result
}
