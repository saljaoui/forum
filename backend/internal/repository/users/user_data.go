package repository

import (
	"fmt"

	"forum-project/backend/internal/database"
	"forum-project/backend/internal/models"

	"github.com/gofrs/uuid/v5"
)


func emailExists(email string) bool {
	var exists bool
	query := "SELECT EXISTS (select email from user where email=?)"
	database.SelectOneRow(query, email, &exists)
	return exists
}

func updateuuidUser(uudi uuid.UUID, userId int64) {
	db := database.Config()
	_, err := db.Exec("UPDATE user SET UUID=? WHERE id=?", uudi, userId)
	if err != nil {
		fmt.Println("Error To Update User uuid")
	}
}

func insertUser(users *models.User, password string) error {
	db := database.Config()
	stm := "INSERT INTO user (firstname,lastname,email,password) VALUES(?,?,?,?)"
	_, err := db.Exec(stm, users.Firstname, users.Lastname, users.Email, password)
	return err
}

func selectUser(log *models.Login) *models.User {
	db := database.Config()
	user := models.User{}
	query := "select id,email,password, firstname ,lastname FROM user where email=?"
	err := db.QueryRow(query, log.Email, log.Password).Scan(&user.Id, &user.Email,
		&user.Password, &user.Firstname, &user.Lastname)
	if err != nil {
		fmt.Println("Error ", err)
	}
	return &user
}
