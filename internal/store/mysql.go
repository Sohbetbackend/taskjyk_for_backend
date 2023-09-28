package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connectdb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/eMekdep")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return db, err
}
