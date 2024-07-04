package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func InitPostgres() {
	//host := os.Getenv("DB_HOST")
	//port := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")
	//dbUser := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASSWORD")

	host := "localhost"
	port := "5432"
	dbName := "postgres"
	dbUser := "postgres"
	password := "postgres"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(Book{})

	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	fmt.Println("connected to database")
}
