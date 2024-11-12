package repository

import (
	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
	"forum-project/backend/internal/models"
)

func Post(users *models.Post) messages.Messages {
	message := messages.Messages{}
	insert := "INSERT INTO card (user_id,content,created_at) VALUES(?,?,?)"
	// exists := emailExists(users.Email)
	// if exists {
	// 	message.MessageError = "Email User Is Allready Exsist"
	// 	message.ErrorBool = true
	// } else {
	// 	password := hashPassword(&models.User{})
	// 	stm := "INSERT INTO users (firstname,lastname,email,password) VALUES(?,?,?,?)"
	// 	fmt.Println(password)
	database.Exec(stm, users.Firstname, users.Lastname, users.Email, password)
	// 	message.MessageSucc = "User created successfully"
	// 	// check := checkPassword(users.Password, password)
	// }
	return message
}
