package main

import (
	Cuser "github.com/Sohbetbackend/eMekdep/internal/api"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/users", Cuser.GetUsers)
		v1.PUT("/users/:id", Cuser.UpdateUser)
		v1.POST("/users", Cuser.InsertUser)
		v1.DELETE("/users/:id", Cuser.DeleteUser)
	}
	router.Run(":8080")

}
