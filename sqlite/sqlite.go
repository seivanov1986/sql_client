package sqlite

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/seivanov1986/sql_client"
	"net/url"

	_ "github.com/covenantsql/go-sqlite3-encrypt"
)

func NewClient(dataSourcePath string) (*sql_client.DataBaseImpl, error) {
	conn, err := sqlx.Connect("sqlite3", dataSourcePath)
	if err != nil {
		return nil, err
	}

	return &sql_client.DataBaseImpl{DB: conn}, nil
}

func NewEncryptedClient(dataSourcePath, key string) (*sql_client.DataBaseImpl, error) {
	escapedPath := url.QueryEscape(dataSourcePath)
	escapedKey := url.QueryEscape(key)

	dsn := fmt.Sprintf(
		"file:%s?_crypto_key=%s&_pragma_cipher_page_size=4096",
		escapedPath,
		escapedKey,
	)

	conn, err := sqlx.Connect("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return &sql_client.DataBaseImpl{DB: conn}, nil
}
