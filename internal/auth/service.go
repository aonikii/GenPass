package auth

import (
	"GenPass/internal/database"
	"log"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

func registerCheck(u, p string) string {
	if utf8.RuneCountInString(u) < 5 || utf8.RuneCountInString(p) < 5 {
		return "Имя и пароль должны быть >= 5"
	}
	userExists := database.CheckUserExists(u)
	if userExists {
		return "Пользователь с таким именем уже зарегистрирован"
	}
	return "ok"
}

func addUserToDb(u, p string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	database.InsertUsers(u, string(bytes))
}

func enterCheck(u, p string) string {
	userExists := database.CheckUserExists(u)
	if !userExists {
		return "Пользователь не найден"
	}
	err := bcrypt.CompareHashAndPassword([]byte(database.GetHashPassword(u)), []byte(p))
	if err != nil {
		return "Имя или пароль введены не верно"
	}
	return "ok"
}
