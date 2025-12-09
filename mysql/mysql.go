package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/seivanov1986/sql_client"
	dbconfig "github.com/seivanov1986/sql_client/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewClient(cfg *dbconfig.DBconfig) (*sql_client.DataBaseImpl, error) {
	source := fmt.Sprintf(
		"%v:%v@(%v:%v)/%v?parseTime=%t",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.ParseTime,
	)

	conn, err := sqlx.Connect("mysql", source)
	if err != nil {
		return nil, err
	}

	return &sql_client.DataBaseImpl{DB: conn}, nil
}
