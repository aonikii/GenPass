package main

import (
	"GenPass/internal/auth"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("web/templates/*.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { auth.Home(w, r, tmpl) })
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { auth.Login(w, r, tmpl) })
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) { auth.Register(w, r, tmpl) })

	log.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
