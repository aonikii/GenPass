package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectionDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	connStr := os.Getenv("connStr")

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка при открытии подключения:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}

	log.Println("Успешное подключение к базе данных!")
}
