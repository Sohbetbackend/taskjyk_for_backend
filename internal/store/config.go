package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connectdb() error {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/emekdep")
	if err != nil {
		return err
	}
	err = db.Ping()
	return err
}
