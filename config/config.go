package config

import (
	"prakerja6/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja6?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	migration()
}

func migration() {
	DB.AutoMigrate(&models.User{})
}
