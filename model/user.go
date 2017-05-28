package model

import (
	"crypto/sha256"
	"fmt"

	"github.com/dchest/uniuri"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/pbkdf2"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:32;not null;unique"`
	Password string `gorm:"size:255;not null"`
	Token    string `gorm:"size:16;unique"`
}

func GetUserByToken(tk string) *User {
	u := new(User)
	DB.Where(&User{Token: tk}).First(&u)
	return u
}

func getAdmin() *User {
	u := new(User)
	DB.Where(&User{Email: "caiwenpi@gmail.com"}).First(&u)
	return u
}

func CreateUser(email, password string) {
	DB.FirstOrCreate(&User{
		Email:    email,
		Password: EncodePasswd(password),
		Token:    uniuri.New(),
	})
}

func EncodePasswd(s string) string {
	return fmt.Sprintf("%x", pbkdf2.Key([]byte(s), []byte("bnjkjwh"), 10000, 16, sha256.New))
}

func LoginIn(email, password string) *User {
	u := new(User)
	DB.Where(&User{
		Email:    email,
		Password: EncodePasswd(password),
	}).First(&u)
	if u.ID == 0 {
		return u
	}

	u.Token = uniuri.New()
	DB.Save(&u)

	return u
}
