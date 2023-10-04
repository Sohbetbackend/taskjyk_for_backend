package store

import (
	"fmt"

	"github.com/Sohbetbackend/eMekdep/internal/models"
)

func GetAllUsers() []models.Users {

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("error", err.Error())
		return nil
	}

	users := []models.Users{}

	for results.Next() {
		var user models.Users
		err = results.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Middlename, &user.Username, &user.Phone, &user.Email, &user.Birthday, &user.Address)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)

	}
	return users
}

func UpdateUser(m models.Users) {

	update, err := db.Query("UPDATE users SET first_name = ?, last_name = ?, middle_name = ?, user_name = ?, phone = ?, email = ?, birthday = ?, address = ? where id = ?", m.Firstname, m.Lastname, m.Middlename, m.Username, m.Phone, m.Email, m.Birthday, m.Address, m.ID)
	if err != nil {
		fmt.Println("Err", err.Error())
	}

}

func DeleteUser(m models.Users) {
	delete, err := db.Query("DELETE FROM users WHERE id = ?", m.ID)

	if err != nil {
		fmt.Println("Err", err.Error())

		return
	}
	defer delete.Close()
}

func CreateUser(m models.Users) {
	insert, err := db.Query("INSERT INTO users (first_name, last_name, middle_name, user_name, phone, email, birthday, address) VALUES (?,?,?,?,?,?,?,?)", m.Firstname, m.Lastname, m.Middlename, m.Username, m.Phone, m.Email, m.Birthday, m.Address)

	if err != nil {
		panic(err.Error())
	}
}
