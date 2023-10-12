package store

import (
	"log"

	"github.com/Sohbetbackend/eMekdep/internal/models"
)

func GetAllUsers() ([]models.Users, error) {
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("error", err.Error())
	}

	users := []models.Users{}

	for results.Next() {
		var user models.Users
		err = results.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Middlename, &user.Username, &user.Phone, &user.Email, &user.Birthday, &user.Gender, &user.Address)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)

	}
	return users, err
}

func DeleteUser(m models.Users) {

	_, err := db.Query("DELETE FROM users WHERE id = ?", m.ID)
	if err != nil {
		log.Fatal("Err", err.Error())
	}

}

func CreateUser(m models.Users) {
	_, err := db.Query("INSERT INTO users (first_name, last_name, middle_name, user_name, phone, email, birthday, gender, address) VALUES (?,?,?,?,?,?,?,?,?)", m.Firstname, m.Lastname, m.Middlename, m.Username, m.Phone, m.Email, m.Birthday, m.Gender, m.Address)

	if err != nil {
		log.Fatal("Err", err.Error())
	}
}

func UpdateUser(m models.Users) {
	_, err := db.Query("UPDATE users SET first_name = ?,last_name = ?, middle_name = ?, user_name = ?, phone = ?, email = ?, birthday = ?, address = ? where user_id = ?", m.Username, m.Lastname, m.Middlename, m.Username, m.Phone, m.Email, m.Birthday, m.Address, m.ID)

	if err != nil {
		log.Fatal("Err", err.Error())
	}
}
