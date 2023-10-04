package api

import (
	"net/http"

	"github.com/Sohbetbackend/eMekdep/internal/models"
	"github.com/Sohbetbackend/eMekdep/internal/store"
	Conf "github.com/Sohbetbackend/eMekdep/internal/store"
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
	}
	router.Run(":8080")
}

// GET ALL USERS
func GetUsers(c *gin.Context) {
	var err error
	ResultUsers := Conf.GetAllUsers()

	if ResultUsers == nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": err, "message": "Failed update user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "users": ResultUsers})
	}
}

// UPDATE USER
func UpdateUser(c *gin.Context) {
	var err error
	var ID = c.Param("id")

	var updateuser models.Users
	c.BindJSON(&updateuser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": err, "message": "Failed update user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "users": &updateuser, "message": "Success Update User"})
	}
}

// POST NEW USER
func InsertUser(c *gin.Context) {

	var adduser models.Users

	if err := c.BindJSON(&adduser); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": nil, "message": "Failed Input User"})
	} else {
		store.CreateUser(adduser)
		c.IndentedJSON(http.StatusOK, gin.H{"success": true, "users": adduser, "message": "Success Input User"})
	}
}

// DELETE USER
func DeleteUser(c *gin.Context) {
	var err error
	var UserID = c.Param("id")

	user := Conf.DeleteUser(UserID)

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": err, "message": "Failed to delete user"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"success": true, "users": nil, "message": "Successfully Deleted user"})
	}
}
