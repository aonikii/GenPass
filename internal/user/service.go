package user

import (
	"GenPass/internal/database"
	"GenPass/internal/password"
	"log"
)

func newPasswordCreation(site, length, specialSymbol string, userId int) {
	pass := password.GeneratePass(length, specialSymbol)
	database.AddPasswordTodb(site, pass, userId)
}

func getPasswords(userId int) {
	allPass := database.GetPasswords(userId)
	for _, v := range allPass {
		log.Println(v)
	}
}
