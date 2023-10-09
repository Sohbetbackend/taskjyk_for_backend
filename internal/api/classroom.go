package api

import (
	"net/http"
	"strconv"

	"github.com/Sohbetbackend/eMekdep/internal/models"

	"github.com/Sohbetbackend/eMekdep/internal/store"
	"github.com/gin-gonic/gin"
)

func GetClassrooms(c *gin.Context) {
	classrooms, err := store.GetAllClasses()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "classrooms": nil, "message": "Failed to get all user"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "classrooms": &classrooms})
}

func CreateClassroom(c *gin.Context) {
	addclassroom := models.Classroom{}

	err := c.BindJSON(&addclassroom)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "classroom": nil, "message": "Failed to create class"})
	}
	store.InsertClass(addclassroom)
	c.JSON(http.StatusOK, gin.H{"success": true, "classroom": addclassroom, "message": "Successfully added class"})
}

func DeleteClassroom(c *gin.Context) {
	var err error
	var deleteclassroom models.Classroom

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "classroom": nil, "message": "Failed to delete class"})
	}
	store.DeleteClass(deleteclassroom)
	c.JSON(http.StatusOK, gin.H{"success": true, "classroom": nil, "message": "Successfully Deleted Class"})
}

func GetClassroom(c *gin.Context) {
	id := c.Param("id")
	id2, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	room := store.GetClass(id2)

	if room != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "classroom": nil, "message": "Failed to get classroom"})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "classroom": room})
}
