package store

import (
	"log"
	"net/http"

	ModelUser "github.com/Sohbetbackend/eMekdep/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	db, err := Connectdb()
	if err != nil {
		log.Println("could not load the database")
	}
	defer db.Close()

	var ResultUsers ModelUser.Users
	allusers := db.QueryRow("SELECT * FROM users").
		Scan(&ResultUsers.ID, &ResultUsers.Firstname, &ResultUsers.Lastname, &ResultUsers.Middlename, &ResultUsers.Username, &ResultUsers.Phone, &ResultUsers.Email, &ResultUsers.Birthday, &ResultUsers.Address)
	if allusers != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": err, "message": "Failed to get users"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "users": &ResultUsers})
	}
}

func UpdateUser(c *gin.Context) {
	db, err := Connectdb()
	if err != nil {
		log.Println("could not load the database")
	}
	defer db.Close()

	var updateuser ModelUser.Users
	c.BindJSON(&updateuser)
	var ID = c.Param("id")

	update, err := db.Query("UPDATE users SET firstname = ?, lastname = ?, middlename = ?, username = ?, phone = ?, email = ?, birthday = ?, address = ? where id = ?", updateuser.Firstname, updateuser.Lastname, updateuser.Middlename, updateuser.Username, updateuser.Phone, updateuser.Email, updateuser.Birthday, updateuser.Address, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": err, "message": "Failed update user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "users": &updateuser, "message": "Success Update User"})
	}
	defer update.Close()
}

func DeleteUser(c *gin.Context) {
	db, err := Connectdb()
	if err != nil {
		log.Println("could not load the database")
	}
	defer db.Close()

	var UserID = c.Param("id")
	delete, err := db.Query("DELETE FROM eMekdep WHERE id = ?", UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": err, "message": "Failed to delete user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "users": nil, "message": "Successfully Deleted user"})
	}
	defer delete.Close()
}

func CreateUser(c *gin.Context) {
	db, err := Connectdb()
	if err != nil {
		log.Println("could not load the database")
	}
	defer db.Close()

	var adduser ModelUser.Users
	c.BindJSON(&adduser)
	insert, err := db.Query("INSERT INTO eMekdep (first_name, last_name, middle_name, user_name, phone, email, birthday, address) VALUES (?,?,?,?,?,?,?,?)", adduser.Firstname, adduser.Lastname, adduser.Middlename, adduser.Username, adduser.Phone, adduser.Email, adduser.Birthday, adduser.Address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "users": err, "message": "Failed Input User"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "users": &adduser, "message": "Success Input User"})
	}
	defer insert.Close()
}
