package setting

import (
	"fmt"
	"path/filepath"

	"github.com/caiwp/impala-web/model"

	"os"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

const (
	TimestampFormat string = "2006-01-02 15:04:05"
)

var (
	cfg *viper.Viper = viper.New()
)

func NewContext(c *cli.Context) error {
	NewConfig(c)
	err := NewDBService(c)
	if err != nil {
		return err
	}
	return nil
}

func NewConfig(c *cli.Context) {
	root := c.String("config")
	f := filepath.Join(root, "local.yaml")
	_, err := os.Stat(f)
	if err != nil {
		f = filepath.Join(root, "app.yaml")
	}
	cfg.SetConfigType("yaml")
	cfg.SetConfigFile(f)

	err = cfg.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Infof("Config file: %s", f)
	logrus.Info(cfg.AllSettings())
}

func NewDBService(c *cli.Context) error {
	user := cfg.GetString("MYSQL.USER")
	pass := cfg.GetString("MYSQL.PASSWORD")
	host := cfg.GetString("MYSQL.HOST")
	port := cfg.GetInt("MYSQL.PORT")
	database := cfg.GetString("MYSQL.DATABASE")

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, database)
	logrus.Debug(dsn)

	model.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}

	model.Init()
	return nil
}
