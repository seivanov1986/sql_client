package mysql

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/seivanov1986/sql_client"

	_ "github.com/go-sql-driver/mysql"
)

type DBconfig struct {
	Port     int
	Host     string
	User     string
	Password string
	Database string
}

func (d *DBconfig) Validate() error {
	if d.Port <= 0 || d.Port > 65535 {
		return errors.New("port field is out of the available range")
	}
	if d.Host == "" {
		return errors.New("host field is empty")
	}
	if d.User == "" {
		return errors.New("user field is empty")
	}
	if d.Password == "" {
		return errors.New("password field is empty")
	}
	if d.Database == "" {
		return errors.New("database field is empty")
	}
	return nil
}

func NewClient(cfg *DBconfig) (*sql_client.DataBaseImpl, error) {
	source := fmt.Sprintf(
		"%v:%v@(%v:%v)/%v",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	conn, err := sqlx.Connect("mysql", source)
	if err != nil {
		return nil, err
	}

	return &sql_client.DataBaseImpl{DB: conn}, nil
}
