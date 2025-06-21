package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectionDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	connStr := os.Getenv("connStr")

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка при открытии подключения:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}

	log.Println("Успешное подключение к базе данных!")
	return db
}

func InsertUsers(username, passwordHash string) {
	_, err := db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, passwordHash)
	if err != nil {
		log.Panic(err)
	}
}

func CheckUserExists(u string) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT username FROM users WHERE username=$1)", u).Scan(&exists)
	if err != nil {
		log.Panic(err)
	}
	return exists
}

func GetHashPassword(u string) string {
	var hashpass string
	err := db.QueryRow("SELECT password_hash FROM users WHERE username=$1", u).Scan(&hashpass)
	if err != nil {
		log.Panic(err)
	}
	return hashpass
}
