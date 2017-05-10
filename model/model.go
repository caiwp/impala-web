package model

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func NewDB(user, password string) {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/bgy?charset=utf8&parseTime=True&loc=Local", user, password))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	logrus.Info("open mysql")

	db.AutoMigrate(&User{})

	createAdmin()
}
