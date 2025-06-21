package auth

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	switch r.Method {
	case "GET":
		tmpl.ExecuteTemplate(w, "login.html", nil)
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")
		res := enterCheck(username, password)
		if res != "ok" {
			tmpl.ExecuteTemplate(w, "login.html", map[string]string{"Error": res})
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func Register(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	switch r.Method {
	case "GET":
		tmpl.ExecuteTemplate(w, "register.html", nil)
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")
		res := registerCheck(username, password)
		if res != "ok" {
			tmpl.ExecuteTemplate(w, "register.html", map[string]string{"Error": res})
			return
		}
		addUserToDb(username, password)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
