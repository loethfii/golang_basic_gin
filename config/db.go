package config

import (
	"golang_basic_gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_basic_sql_2?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(&models.Department{})
	DB.AutoMigrate(&models.Employee{})
	DB.AutoMigrate(&models.Position{})
	DB.AutoMigrate(&models.User{})

	return DB, err
}
