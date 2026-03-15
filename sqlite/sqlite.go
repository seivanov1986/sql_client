package sqlite

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/seivanov1986/sql_client"
	"net/url"

	_ "github.com/mutecomm/go-sqlcipher/v4"
)

func NewClient(path string) (*sql_client.DataBaseImpl, error) {
	// здесь НЕ используем _pragma_key
	conn, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &sql_client.DataBaseImpl{DB: conn}, nil
}

func NewEncryptedClient(path, key string) (*sql_client.DataBaseImpl, error) {
	escapedKey := url.QueryEscape(key)
	dsn := fmt.Sprintf(
		"%s?_pragma_key=%s&_pragma_cipher_page_size=4096",
		path,
		escapedKey,
	)

	conn, err := sqlx.Connect("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return &sql_client.DataBaseImpl{DB: conn}, nil
}
