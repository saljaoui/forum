package models

import (
	"golang.org/x/crypto/bcrypt"
)
 

type Post struct {
	ID        int64
	UserID  int64
	Title     string
	Content  string
	CreatedAt time.Time
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
