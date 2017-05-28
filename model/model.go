package model

import (
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

// Init 初始化
func Init() {
	DB.AutoMigrate(&User{})

	admin := getAdmin()
	logrus.Debug(admin)
	if admin.ID == 0 {
		CreateUser("caiwenpi@gmail.com", "123123")
	}
}
