package dbconfig

import (
	"errors"
)

type DBconfig struct {
	Port     int
	Host     string
	User     string
	Password string
	Database string
	Schema   *string
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
