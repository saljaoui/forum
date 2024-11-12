package repository

import (
	"fmt"

	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
	"forum-project/backend/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct{}

func Register(users *models.User) messages.Messages {
	message := messages.Messages{}
	exists := emailExists(users.Email)
	if exists {
		message.MessageError = "Email User Is Allready Exsist"
		message.ErrorBool = true
	} else {
		password := hashPassword(&models.User{})
		stm := "INSERT INTO users (firstname,lastname,email,password) VALUES(?,?,?,?)"
		fmt.Println(password)
		database.Exec(stm, users.Firstname, users.Lastname, users.Email, password)
		message.MessageSucc = "User created successfully"
		// check := checkPassword(users.Password, password)
	}
	return message
}

// func Login(user *models.User){
// 	query:="select * from users where email=?"
// }

func hashPassword(pass *models.User) string {
	haspassword, err := bcrypt.GenerateFromPassword([]byte(pass.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error", err)
	}
	pass.Password = string(haspassword)
	return pass.Password
}

func checkPassword(passwordUser, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordUser), []byte(password))
	return err == nil
}

func emailExists(email string) bool {
	var exists bool
	query := "SELECT EXISTS (select email from users where email=?)"
	database.SelectOneRow(query, email, &exists)
	return exists
}

func DisplyInfoUser(id models.User) {
}
