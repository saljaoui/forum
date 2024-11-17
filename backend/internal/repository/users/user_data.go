package user

import (
	"fmt"

	"forum-project/backend/internal/database"

	"github.com/gofrs/uuid/v5"
)

func emailExists(email string) bool {
	var exists bool
	query := "SELECT EXISTS (select email from user where email=?)"

	err := database.SelectOneRow(query, email).Scan(&exists)
	if err != nil {
		fmt.Println("Error to EXISTS this Email", err)
	}
	return exists
}

func updateUUIDUser(uudi uuid.UUID, userId int64) {
	stm := "UPDATE user SET UUID=? WHERE id=?"
	_, err := database.Exec(stm, uudi, userId)
	if err != nil {
		fmt.Println("Error To Update User uuid")
	}
}

func insertUser(users *User, password string) error {
	stm := "INSERT INTO user (firstname,lastname,email,password) VALUES(?,?,?,?)"
	_, err := database.Exec(stm, users.Firstname, users.Lastname, users.Email, password)
	return err
}

func selectUser(log *Login) *User {
	user := User{}
	query := "select id,email,password, firstname ,lastname FROM user where email=?"
	err := database.SelectOneRow(query, log.Email, log.Password).Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname)
	if err != nil {
		fmt.Println("error to select user", err)
	}
	return &user
}

func checkAuthenticat(id string) bool {
	stm := `SELECT EXISTS (SELECT 1 FROM user WHERE UUID =  ?)  `
	var exists bool
	err := database.SelectOneRow(stm, id, id).Scan(&exists)
	return err == nil
}
