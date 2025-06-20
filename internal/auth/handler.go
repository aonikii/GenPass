package auth

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	err := tmpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
	}
}

func Register(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	err := tmpl.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
	}
}
