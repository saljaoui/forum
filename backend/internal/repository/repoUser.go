package repository

import (
	"fmt"
	"strings"

	messages "forum-project/backend/internal/Messages"
	"forum-project/backend/internal/database"
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
		message.ErrorBool = true
	} else {
		password := hashPassword(users.Password)
		stm := "INSERT INTO user (firstname,lastname,email,password) VALUES(?,?,?,?)"

		database.Exec(stm, users.Firstname, users.Lastname, users.Email, password)
		message.MessageSucc = "User created successfully"
	}
	return message
}

func Login(log *models.Login) (models.Loged, messages.Messages) {
	user := models.User{}
	message := messages.Messages{}
	db := database.Config()
	if log.Email == "" || !emailExists(log.Email) {
		message.ErrorBool = true
		message.MessageError = "Envalid Email"
		return models.Loged{}, message
	} else {

		query := "select id,email,password, firstname ,lastname FROM user where email=?"
		err := db.QueryRow(query, log.Email, log.Password).Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname)
		if err != nil {
			fmt.Println("Error ", err)
		}

		if CheckPasswordHash(user.Password, log.Password) {
			uuid, err := uuid.NewV4()
			if err != nil {
				fmt.Println("Error to Generate uuid", err)
			}
			loged := models.Loged{
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
			return models.Loged{}, message
		}
	}
}

func updateuuidUser(uudi uuid.UUID, userId int64) {
	db := database.Config()
	_, err := db.Exec("UPDATE user SET UUID=? WHERE id=?", uudi, userId)
	if err != nil {
		fmt.Println("Error To Update User uuid")
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

func emailExists(email string) bool {
	var exists bool
	query := "SELECT EXISTS (select email from user where email=?)"
	database.SelectOneRow(query, email, &exists)
	return exists
}
