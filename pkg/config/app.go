package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

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

        d, err := sql.Open("mysql", dsn)
        
    if err != nil {
        log.Fatal("Erro ao abrir conexão com o banco de dados:", err)
    }

    if err := d.Ping(); err != nil {
        log.Fatal("Erro ao conectar ao banco de dados:", err)
    }
    
    fmt.Println("Conectado ao banco de dados com sucesso!")
    db =d
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)
}

// func CreateTable() {
//     query := `
//     CREATE TABLE IF NOT EXISTS books (
//         id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
//         created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
//         updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
//         deleted_at DATETIME(3) DEFAULT NULL,
//         name LONGTEXT,
//         author LONGTEXT,
//         publication LONGTEXT,
//         PRIMARY KEY (id),
//         KEY idx_books_deleted_at (deleted_at)
//     ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

//     // Execute the query to create the table
//     if _, err := db.Exec(query); err != nil {
//         log.Fatal("Error creating table:", err)
//     }
//     fmt.Println("Table 'books' created or already exists.")
// }

func GetDB() *sql.DB {
    if db == nil {
        log.Fatal("A conexão com o banco de dados não foi estabelecida.")
    }
    return db
}