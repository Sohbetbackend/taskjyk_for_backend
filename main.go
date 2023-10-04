package main

import (
	"github.com/Sohbetbackend/eMekdep/internal/api"
	"github.com/Sohbetbackend/eMekdep/internal/store"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	api.SetupRoutes()
	store.Connectdb()
}
