package store

import (
	"log"

	"github.com/Sohbetbackend/eMekdep/internal/models"
)

func GetAllClasses() ([]models.Classroom, error) {
	result, err := db.Query("SELECT * FROM classroom")
	if err != nil {
		log.Fatal("err", err.Error())
	}

	classroom := []models.Classroom{}

	for result.Next() {
		var room models.Classroom
		err = result.Scan(&room.ID, &room.Name, &room.Desc)

		if err != nil {
			panic(err.Error())
		}
		classroom = append(classroom, room)
	}

	return classroom, err
}

func InsertClass(m models.Classroom) {
	_, err := db.Query("INSERT INTO classroom (name, description) VALUES (?,?)", m.Name, m.Desc)

	if err != nil {
		log.Fatal("Err", err.Error())
	}
}

func DeleteClass(m models.Classroom) {
	_, err := db.Query("DELETE FROM classroom WHERE id = ?", m.ID)

	if err != nil {
		log.Fatal("Err", err.Error())
	}
}

func GetClass(id int) *models.Classroom {
	room := &models.Classroom{}

	getclassroom, err := db.Query("SELECT * FROM classroom WHERE id=?", id)

	if err != nil {
		log.Fatal("err", err.Error())
	}

	if getclassroom.Next() {
		err = getclassroom.Scan(&room.Name, &room.Desc)

		if err != nil {
			return nil
		}
	}
	return room
}

func UpdateClass(m models.Classroom) {

	_, err := db.Query("UPDATE classroom SET name = ?, description = ?", m.Name, m.Desc)
	if err != nil {
		log.Fatal("Error", err.Error())
	}
}
