package api

import (
	Cuser "github.com/Sohbetbackend/eMekdep/internal/store"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/users", Cuser.GetAllUsers)
		v1.PUT("/users/:id", Cuser.UpdateUser)
		v1.POST("/users", Cuser.CreateUser)
		v1.DELETE("/users/:id", Cuser.DeleteUser)
	}
}
