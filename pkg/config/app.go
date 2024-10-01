package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
    // Load the .env file
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env", err)
    }

    // Prepare the DSN
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"))

    // Open the database connection
    d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Erro ao conectar com o banco de dados", err)
    }

    db = d
}

func GetDB() *gorm.DB {
    if db == nil {
        log.Fatal("A conexão com o banco de dados não foi estabelecida.")
    }
    return db
}
