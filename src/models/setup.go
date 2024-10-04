package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(DbUser, DbPassword, DbHost, DbPort, DbName string) *gorm.DB {
	dsn := DbUser + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB = database
	if err := DB.AutoMigrate(&Product{}); err != nil {
		log.Fatalf("Auto-migration failed: %v", err)
	}

	return DB
}
