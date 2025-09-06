package user

import (
	"GenPass/internal/database"
	"GenPass/internal/models"
	"GenPass/internal/password"
)

func newPasswordCreation(site, length, specialSymbol string, userId int) {
	pass := password.GeneratePass(length, specialSymbol)
	database.AddPasswordTodb(site, pass, userId)
}

func getPasswords(userId int) []models.Password {
	allPass := database.GetPasswords(userId)
	return allPass
}
