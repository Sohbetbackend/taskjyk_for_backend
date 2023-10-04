package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connectdb() error {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/eMekdep")
	if err != nil {
		return nil
	}
	defer db.Close()
	err = db.Ping()
	return err
}
