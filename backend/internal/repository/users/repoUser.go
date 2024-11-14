package repository

import (
	"fmt"
	"strings"

	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/models"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	models.User
}

func Register(users *models.User) messages.Messages {
	message := messages.Messages{}
	if strings.Trim(users.Firstname, " ") == "" || strings.Trim(users.Email, " ") == "" ||
		strings.Trim(users.Lastname, " ") == "" || strings.Trim(users.Password, " ") == "" {
		message.MessageError = "All Input is Requerd"
		message.ErrorBool = true
		return message
	}
	exists := emailExists(users.Email)
	if exists {
		message.MessageError = "Email User Is Allready Exsist"
	} else {
		password := hashPassword(users.Password)
		err := insertUser(users, password)
		if err != nil {
			message.MessageError = "Error To Create this user"
		} else {
			message.MessageSucc = "User created successfully"
		}
	}
	return message
}

func Login(log *models.Login) (models.ResponceUser, messages.Messages) {
	message := messages.Messages{}
	if log.Email == "" || !emailExists(log.Email) {
		message.MessageError = "Envalid Email"
		return models.ResponceUser{}, message
	} else {
		user := selectUser(log)
		if CheckPasswordHash(user.Password, log.Password) {
			uuid, err := uuid.NewV4()
			if err != nil {
				fmt.Println("Error to Generate uuid", err)
			}
			loged := models.ResponceUser{
				Id:        user.Id,
				UUID:      uuid,
				Email:     user.Email,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
			}
			updateuuidUser(uuid, user.Id)
			return loged, messages.Messages{}
		} else {
			message.MessageError = "Email Or Password Encorect "
			return models.ResponceUser{}, message
		}
	}
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) string {
	fmt.Println(password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error", err)
	}
	return string(hashedPassword)
}
