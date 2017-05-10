package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string    `gorm:"size:32;not null;unique"`
	Password string    `gorm:"size:255;not null"`
}

func SignIn(email, password string) (*User, error) {
	u := &User{
		Email:    email,
		Password: password,
	}

	if u.check() {
		return u, nil
	}

	return nil, fmt.Errorf("Login failed")
}

func (u User) check() bool {
	if u.Email == "caiwenpi@gmail.com" && u.Password == "123123" {
		return true
	}
	return false
}

func createAdmin() {
	db.FirstOrCreate(&User{
		Email:    "caiwenpi@gmail.com",
		Password: "123123",
	})
}
