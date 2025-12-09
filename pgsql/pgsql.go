package pgsql

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/seivanov1986/sql_client"
	dbconfig "github.com/seivanov1986/sql_client/config"

	_ "github.com/lib/pq"
)

const defaultSchema = "public"

func NewClient(cfg *dbconfig.DBconfig) (*sql_client.DataBaseImpl, error) {
	var schema string = defaultSchema
	if cfg.Schema != nil {
		schema = *cfg.Schema
	}

	source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		schema,
	)

	conn, err := sqlx.Connect("postgres", source)
	if err != nil {
		return nil, err
	}

	return &sql_client.DataBaseImpl{DB: conn}, nil
}
