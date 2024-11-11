package repository

import (
	"fmt"

	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
	"forum-project/backend/internal/models"
)

type UserModel struct{}

func Register(users *models.User) messages.UserAllReadyExists {
	// mess := messages.Message()
	open := database.Config()
	tes := messages.UserAllReadyExists{}
	exists := emailExists(users.Email)
	if exists {
		tes.MessageError = "Email User Is Allready Exsist"
		tes.ErrorBool = true
	} else {
		stm := "INSERT INTO users (firstname,lastname,email,password) VALUES(?,?,?,?)"
		_, err := open.Exec(stm, users.Firstname, users.Lastname, users.Email, users.Password)
		if err != nil {
			fmt.Println(err)
		}
		tes.MessageSucc = "User created successfully"
		tes.SuccBool = true
	}
	return tes
}

func emailExists(email string) bool {
	open := database.Config()
	var exists bool
	query := "SELECT EXISTS (select email from users where email=?)"
	err := open.QueryRow(query, email).Scan(&exists)
	if err != nil {
		fmt.Println(err)
	}
	return exists
}
