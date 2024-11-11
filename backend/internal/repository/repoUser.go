package repository

import (
	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
	"forum-project/backend/internal/models"
)

type UserModel struct{}

func Register(users *models.User) messages.UserAllReadyExists {
	tes := messages.UserAllReadyExists{}
	exists := emailExists(users.Email)
	if exists {
		tes.MessageError = "Email User Is Allready Exsist"
		tes.ErrorBool = true
	} else {
		stm := "INSERT INTO users (firstname,lastname,email,password) VALUES(?,?,?,?)"
		database.Exec(stm, users.Firstname, users.Lastname, users.Email, users.Password)
		tes.MessageSucc = "User created successfully"
		tes.SuccBool = true
	}
	return tes
}

func emailExists(email string) bool {
	var exists bool
	query := "SELECT EXISTS (select email from users where email=?)"
	database.SelectOneRow(query, email, &exists)
	return exists
}
