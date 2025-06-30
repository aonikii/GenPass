package database

import (
	"GenPass/internal/models"
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

func GetUserInfo(u string) *models.User {
	var userInfo models.User

	err := db.QueryRow(
		"SELECT id, username, password_hash FROM users WHERE username=$1", u).Scan(&userInfo.Id, &userInfo.Username, &userInfo.PasswordHash)
	if err != nil {
		log.Panic(err)
	}
	return &userInfo
}

func AddPasswordTodb(site, pass string, userId int) {
	_, err := db.Exec("INSERT INTO passwords (site, password, user_id) VALUES ($1, $2, $3)", site, pass, userId)
	if err != nil {
		log.Panic(err)
	}
}
