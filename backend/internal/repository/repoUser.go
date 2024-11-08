package repository

import (
	"database/sql"
	"fmt"

	"forum-project/backend/internal/models"
)

type Connect struct {
	dbConect models.DB
}

func NewConnect(db *sql.DB) Connect {
	return Connect{
		dbConect: models.DB{
			Db: db,
		},
	}
}

func (cnx Connect) Register(users *models.User) error {
	stm := "INSERT INTO users (firstname,lastname,email,password) VALUES(?,?,?,?)"
	_, err := cnx.dbConect.Db.Exec(stm, users.Firstname, users.Lastname, users.Email, users.Password)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(users.Firstname,users.Lastname,users.Email)
	return nil
}
