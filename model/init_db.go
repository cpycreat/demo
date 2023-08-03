package model

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// 数据库配置
const (
	HOST     = "127.0.0.1"
	PORT     = "3306"
	USER     = "root"
	PASS     = "sztu@150037"
	DATABASE = "dy"
)

var DB *gorm.DB

func InitDB() {
	//连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DATABASE)
	fmt.Println(dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

}
