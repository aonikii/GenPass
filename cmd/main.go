package main

import (
	"GenPass/internal/auth"
	"GenPass/internal/database"
	"GenPass/internal/sessions"
	"GenPass/internal/user"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("web/templates/*.html"))

func main() {
	db := database.ConnectionDb()
	defer db.Close()

	sessions.SessionsInit()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { auth.Home(w, r, tmpl) })
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { auth.Login(w, r, tmpl) })
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) { auth.Register(w, r, tmpl) })
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) { user.Dashboard(w, r, tmpl) })
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
