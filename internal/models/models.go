package models

type User struct {
	Id           int
	Username     string
	PasswordHash string
}

type Password struct {
	Site         string
	PasswordText string
}
