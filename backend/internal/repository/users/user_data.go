package repository

import (
	"fmt"

	"forum-project/backend/internal/database"

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

func insertUser(users *User, password string) error {
	db := database.Config()

	stm := "INSERT INTO user (firstname,lastname,email,password) VALUES(?,?,?,?)"
	_, err := db.Exec(stm, users.Firstname, users.Lastname, users.Email, password)
	return err
}

func selectUser(log *Login) *User {
	db := database.Config()
	user := User{}
	query := "select id,email,password, firstname ,lastname FROM user where email=?"
	err := db.QueryRow(query, log.Email, log.Password).Scan(&user.Id, &user.Email,
		&user.Password, &user.Firstname, &user.Lastname)
	if err != nil {
		fmt.Println("Error ", err)
	}
	return &user
}

func checkAuthenticat(id string) bool {
	db := database.Config()
	stm := `SELECT 
		EXISTS (SELECT 1 FROM user WHERE UUID =  ?) as ex,
		(SELECT UUID FROM user WHERE UUID =?)  as  uuid`
	var exists bool
	var uuid string
	err := db.QueryRow(stm, id, id).Scan(&exists, &uuid)
	fmt.Println(exists, uuid)
	return err == nil
}
