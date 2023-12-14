package sql_client

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

type dataBase struct {
	db *sqlx.DB
}

func New() *dataBase {
	host := os.Getenv("SQL_HOST")
	port := os.Getenv("SQL_PORT")
	user := os.Getenv("SQL_USER")
	pass := os.Getenv("SQL_PASSWD")
	dbname := os.Getenv("SQL_DBNAME")

	source := fmt.Sprintf(
		"%v:%v@(%v:%v)/%v",
		user, pass, host, port, dbname,
	)

	conn, err := sqlx.Connect("mysql", source)
	if err == nil {
		conn.SetConnMaxLifetime(60 * time.Second)
		conn.SetMaxIdleConns(10)
		conn.SetMaxOpenConns(10)
	}
	return &dataBase{
		db: conn,
	}
}
