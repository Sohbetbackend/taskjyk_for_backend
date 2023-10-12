package api

import (
	"net/http"

	"github.com/Sohbetbackend/eMekdep/internal/models"
	"github.com/Sohbetbackend/eMekdep/internal/store"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/users", GetUsers)
		v1.PUT("/users/:id", UpdateUser)
		v1.POST("/users", InsertUser)
		v1.DELETE("/users/:id", DeleteUser)
		v1.GET("/classrooms", GetClassrooms)
		v1.POST("/classrooms", CreateClassroom)
		v1.DELETE("/classrooms/:id", DeleteClassroom)
	}
	router.Run(":8080")
}

// GET ALL USERS
func GetUsers(c *gin.Context) {
	resultusers, err := store.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": nil, "message": "Failed to get all user"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "users": resultusers})

}

// UPDATE USER
func UpdateUser(c *gin.Context) {

	updateuser := models.Users{}
	err := c.BindJSON(&updateuser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": nil, "message": "Failed to update user"})
	}
	store.UpdateUser(updateuser)
	c.JSON(http.StatusOK, gin.H{"result": nil, "message": "Success Update User"})
}

// // POST NEW USER
func InsertUser(c *gin.Context) {
	adduser := models.Users{}

	err := c.BindJSON(&adduser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": nil, "message": "Failed Input User"})
	}
	store.CreateUser(adduser)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "users": adduser, "message": "Success Input User"})

}

// // DELETE USER
func DeleteUser(c *gin.Context) {
	var err error

	var deleteuser models.Users

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": nil, "message": "Failed to delete user"})
	}
	store.DeleteUser(deleteuser)
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "users": nil, "message": "Successfully Deleted user"})
}
