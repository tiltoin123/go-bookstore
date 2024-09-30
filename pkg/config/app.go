package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var( db *gorm.DB)//what this line does

func Connect(){
	err := godotenv.Load()

    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    os.Getenv("DB_USER"), 
    os.Getenv("DB_PASSWORD"), 
    os.Getenv("DB_HOST"), 
    os.Getenv("DB_PORT"), 
    os.Getenv("DB_NAME"))

	d, err := gorm.Open("mysql",dsn)
	if err!= nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
    if db == nil {
        log.Fatal("A conexão com o banco de dados não foi estabelecida.")
    }
    return db
}
