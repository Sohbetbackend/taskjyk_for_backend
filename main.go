package main

import (
	run "github.com/Sohbetbackend/eMekdep/internal/api"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	run.SetupRoutes()
}
