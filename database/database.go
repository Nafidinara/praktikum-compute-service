package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Praktikum/models"
	"Praktikum/utils"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// var dsn string = "root:afara123@tcp(127.0.0.1:3307)/deliveries?charset=utf8mb4&parseTime=True&loc=Local"

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetConfig("DB_USERNAME"),
		utils.GetConfig("DB_PASSWORD"),
		utils.GetConfig("DB_HOST"),
		utils.GetConfig("DB_PORT"),
		utils.GetConfig("DB_NAME"),
	)

	fmt.Println("ssssssss")
	fmt.Println("ssssssss:" + dsn)
	fmt.Println("ssssssss")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatalf("error when creating a connection to the database: % \n", err)
	}

	log.Println("connected to the database")

}

func Migrate() {
	err := DB.AutoMigrate(models.User{})
	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}
