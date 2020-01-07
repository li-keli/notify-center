package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var conn *gorm.DB

func NewDB() {
	db, err := gorm.Open("mysql", "root:123@tcp(172.16.2.161:3300)/notify?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logrus.Fatalln("数据库连接失败", err)
	}
	logrus.Info("MySql连接成功...")
	//db.LogMode(true)
	db.SingularTable(true)

	conn = db
}
