package user

import (
	"database/sql"
	"fmt"

	"forum-project/backend/internal/database"
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

func updateUUIDUser(uudi string, userId int64) error {
	stm := "UPDATE user SET UUID=? WHERE id=?"
	_, err := database.Exec(stm, uudi, userId)
	return err
}

func insertUser(users *User, password string) ( sql.Result,error) {
	stm := "INSERT INTO user (firstname,lastname,email,password) VALUES(?,?,?,?)"
	row, err := database.Exec(stm, users.Firstname, users.Lastname, users.Email, password)
	return row,err 
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

func CheckUser(id int) bool {
	stm := `SELECT EXISTS (SELECT 1 FROM user WHERE id =  ?)  `
	var exists bool
	err := database.SelectOneRow(stm, id, id).Scan(&exists)
	return err == nil
}