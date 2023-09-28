package api

import (
	"log"
	"net/http"

	ModelUser "github.com/Sohbetbackend/eMekdep/internal/models"
	Conf "github.com/Sohbetbackend/eMekdep/internal/store"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var db, err = Conf.Connectdb()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": "Missing connection"})
		log.Println("Missing connection")
		return
	}
	defer db.Close()

	var ResultUsers ModelUser.Users
	var allusers = db.QueryRow("SELECT * FROM users").
		Scan(&ResultUsers.ID, &ResultUsers.Firstname, &ResultUsers.Lastname, &ResultUsers.Middlename, &ResultUsers.Username, &ResultUsers.Phone, &ResultUsers.Email, &ResultUsers.Birthday, &ResultUsers.Address)
	if allusers != nil {
		c.JSON(http.StatusNotFound, gin.H{"users": &ResultUsers})
	} else {
		c.JSON(http.StatusOK, gin.H{"users": &ResultUsers})
	}
}

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

	_, err := db.Query("UPDATE users SET firstname = ?, lastname = ?, middlename = ?, username = ?, phone = ?, email = ?, birthday = ?, address = ? where id = ?", firstname, lastname, middlename, username, phone, email, birthday, address, ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"users": err, "message": "Failed update user"})
	} else {
		c.JSON(http.StatusOK, gin.H{"users": nil, "message": "Success Update User"})
	}
}
