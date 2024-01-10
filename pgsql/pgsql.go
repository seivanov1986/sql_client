package pgsql

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/seivanov1986/sql_client"
	dbconfig "github.com/seivanov1986/sql_client/config"

	_ "github.com/lib/pq"
)

func NewClient(cfg *dbconfig.DBconfig) (*sql_client.DataBaseImpl, error) {
	source := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)

	conn, err := sqlx.Connect("postgres", source)
	if err != nil {
		return nil, err
	}

	return &sql_client.DataBaseImpl{DB: conn}, nil
}
