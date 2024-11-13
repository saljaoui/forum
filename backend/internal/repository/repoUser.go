package repository

import (
	"fmt"

	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
	"forum-project/backend/internal/models"

	"github.com/gofrs/uuid/v5"
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
		stm := "INSERT INTO user (firstname,lastname,email,password,UUID) VALUES(?,?,?,?,?)"
		uuid, err := uuid.NewV4()
		if err != nil {
			fmt.Println("Error to Generate uuid", err)
		}
		database.Exec(stm, users.Firstname, users.Lastname, users.Email, password, uuid)
		message.MessageSucc = "User created successfully"
		// check := checkPassword(users.Password, password)
	}
	return message
}

func Login(user *models.Login) {
	db := database.Config()
	passwordhased := ""
	id := 0
	email := ""
	query := "select id,email,password from user where email=? and password=?"

	err := db.QueryRow(query, user.Email, user.Password).Scan(&id, &email, &passwordhased)
	if err != nil {
		fmt.Println("errror", err)
	}
	if checkPassword(user.Password, passwordhased) {
	} else {
		fmt.Println("Your Password Encorect")
	}
	fmt.Println(id, email, passwordhased)
}

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
	query := "SELECT EXISTS (select email from user where email=?)"
	database.SelectOneRow(query, email, &exists)
	return exists
}

func DisplyInfoUser(id models.User) {
	rows := database.SelectRows("select * from post", models.User{})
	for rows.Next() {
	}
}
