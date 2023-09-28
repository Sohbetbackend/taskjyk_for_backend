package api

import (
	"log"
	"net/http"

	ModelUser "github.com/Sohbetbackend/eMekdep/internal/models"
	Conf "github.com/Sohbetbackend/eMekdep/internal/store"
	"github.com/gin-gonic/gin"
)

// GET ALL USERS
func GetUsers(c *gin.Context) {
	var db, err = Conf.Connectdb()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing connection"})
		log.Println("Missing connection")
		return
	}
	defer db.Close()

	var ResultUsers ModelUser.Users
	var allusers = db.QueryRow("SELECT * FROM eMekdep").
		Scan(&ResultUsers.ID, &ResultUsers.Firstname, &ResultUsers.Lastname, &ResultUsers.Middlename, &ResultUsers.Username, &ResultUsers.Phone, &ResultUsers.Email, &ResultUsers.Birthday, &ResultUsers.Address)
	if allusers != nil {
		c.JSON(http.StatusNotFound, gin.H{"users": &ResultUsers})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "users": &ResultUsers})
	}
}

// UPDATE USER
func UpdateUser(c *gin.Context) {
	var db, errdb = Conf.Connectdb()
	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()
	var txtuser ModelUser.Users
	c.BindJSON(&txtuser)
	var ID = c.Param("id")
	var firstname = txtuser.Firstname
	var lastname = txtuser.Lastname
	var middlename = txtuser.Middlename
	var username = txtuser.Username
	var phone = txtuser.Phone
	var email = txtuser.Email
	var birthday = txtuser.Birthday
	var address = txtuser.Address

	_, err := db.Query("UPDATE eMekdep SET first_name = ?, last_name = ?, middle_name = ?, user_name = ?, phone = ?, email = ?, birthday = ?, address = ? where id = ?", firstname, lastname, middlename, username, phone, email, birthday, address, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"users": err, "message": "Failed update user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"users": nil, "message": "Success Update User"})
	}
}

// POST NEW USER
func InsertUser(c *gin.Context) {
	var db, errdb = Conf.Connectdb()
	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()

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

	_, err := db.Query("INSERT INTO eMekdep (first_name, last_name, middle_name, user_name, phone, email, birthday, address) VALUES (?,?,?,?,?,?,?,?)", firstname, lastname, middlename, username, phone, email, birthday, address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": err, "message": "Failed Input User"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": nil, "message": "Success Input User"})
	}
}

func DeleteUser(c *gin.Context) {
	var db, errdb = Conf.Connectdb()
	if errdb != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing Connection"})
		log.Println("Missing Connection")
		return
	}
	defer db.Close()

	var UserID = c.Param("id")

	_, err := db.Query("DELETE FROM eMekdep WHERE id = ?", UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": err, "message": "Failed to delete user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": nil, "message": "Successfully Deleted user"})
	}
}