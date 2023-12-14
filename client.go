package sql_client

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type sqlxTransaction struct {
	*sqlx.Tx
}

func (d *dataBase) NewTransaction() (*sqlxTransaction, error) {
	tx, _ := d.db.Beginx()
	return &sqlxTransaction{tx}, nil
}

func (d *dataBase) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.db.SelectContext(ctx, dest, query, args...)
}

func (d *dataBase) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

func (d *dataBase) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return d.db.NamedExecContext(ctx, query, arg)
}

func (d *dataBase) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.db.GetContext(ctx, dest, query, args...)
}

func (d *dataBase) DeleteIn(ctx context.Context, query string, args ...interface{}) error {
	query, inArgs, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}

	_, err = d.db.ExecContext(ctx, query, inArgs...)
	return err
}

func (d *sqlxTransaction) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.SelectContext(ctx, dest, query, args...)
}

func (d *sqlxTransaction) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.ExecContext(ctx, query, args...)
}

func (d *sqlxTransaction) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return d.NamedExecContext(ctx, query, arg)
}

func (d *sqlxTransaction) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.GetContext(ctx, dest, query, args...)
}

func (d *sqlxTransaction) DeleteIn(ctx context.Context, query string, args ...interface{}) error {
	query, inArgs, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}

	_, err = d.ExecContext(ctx, query, inArgs...)
	return err
}

func (tr *transactionManager) MakeTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	trx, err := tr.db.NewTransaction()
	if err != nil {
		return err
	}

	err = fn(context.WithValue(ctx, "trx", trx))
	if err != nil {
		trx.Rollback()
		return err
	}

	trx.Commit()
	return nil
}

func (tr *transactionManager) FindTransaction(ctx context.Context) *sqlxTransaction {
	trx := ctx.Value("trx")
	result, ok := trx.(*sqlxTransaction)
	if !ok {
		return nil
	}

	return result
}
