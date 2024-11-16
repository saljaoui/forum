package repository

import (
	"fmt"
	"strings"
	"time"

	messages "forum-project/backend/internal/Messages"

	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int64     `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdat"`
	UUID      uuid.UUID `json:"uuid"`
}
type ResponceUser struct {
	Id        int64     `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	UUID      uuid.UUID `json:"uuid"`
}
type Login struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (users *User) Register() messages.Messages {
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

func (log *Login) Authentication() (ResponceUser, messages.Messages, uuid.UUID) {
	message := messages.Messages{}
	if log.Email == "" || !emailExists(log.Email) {
		message.MessageError = "Envalid Email"
		return ResponceUser{}, message, uuid.UUID{}
	} else {
		user := selectUser(log)
		if checkPasswordHash(user.Password, log.Password) {
			uuid, err := uuid.NewV4()
			if err != nil {
				fmt.Println("Error to Generate uuid", err)
			}
			loged := ResponceUser{
				Id:        user.Id,
				UUID:      uuid,
				Email:     user.Email,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
			}
			updateuuidUser(uuid, user.Id)
			return loged, messages.Messages{}, uuid
		} else {
			message.MessageError = "Email Or Password Encorect "
			return ResponceUser{}, message, uuid.UUID{}
		}
	}
}

func checkPasswordHash(hash, password string) bool {
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

func (us *User) AuthenticatLogin(UUID string) (m messages.Messages) {
	exists := checkAuthenticat(UUID)
	if exists {
		m.MessageError = " Unauthorized Token"
		return m
	}
	m.MessageSucc = "welcom"
	return m
}
