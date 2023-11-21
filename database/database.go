package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kyon1313/books/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Migration() {
	DSN := DbCred()

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to the database")
	DB.AutoMigrate(&model.AuthorTable{}, &model.BookTable{}, model.Publisher{}, model.AuthorBook{})
}

func DbCred() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("HOST")
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")
	sslmode := os.Getenv("SSLMODE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, username, password, dbname, port, sslmode)
}
