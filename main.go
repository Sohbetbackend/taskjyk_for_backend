package main

import (
	"log"

	"github.com/Sohbetbackend/eMekdep/internal/api"
	"github.com/Sohbetbackend/eMekdep/internal/store"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := store.Connectdb()
	if err != nil {
		log.Fatal(err)
	}

	api.SetupRoutes()

}
