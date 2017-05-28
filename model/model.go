package model

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// Init 初始化
func Init() error {
	var err error

	dsn := fmt.Sprintf("%s:%s@/bgy?charset=utf8&parseTime=True&loc=Local", "root", "123456")
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB.AutoMigrate(&User{})

	admin := getAdmin()
	logrus.Debug(admin)
	if admin.ID == 0 {
		CreateUser("caiwenpi@gmail.com", "123123")
	}

	return nil
}
