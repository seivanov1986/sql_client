package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/seivanov1986/sql_client"

	_ "github.com/mattn/go-sqlite3"
)

func NewClient(dataSourcePath string) (*sql_client.DataBaseImpl, error) {
	conn, err := sqlx.Connect("sqlite3", dataSourcePath)
	if err != nil {
		return nil, err
	}

	return &sql_client.DataBaseImpl{DB: conn}, nil
}
