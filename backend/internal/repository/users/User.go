package user

import (
	"fmt"
	"strings"
	"time"

	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"

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
	Id        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	UUID      string `json:"uuid"`
}
type Login struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func generatUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error to Generate uuid", err)
	}
	return uuid.String()
}

func checkInputs() {
}

func (users *User) Register() (ResponceUser, messages.Messages, string) {
	message := messages.Messages{}
	uuid := generatUUID()
	loged := ResponceUser{
		Id:        users.Id,
		UUID:      uuid,
		Email:     users.Email,
		Firstname: users.Firstname,
		Lastname:  users.Lastname,
	}

	if strings.Trim(users.Firstname, " ") == "" || strings.Trim(users.Email, " ") == "" ||
		strings.Trim(users.Lastname, " ") == "" || strings.Trim(users.Password, " ") == "" {
		message.MessageError = "All Input is Required"
		return ResponceUser{}, message, ""
	}
	exists := emailExists(users.Email)
	if exists {
		message.MessageError = "Email user already exists"
		return ResponceUser{}, message, ""
	}

	password := hashPassword(users.Password)
	rows, err := insertUser(users, password)
	if err != nil {
		message.MessageError = "Error creating this user."
		return loged, message, uuid
	}

	user_id, err := rows.LastInsertId()
	if err != nil {
		message.MessageError = err.Error()
		return ResponceUser{}, message, ""
	} else {
		err = updateUUIDUser(uuid, user_id)
		if err != nil {
			fmt.Println("Error to Update")
		}
		message.MessageSucc = "User Created Successfully."
	}
	loged.Id = user_id
	return loged, message, uuid
}

func (log *Login) Authentication() (ResponceUser, messages.Messages, uuid.UUID) {
	message := messages.Messages{}

	st := "select * from user"
	rows := database.SelectRows(st)
	users := []User{}
	for rows.Next() {
		us := User{}
		rows.Scan(&us.Id, &us.Firstname, &us.Lastname, &us.Email, &us.Password, &us.CreatedAt,&us.UUID)
		users = append(users, us)
	}
	fmt.Println(users)

	if log.Email == "" || !emailExists(log.Email) {
		message.MessageError = "Invalid email"
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
				UUID:      uuid.String(),
				Email:     user.Email,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
			}
			err = updateUUIDUser(uuid.String(), user.Id)
			if err != nil {
				fmt.Println("Error to Update")
			}
			return loged, messages.Messages{}, uuid
		} else {
			message.MessageError = "Email or password incorrect."
			return ResponceUser{}, message, uuid.UUID{}
		}
	}
}

func (Log *Login) LogOut() (m messages.Messages) {
	err := updateUUIDUser("null", Log.Id)
	if err != nil {
		m.MessageError = "Error To Update user"
		return m
	} else {
		m.MessageSucc = "Update Seccesfly"
		return m
	}
}

func checkPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error", err)
	}
	return string(hashedPassword)
}

func (us *User) AuthenticatLogin(UUID string) (m messages.Messages) {
	exists := checkAuthenticat(UUID)
	if !exists {
		m.MessageError = "Unauthorized token"
	}
	return
}
