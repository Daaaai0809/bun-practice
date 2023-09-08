package infra

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var cnt = 0

func Connection(config mysql.Config) (*sql.DB, error) {
	for {
		db, err := sql.Open("mysql", config.FormatDSN())
		if err != nil {
			return nil, err
		}

		if err := db.Ping(); err != nil {
			cnt++
			if cnt > 10 {
				return nil, err
			}
			continue
		}

		return db, nil
	}
}