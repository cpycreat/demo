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
		panic("无法连接数据库:" + err.Error())
	}
	//自动迁移

	err = DB.AutoMigrate(&User{}, &Vedio{})
	if err != nil {
		panic("数据库迁移失败:" + err.Error())
	}
	fmt.Println("数据库初始化成功")
	// tempuser := User{
	// 	Username: "admin1",
	// 	Password: "admin",
	// 	UpVedio: []*Vedio{
	// 		{
	// 			VID:      1,
	// 			VedioURL: "http://localhost:8080/static/video/1.mp4",
	// 			CoverURL: "http://localhost:8080/static/video/1.jpg",
	// 		},
	// 		{
	// 			VID:      2,
	// 			VedioURL: "http://localhost:8080/static/video/2.mp4",
	// 			CoverURL: "http://localhost:8080/static/video/2.jpg",
	// 		},
	// 	},
	// }

	// DB.Create(&tempuser)
}
