package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func Connect() {
    err := godotenv.Load("../.env")

    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env", err)
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"))
        fmt.Println(dsn)
        d, err := sql.Open("mysql", dsn)
        
    if err != nil {
        log.Fatal("Erro ao abrir conexão com o banco de dados:", err)
    }

    if err := d.Ping(); err != nil {
        log.Fatal("Erro ao conectar ao banco de dados:", err)
    }
    
    fmt.Println("Conectado ao banco de dados com sucesso!")
    db =d
}

func GetDB() *sql.DB {
    fmt.Println("get db @#$@#$$@#@#$",db.Stats())
    if db == nil {
        log.Fatal("A conexão com o banco de dados não foi estabelecida.")
    }
    return db
}