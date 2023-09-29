package api

import (
	"net/http"

	ModelUser "github.com/Sohbetbackend/eMekdep/internal/models"
	"github.com/gin-gonic/gin"
)

// GET ALL USERS
func GetUsers(c *gin.Context) {

	var ResultUsers ModelUser.Users

	allusers := db.QueryRow("SELECT * FROM users").
		Scan(&ResultUsers.ID, &ResultUsers.Firstname, &ResultUsers.Lastname, &ResultUsers.Middlename, &ResultUsers.Username, &ResultUsers.Phone, &ResultUsers.Email, &ResultUsers.Birthday, &ResultUsers.Address)
	if allusers != nil {
		c.JSON(http.StatusNotFound, gin.H{"users": &ResultUsers})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "users": &ResultUsers})
	}
}

// UPDATE USER
func UpdateUser(c *gin.Context) {
	var updateuser ModelUser.Users
	c.BindJSON(&updateuser)
	var ID = c.Param("id")
	var firstname = updateuser.Firstname
	var lastname = updateuser.Lastname
	var middlename = updateuser.Middlename
	var username = updateuser.Username
	var phone = updateuser.Phone
	var email = updateuser.Email
	var birthday = updateuser.Birthday
	var address = updateuser.Address

	_, err := db.Query("UPDATE users SET first_name = ?, last_name = ?, middle_name = ?, user_name = ?, phone = ?, email = ?, birthday = ?, address = ? where id = ?", firstname, lastname, middlename, username, phone, email, birthday, address, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"users": err, "message": "Failed update user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"users": nil, "message": "Success Update User"})
	}
}

// POST NEW USER
func InsertUser(c *gin.Context) {

	var adduser ModelUser.Users
	c.BindJSON(&adduser)
	var firstname = adduser.Firstname
	var lastname = adduser.Lastname
	var middlename = adduser.Middlename
	var username = adduser.Username
	var phone = adduser.Phone
	var email = adduser.Email
	var birthday = adduser.Birthday
	var address = adduser.Address

	_, err := db.Query("INSERT INTO users (first_name, last_name, middle_name, user_name, phone, email, birthday, address) VALUES (?,?,?,?,?,?,?,?)", firstname, lastname, middlename, username, phone, email, birthday, address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": err, "message": "Failed Input User"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": nil, "message": "Success Input User"})
	}
}

func DeleteUser(c *gin.Context) {

	var UserID = c.Param("id")

	_, err := db.Query("DELETE FROM users WHERE id = ?", UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": err, "message": "Failed to delete user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": nil, "message": "Successfully Deleted user"})
	}
}
