package resource

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"os"


)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() Database {


	if err := godotenv.Load(); err != nil {
        panic("Error loading .env file")
    }
	port := os.Getenv("SERVER_PORT")
	fmt.Println("_________________", port)
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_HOST"),os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: nil})
	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + "receipt"+ ";")
	if err != nil {
		panic(fmt.Sprintf("Error", err))
	
	}
	return Database{
		DB: db,
	}
}