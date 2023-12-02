package db

import (
	"fmt"
	"os"

	"github.com/FKuiv/LocalChat/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func Init() *Database {
	enverr := godotenv.Load()
	if enverr != nil {
		fmt.Println("Error loading the env file:", enverr)
	}

	db_user := os.Getenv("POSTGRES_USER")
	db_password := os.Getenv("POSTGRES_PASSWORD")
	db_host := os.Getenv("POSTGRES_HOST")
	db_port := os.Getenv("POSTGRES_PORT")
	db_dbname := os.Getenv("POSTGRES_DB")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_password, db_host, db_port, db_dbname)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Error opening connection to database")
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&models.User{}, &models.Message{}, &models.Group{}, &models.Session{})

	return &Database{DB: db}
}

func (db *Database) GetDB() *gorm.DB {
	return db.DB
}
